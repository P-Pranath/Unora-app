// internal/streak/services/streak_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/checkin"
	"github.com/UnoraApp/be/ent/generated/connection"
	"github.com/UnoraApp/be/ent/generated/nudge"
	"github.com/UnoraApp/be/ent/generated/photo"
	"github.com/UnoraApp/be/ent/generated/streak"
	"github.com/UnoraApp/be/internal/streak/dto"
	"github.com/UnoraApp/be/pkg/storage"
)

// StreakService handles streak-related business logic
type StreakService struct {
	entClient     *ent.Client
	storageClient storage.Client
}

// NewStreakService creates a new streak service
func NewStreakService(entClient *ent.Client, storageClient storage.Client) *StreakService {
	return &StreakService{
		entClient:     entClient,
		storageClient: storageClient,
	}
}

// GetStreakByConnection gets the streak for a connection
func (s *StreakService) GetStreakByConnection(ctx context.Context, userID, connectionID string) (*dto.StreakResponse, error) {
	// Verify user is part of connection
	conn, err := s.entClient.Connection.
		Query().
		Where(connection.IDEQ(connectionID)).
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		WithStreak(func(q *ent.StreakQuery) {
			q.WithCheckIns()
		}).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("connection not found")
		}
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	if conn.Edges.Streak == nil {
		return nil, fmt.Errorf("streak not found for this connection")
	}

	return s.streakToResponse(ctx, conn.Edges.Streak, userID, conn)
}

// CheckIn performs a daily check-in
func (s *StreakService) CheckIn(ctx context.Context, userID, connectionID string, req *dto.CheckInRequest) (*dto.CheckInResponse, error) {
	// Get the streak
	conn, err := s.entClient.Connection.
		Query().
		Where(connection.IDEQ(connectionID)).
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		WithStreak().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("connection not found: %w", err)
	}

	st := conn.Edges.Streak
	if st == nil {
		return nil, fmt.Errorf("streak not found")
	}

	// Check if streak is active
	if st.StreakState != streak.StreakStateActive && st.StreakState != streak.StreakStateAtRisk {
		return nil, fmt.Errorf("streak is not active")
	}

	// Check if already checked in today
	today := time.Now().Truncate(24 * time.Hour)
	exists, _ := s.entClient.CheckIn.
		Query().
		Where(checkin.StreakIDEQ(st.ID)).
		Where(checkin.UserIDEQ(userID)).
		Where(checkin.DayNumberEQ(st.CurrentDay)).
		Where(checkin.CreatedAtGTE(today)).
		Exist(ctx)
	if exists {
		return nil, fmt.Errorf("already checked in today")
	}

	// Create check-in
	checkInID := uuid.New().String()
	eventData := make(map[string]interface{})
	if req != nil && req.Activity != "" {
		eventData["activity"] = req.Activity
	}

	ci, err := s.entClient.CheckIn.
		Create().
		SetID(checkInID).
		SetStreakID(st.ID).
		SetUserID(userID).
		SetDayNumber(st.CurrentDay).
		SetCheckInType(checkin.CheckInTypeManual).
		SetEventData(eventData).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create check-in: %w", err)
	}

	// Check if both users have checked in today
	partnerID := conn.UserAID
	if conn.UserAID == userID {
		partnerID = conn.UserBID
	}

	partnerCheckedIn, _ := s.entClient.CheckIn.
		Query().
		Where(checkin.StreakIDEQ(st.ID)).
		Where(checkin.UserIDEQ(partnerID)).
		Where(checkin.DayNumberEQ(st.CurrentDay)).
		Where(checkin.CreatedAtGTE(today)).
		Exist(ctx)

	// If both checked in, advance to next day
	if partnerCheckedIn {
		newDay := st.CurrentDay + 1
		newState := streak.StreakStateActive

		if newDay > 15 {
			// Streak completed!
			newDay = 15
			newState = streak.StreakStateCompleted
			_, err = st.Update().
				SetCurrentDay(newDay).
				SetStreakState(newState).
				SetCompletedAt(time.Now()).
				Save(ctx)
		} else {
			_, err = st.Update().
				SetCurrentDay(newDay).
				SetStreakState(newState).
				Save(ctx)
		}
		if err != nil {
			// Log error but don't fail the check-in
		}
	}

	activity := ""
	if v, ok := eventData["activity"].(string); ok {
		activity = v
	}

	return &dto.CheckInResponse{
		ID:          ci.ID,
		DayNumber:   ci.DayNumber,
		UserID:      ci.UserID,
		CheckInType: string(ci.CheckInType),
		Activity:    activity,
		CreatedAt:   ci.CreatedAt,
	}, nil
}

