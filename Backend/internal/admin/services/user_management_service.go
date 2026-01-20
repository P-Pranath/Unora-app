// internal/admin/services/user_management_service.go
package services

import (
	"context"
	"fmt"
	"time"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/connection"
	"github.com/UnoraApp/be/ent/generated/credittransaction"
	"github.com/UnoraApp/be/ent/generated/user"
	"github.com/UnoraApp/be/ent/generated/userreport"
	"github.com/UnoraApp/be/internal/admin/dto"
	"github.com/google/uuid"
)

// UserManagementService handles admin user management
type UserManagementService struct {
	entClient *ent.Client
}

// NewUserManagementService creates a new user management service
func NewUserManagementService(entClient *ent.Client) *UserManagementService {
	return &UserManagementService{
		entClient: entClient,
	}
}

// ListUsers returns paginated list of users
func (s *UserManagementService) ListUsers(ctx context.Context, page, pageSize int, status, search string) (*dto.AdminUserListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	query := s.entClient.User.Query()

	// Apply status filter
	if status != "" {
		query = query.Where(user.AccountStatusEQ(user.AccountStatus(status)))
	}

	// Apply search filter
	if search != "" {
		query = query.Where(
			user.Or(
				user.EmailContainsFold(search),
				user.FirstNameContainsFold(search),
				user.LastNameContainsFold(search),
			),
		)
	}

	// Get total count
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count users: %w", err)
	}

	// Get users
	users, err := query.
		Order(ent.Desc(user.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	result := make([]dto.AdminUserResponse, len(users))
	for i, u := range users {
		result[i] = s.userToAdminResponse(ctx, u)
	}

	totalPages := (total + pageSize - 1) / pageSize

	return &dto.AdminUserListResponse{
		Users:      result,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetUser returns detailed user info
func (s *UserManagementService) GetUser(ctx context.Context, userID string) (*dto.AdminUserResponse, error) {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	result := s.userToAdminResponse(ctx, u)
	return &result, nil
}

// SuspendUser suspends a user account
func (s *UserManagementService) SuspendUser(ctx context.Context, userID string, req *dto.SuspendUserRequest) error {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	now := time.Now()
	_, err = u.Update().
		SetAccountStatus(user.AccountStatusSuspended).
		SetSuspendedAt(now).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to suspend user: %w", err)
	}

	return nil
}

// UnsuspendUser reactivates a suspended user
func (s *UserManagementService) UnsuspendUser(ctx context.Context, userID string) error {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	_, err = u.Update().
		SetAccountStatus(user.AccountStatusActive).
		ClearSuspendedAt().
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to unsuspend user: %w", err)
	}

	return nil
}

// DeleteUser soft deletes a user
func (s *UserManagementService) DeleteUser(ctx context.Context, userID string) error {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	now := time.Now()
	_, err = u.Update().
		SetAccountStatus(user.AccountStatusDeleted).
		SetDeletedAt(now).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// AdjustCredits adjusts user credit balance
func (s *UserManagementService) AdjustCredits(ctx context.Context, userID string, req *dto.AdjustCreditsRequest) error {
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	newBalance := u.CreditBalance + req.Amount
	if newBalance < 0 {
		return fmt.Errorf("cannot reduce balance below zero")
	}

	// Update balance
	_, err = u.Update().
		SetCreditBalance(newBalance).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update balance: %w", err)
	}

	// Create transaction record
	txID := uuid.New().String()
	_, err = s.entClient.CreditTransaction.
		Create().
		SetID(txID).
		SetUserID(userID).
		SetTransactionType(credittransaction.TransactionTypeAdminAdjustment).
		SetCreditAmount(req.Amount).
		SetBalanceAfter(newBalance).
		SetNillableDescription(strPtr(fmt.Sprintf("%s: %s", req.Reason, req.Description))).
		SetNillableReferenceType(strPtr("admin_adjustment")).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to record transaction: %w", err)
	}

	return nil
}

func (s *UserManagementService) userToAdminResponse(ctx context.Context, u *ent.User) dto.AdminUserResponse {
	// Count photos
	photoCount, _ := s.entClient.Photo.
		Query().
		Where().
		Count(ctx)

	// Count connections
	connectionCount, _ := s.entClient.Connection.
		Query().
		Where(
			connection.Or(
				connection.UserAIDEQ(u.ID),
				connection.UserBIDEQ(u.ID),
			),
		).
		Where(connection.ConnectionStatusEQ(connection.ConnectionStatusActive)).
		Count(ctx)

	// Count reports received
	reportsReceived, _ := s.entClient.UserReport.
		Query().
		Where(userreport.ReportedUserIDEQ(u.ID)).
		Count(ctx)

	return dto.AdminUserResponse{
		ID:               u.ID,
		Email:            ptrToString(u.Email),
		PhoneNumber:      ptrToString(u.PhoneNumber),
		FirstName:        ptrToString(u.FirstName),
		LastName:         ptrToString(u.LastName),
		AccountStatus:    string(u.AccountStatus),
		OnboardingStatus: string(u.OnboardingStatus),
		CreditBalance:    u.CreditBalance,
		PhotoCount:       photoCount,
		ConnectionCount:  connectionCount,
		ReportsReceived:  reportsReceived,
		CreatedAt:        u.CreatedAt,
		LastActiveAt:     u.LastActiveAt,
		SuspendedAt:      u.SuspendedAt,
	}
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
