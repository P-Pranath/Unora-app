// internal/safety/services/safety_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/connection"
	"github.com/UnoraApp/be/ent/generated/interest"
	"github.com/UnoraApp/be/ent/generated/userblock"
	"github.com/UnoraApp/be/ent/generated/userreport"
	"github.com/UnoraApp/be/internal/safety/dto"
)

// SafetyService handles safety-related business logic
type SafetyService struct {
	entClient *ent.Client
}

// NewSafetyService creates a new safety service
func NewSafetyService(entClient *ent.Client) *SafetyService {
	return &SafetyService{
		entClient: entClient,
	}
}

// BlockUser blocks a user
func (s *SafetyService) BlockUser(ctx context.Context, blockerID string, req *dto.BlockUserRequest) (*dto.BlockResponse, error) {
	if blockerID == req.UserID {
		return nil, fmt.Errorf("cannot block yourself")
	}

	// Check if already blocked
	exists, _ := s.entClient.UserBlock.
		Query().
		Where(userblock.BlockerUserIDEQ(blockerID)).
		Where(userblock.BlockedUserIDEQ(req.UserID)).
		Where(userblock.UnblockedAtIsNil()).
		Exist(ctx)
	if exists {
		return nil, fmt.Errorf("user already blocked")
	}

	// Create block
	blockID := uuid.New().String()
	block, err := s.entClient.UserBlock.
		Create().
		SetID(blockID).
		SetBlockerUserID(blockerID).
		SetBlockedUserID(req.UserID).
		SetNillableReason(strPtr(req.Reason)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to block user: %w", err)
	}

	// Terminate any existing connection
	s.terminateConnection(ctx, blockerID, req.UserID)

	// Cancel any pending interests
	s.cancelInterests(ctx, blockerID, req.UserID)

	// Get blocked user name
	blockedUser, _ := s.entClient.User.Get(ctx, req.UserID)
	blockedName := ""
	if blockedUser != nil && blockedUser.FirstName != nil {
		blockedName = *blockedUser.FirstName
	}

	return &dto.BlockResponse{
		ID:            block.ID,
		BlockedUserID: block.BlockedUserID,
		BlockedName:   blockedName,
		CreatedAt:     block.CreatedAt,
	}, nil
}

// UnblockUser unblocks a user
func (s *SafetyService) UnblockUser(ctx context.Context, blockerID string, req *dto.UnblockUserRequest) error {
	// Find active block
	block, err := s.entClient.UserBlock.
		Query().
		Where(userblock.BlockerUserIDEQ(blockerID)).
		Where(userblock.BlockedUserIDEQ(req.UserID)).
		Where(userblock.UnblockedAtIsNil()).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("user is not blocked")
		}
		return fmt.Errorf("failed to find block: %w", err)
	}

	// Update block with unblocked time
	_, err = block.Update().
		SetUnblockedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to unblock: %w", err)
	}

	return nil
}