// GetTodayStreaks returns all streaks requiring check-in today
func (s *StreakService) GetTodayStreaks(ctx context.Context, userID string) (*dto.TodayStreaksResponse, error) {
	// Get all active connections for the user
	connections, err := s.entClient.Connection.
		Query().
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		Where(connection.ConnectionStatusEQ(connection.ConnectionStatusActive)).
		Where(connection.DeletedAtIsNil()).
		WithStreak().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connections: %w", err)
	}

	today := time.Now().Truncate(24 * time.Hour)
	items := make([]dto.TodayStreakItem, 0)
	totalPending := 0
	totalCompleted := 0

	for _, conn := range connections {
		st := conn.Edges.Streak
		if st == nil || (st.StreakState != streak.StreakStateActive && st.StreakState != streak.StreakStateAtRisk) {
			continue
		}

		partnerID := conn.UserAID
		if conn.UserAID == userID {
			partnerID = conn.UserBID
		}

		// Check if user checked in today
		myCheckIn, _ := s.entClient.CheckIn.
			Query().
			Where(checkin.StreakIDEQ(st.ID)).
			Where(checkin.UserIDEQ(userID)).
			Where(checkin.DayNumberEQ(st.CurrentDay)).
			Where(checkin.CreatedAtGTE(today)).
			Exist(ctx)

		// Check if partner checked in today
		partnerCheckIn, _ := s.entClient.CheckIn.
			Query().
			Where(checkin.StreakIDEQ(st.ID)).
			Where(checkin.UserIDEQ(partnerID)).
			Where(checkin.DayNumberEQ(st.CurrentDay)).
			Where(checkin.CreatedAtGTE(today)).
			Exist(ctx)

		// Get partner info
		partner, _ := s.entClient.User.Get(ctx, partnerID)
		partnerPhoto, _ := s.entClient.Photo.
			Query().
			Where(photo.UserIDEQ(partnerID)).
			Where(photo.DeletedAtIsNil()).
			Order(ent.Asc(photo.FieldDisplayOrder)).
			First(ctx)

		partnerName := ""
		partnerPhotoURL := ""
		if partner != nil && partner.FirstName != nil {
			partnerName = *partner.FirstName
		}
		if partnerPhoto != nil {
			partnerPhotoURL, _ = s.storageClient.GetPresignedDownloadURL(ctx, partnerPhoto.StorageKey)
		}

		needsCheckIn := !myCheckIn

		item := dto.TodayStreakItem{
			StreakID:         st.ID,
			ConnectionID:     conn.ID,
			State:            string(st.StreakState),
			CurrentDay:       st.CurrentDay,
			PartnerName:      partnerName,
			PartnerPhotoURL:  partnerPhotoURL,
			NeedsCheckIn:     needsCheckIn,
			PartnerCheckedIn: partnerCheckIn,
		}
		items = append(items, item)

		if needsCheckIn {
			totalPending++
		} else {
			totalCompleted++
		}
	}

	return &dto.TodayStreaksResponse{
		Streaks:        items,
		TotalPending:   totalPending,
		TotalCompleted: totalCompleted,
	}, nil
}

// SendNudge sends a nudge to streak partner
func (s *StreakService) SendNudge(ctx context.Context, userID, connectionID string, req *dto.SendNudgeRequest) (*dto.NudgeResponse, error) {
	// Get connection and streak
	conn, err := s.entClient.Connection.
		Query().
		Where(connection.IDEQ(connectionID)).
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		WithStreak().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("connection not found: %w", err)
	}

	st := conn.Edges.Streak
	if st == nil {
		return nil, fmt.Errorf("streak not found")
	}

	receiverID := conn.UserAID
	if conn.UserAID == userID {
		receiverID = conn.UserBID
	}

	// Check if already nudged today
	today := time.Now().Truncate(24 * time.Hour)
	exists, _ := s.entClient.Nudge.
		Query().
		Where(nudge.StreakIDEQ(st.ID)).
		Where(nudge.SenderUserIDEQ(userID)).
		Where(nudge.DayNumberEQ(st.CurrentDay)).
		Where(nudge.CreatedAtGTE(today)).
		Exist(ctx)
	if exists {
		return nil, fmt.Errorf("already sent a nudge today")
	}

	// Create nudge
	nudgeID := uuid.New().String()
	message := ""
	if req != nil {
		message = req.Message
	}

	n, err := s.entClient.Nudge.
		Create().
		SetID(nudgeID).
		SetStreakID(st.ID).
		SetSenderUserID(userID).
		SetReceiverUserID(receiverID).
		SetDayNumber(st.CurrentDay).
		SetNudgeStatus(nudge.NudgeStatusSent).
		SetNillableMessage(strPtr(message)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to send nudge: %w", err)
	}

	// Get sender info
	sender, _ := s.entClient.User.Get(ctx, userID)
	senderName := ""
	if sender != nil && sender.FirstName != nil {
		senderName = *sender.FirstName
	}

	return &dto.NudgeResponse{
		ID:         n.ID,
		StreakID:   n.StreakID,
		DayNumber:  n.DayNumber,
		Status:     string(n.NudgeStatus),
		Message:    ptrToString(n.Message),
		SenderID:   n.SenderUserID,
		SenderName: senderName,
		CreatedAt:  n.CreatedAt,
		SeenAt:     n.SeenAt,
	}, nil
}

// GetReceivedNudges returns nudges received by the user
func (s *StreakService) GetReceivedNudges(ctx context.Context, userID string) ([]*dto.NudgeResponse, error) {
	nudges, err := s.entClient.Nudge.
		Query().
		Where(nudge.ReceiverUserIDEQ(userID)).
		Where(nudge.NudgeStatusIn(nudge.NudgeStatusSent, nudge.NudgeStatusSeen)).
		Order(ent.Desc(nudge.FieldCreatedAt)).
		Limit(50).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get nudges: %w", err)
	}

	result := make([]*dto.NudgeResponse, len(nudges))
	for i, n := range nudges {
		sender, _ := s.entClient.User.Get(ctx, n.SenderUserID)
		senderName := ""
		if sender != nil && sender.FirstName != nil {
			senderName = *sender.FirstName
		}

		result[i] = &dto.NudgeResponse{
			ID:         n.ID,
			StreakID:   n.StreakID,
			DayNumber:  n.DayNumber,
			Status:     string(n.NudgeStatus),
			Message:    ptrToString(n.Message),
			SenderID:   n.SenderUserID,
			SenderName: senderName,
			CreatedAt:  n.CreatedAt,
			SeenAt:     n.SeenAt,
		}
	}

	return result, nil
}

// GetRecoveryOptions returns recovery options for a streak
func (s *StreakService) GetRecoveryOptions(ctx context.Context, userID, streakID string) (*dto.RecoveryOptionsResponse, error) {
	st, err := s.entClient.Streak.Get(ctx, streakID)
	if err != nil {
		return nil, fmt.Errorf("streak not found: %w", err)
	}

	// Check if in payment window
	if st.StreakState != streak.StreakStatePaymentWindow && st.StreakState != streak.StreakStateAtRisk {
		return &dto.RecoveryOptionsResponse{
			CanRecover: false,
		}, nil
	}

	// Get user credits
	u, _ := s.entClient.User.Get(ctx, userID)
	userCredits := 0
	if u != nil {
		userCredits = u.CreditBalance
	}

	// Recovery cost increases with day
	recoveryCost := 20 + (st.CurrentDay * 5)

	return &dto.RecoveryOptionsResponse{
		CanRecover:       true,
		RecoveryDeadline: st.RecoveryDeadlineAt,
		CreditCost:       recoveryCost,
		UserCredits:      userCredits,
	}, nil
}