// GetBlockedUsers returns list of users blocked by the requester
func (s *SafetyService) GetBlockedUsers(ctx context.Context, userID string) (*dto.BlockedUsersResponse, error) {
	blocks, err := s.entClient.UserBlock.
		Query().
		Where(userblock.BlockerUserIDEQ(userID)).
		Where(userblock.UnblockedAtIsNil()).
		Order(ent.Desc(userblock.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get blocks: %w", err)
	}

	result := make([]dto.BlockResponse, len(blocks))
	for i, b := range blocks {
		blockedUser, _ := s.entClient.User.Get(ctx, b.BlockedUserID)
		blockedName := ""
		if blockedUser != nil && blockedUser.FirstName != nil {
			blockedName = *blockedUser.FirstName
		}

		result[i] = dto.BlockResponse{
			ID:            b.ID,
			BlockedUserID: b.BlockedUserID,
			BlockedName:   blockedName,
			CreatedAt:     b.CreatedAt,
		}
	}

	return &dto.BlockedUsersResponse{
		Blocks: result,
		Total:  len(result),
	}, nil
}

// GetBlockStatus checks if a user is blocked or blocking
func (s *SafetyService) GetBlockStatus(ctx context.Context, userID, otherUserID string) (*dto.BlockStatusResponse, error) {
	// Check if I blocked them
	isBlocked, _ := s.entClient.UserBlock.
		Query().
		Where(userblock.BlockerUserIDEQ(userID)).
		Where(userblock.BlockedUserIDEQ(otherUserID)).
		Where(userblock.UnblockedAtIsNil()).
		Exist(ctx)

	// Check if they blocked me
	isBlockedBy, _ := s.entClient.UserBlock.
		Query().
		Where(userblock.BlockerUserIDEQ(otherUserID)).
		Where(userblock.BlockedUserIDEQ(userID)).
		Where(userblock.UnblockedAtIsNil()).
		Exist(ctx)

	return &dto.BlockStatusResponse{
		IsBlocked:   isBlocked,
		IsBlockedBy: isBlockedBy,
	}, nil
}

// IsBlocked checks if there's any block between two users
func (s *SafetyService) IsBlocked(ctx context.Context, userA, userB string) bool {
	blocked, _ := s.entClient.UserBlock.
		Query().
		Where(
			userblock.Or(
				userblock.And(
					userblock.BlockerUserIDEQ(userA),
					userblock.BlockedUserIDEQ(userB),
				),
				userblock.And(
					userblock.BlockerUserIDEQ(userB),
					userblock.BlockedUserIDEQ(userA),
				),
			),
		).
		Where(userblock.UnblockedAtIsNil()).
		Exist(ctx)
	return blocked
}

// GetBlockedUserIDs returns IDs of users blocked by or blocking the user
func (s *SafetyService) GetBlockedUserIDs(ctx context.Context, userID string) ([]string, error) {
	blocks, err := s.entClient.UserBlock.
		Query().
		Where(
			userblock.Or(
				userblock.BlockerUserIDEQ(userID),
				userblock.BlockedUserIDEQ(userID),
			),
		).
		Where(userblock.UnblockedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	userIDs := make([]string, 0, len(blocks))
	seen := make(map[string]bool)
	for _, b := range blocks {
		if b.BlockerUserID != userID && !seen[b.BlockerUserID] {
			userIDs = append(userIDs, b.BlockerUserID)
			seen[b.BlockerUserID] = true
		}
		if b.BlockedUserID != userID && !seen[b.BlockedUserID] {
			userIDs = append(userIDs, b.BlockedUserID)
			seen[b.BlockedUserID] = true
		}
	}
	return userIDs, nil
}

// ReportUser reports a user
func (s *SafetyService) ReportUser(ctx context.Context, reporterID string, req *dto.ReportUserRequest) (*dto.ReportResponse, error) {
	if reporterID == req.UserID {
		return nil, fmt.Errorf("cannot report yourself")
	}

	// Create report
	reportID := uuid.New().String()
	report, err := s.entClient.UserReport.
		Create().
		SetID(reportID).
		SetReporterUserID(reporterID).
		SetReportedUserID(req.UserID).
		SetReportReason(userreport.ReportReason(req.Reason)).
		SetNillableDescription(strPtr(req.Description)).
		SetNillableReferenceType(strPtr(req.ReferenceType)).
		SetNillableReferenceID(strPtr(req.ReferenceID)).
		SetReportStatus(userreport.ReportStatusPending).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create report: %w", err)
	}

	return &dto.ReportResponse{
		ID:             report.ID,
		ReportedUserID: report.ReportedUserID,
		Reason:         string(report.ReportReason),
		Description:    ptrToString(report.Description),
		Status:         string(report.ReportStatus),
		CreatedAt:      report.CreatedAt,
	}, nil
}

// Helper: terminate connection between users
func (s *SafetyService) terminateConnection(ctx context.Context, userA, userB string) {
	// Ensure ordering
	if userA > userB {
		userA, userB = userB, userA
	}

	s.entClient.Connection.
		Update().
		Where(connection.UserAIDEQ(userA)).
		Where(connection.UserBIDEQ(userB)).
		Where(connection.ConnectionStatusEQ(connection.ConnectionStatusActive)).
		SetConnectionStatus(connection.ConnectionStatusTerminated).
		SetTerminatedAt(time.Now()).
		Save(ctx)
}

// Helper: cancel pending interests between users
func (s *SafetyService) cancelInterests(ctx context.Context, userA, userB string) {
	s.entClient.Interest.
		Update().
		Where(
			interest.Or(
				interest.And(
					interest.SenderUserIDEQ(userA),
					interest.ReceiverUserIDEQ(userB),
				),
				interest.And(
					interest.SenderUserIDEQ(userB),
					interest.ReceiverUserIDEQ(userA),
				),
			),
		).
		Where(interest.InterestStatusEQ(interest.InterestStatusPending)).
		SetInterestStatus(interest.InterestStatusExpired).
		Save(ctx)
}

// Helper functions
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