// RecoverStreak recovers a broken streak using credits
func (s *StreakService) RecoverStreak(ctx context.Context, userID, streakID string, req *dto.RecoverStreakRequest) (*dto.RecoverStreakResponse, error) {
	st, err := s.entClient.Streak.Get(ctx, streakID)
	if err != nil {
		return nil, fmt.Errorf("streak not found: %w", err)
	}

	// Check if can recover
	if st.StreakState != streak.StreakStatePaymentWindow && st.StreakState != streak.StreakStateAtRisk {
		return nil, fmt.Errorf("streak cannot be recovered")
	}

	// Check deadline
	if st.RecoveryDeadlineAt != nil && time.Now().After(*st.RecoveryDeadlineAt) {
		return nil, fmt.Errorf("recovery deadline has passed")
	}

	// Get user
	u, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Calculate cost
	recoveryCost := 20 + (st.CurrentDay * 5)

	if req.PayWithCredits {
		if u.CreditBalance < recoveryCost {
			return nil, fmt.Errorf("insufficient credits")
		}

		// Deduct credits
		newBalance := u.CreditBalance - recoveryCost
		_, err = u.Update().
			SetCreditBalance(newBalance).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to deduct credits: %w", err)
		}

		// Restore streak
		_, err = st.Update().
			SetStreakState(streak.StreakStateActive).
			ClearRecoveryDeadlineAt().
			ClearBreakerUserID().
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to restore streak: %w", err)
		}

		return &dto.RecoverStreakResponse{
			Success:        true,
			NewCredits:     newBalance,
			StreakRestored: true,
			Message:        "Streak recovered successfully!",
		}, nil
	}

	return nil, fmt.Errorf("payment method not supported")
}

func (s *StreakService) streakToResponse(ctx context.Context, st *ent.Streak, userID string, conn *ent.Connection) (*dto.StreakResponse, error) {
	today := time.Now().Truncate(24 * time.Hour)

	// Check my check-in
	myCheckIn, _ := s.entClient.CheckIn.
		Query().
		Where(checkin.StreakIDEQ(st.ID)).
		Where(checkin.UserIDEQ(userID)).
		Where(checkin.DayNumberEQ(st.CurrentDay)).
		Where(checkin.CreatedAtGTE(today)).
		Exist(ctx)

	// Check partner check-in
	partnerID := conn.UserAID
	if conn.UserAID == userID {
		partnerID = conn.UserBID
	}
	partnerCheckIn, _ := s.entClient.CheckIn.
		Query().
		Where(checkin.StreakIDEQ(st.ID)).
		Where(checkin.UserIDEQ(partnerID)).
		Where(checkin.DayNumberEQ(st.CurrentDay)).
		Where(checkin.CreatedAtGTE(today)).
		Exist(ctx)

	// Get check-ins
	checkIns, _ := s.entClient.CheckIn.
		Query().
		Where(checkin.StreakIDEQ(st.ID)).
		Order(ent.Asc(checkin.FieldDayNumber)).
		All(ctx)

	checkInResponses := make([]dto.CheckInResponse, len(checkIns))
	for i, ci := range checkIns {
		activity := ""
		if ci.EventData != nil {
			if v, ok := ci.EventData["activity"].(string); ok {
				activity = v
			}
		}
		checkInResponses[i] = dto.CheckInResponse{
			ID:          ci.ID,
			DayNumber:   ci.DayNumber,
			UserID:      ci.UserID,
			CheckInType: string(ci.CheckInType),
			Activity:    activity,
			CreatedAt:   ci.CreatedAt,
		}
	}

	return &dto.StreakResponse{
		ID:                  st.ID,
		ConnectionID:        st.ConnectionID,
		State:               string(st.StreakState),
		CurrentDay:          st.CurrentDay,
		ResetCount:          st.ResetCount,
		MyCheckInToday:      myCheckIn,
		PartnerCheckInToday: partnerCheckIn,
		HealthScore:         st.StreakHealthScore,
		RecoveryDeadline:    st.RecoveryDeadlineAt,
		CheckIns:            checkInResponses,
		CreatedAt:           st.CreatedAt,
		UpdatedAt:           st.UpdatedAt,
		CompletedAt:         st.CompletedAt,
	}, nil
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
