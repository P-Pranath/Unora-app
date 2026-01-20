// internal/admin/services/streak_management_service.go
package services

import (
	"context"
	"fmt"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/streak"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// StreakManagementService handles admin streak management
type StreakManagementService struct {
	entClient *ent.Client
}

// NewStreakManagementService creates a new streak management service
func NewStreakManagementService(entClient *ent.Client) *StreakManagementService {
	return &StreakManagementService{
		entClient: entClient,
	}
}

// ListStreaks returns paginated list of streaks
func (s *StreakManagementService) ListStreaks(ctx context.Context, page, pageSize int, state string) (*dto.AdminStreakListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	query := s.entClient.Streak.Query()

	if state != "" {
		query = query.Where(streak.StreakStateEQ(streak.StreakState(state)))
	}

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count streaks: %w", err)
	}

	streaks, err := query.
		WithConnection().
		Order(ent.Desc(streak.FieldCurrentDay)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get streaks: %w", err)
	}

	result := make([]dto.AdminStreakResponse, len(streaks))
	for i, st := range streaks {
		userAName, userBName := "", ""
		if st.Edges.Connection != nil {
			userA, _ := s.entClient.User.Get(ctx, st.Edges.Connection.UserAID)
			userB, _ := s.entClient.User.Get(ctx, st.Edges.Connection.UserBID)
			if userA != nil && userA.FirstName != nil {
				userAName = *userA.FirstName
			}
			if userB != nil && userB.FirstName != nil {
				userBName = *userB.FirstName
			}
		}

		result[i] = dto.AdminStreakResponse{
			ID:            st.ID,
			ConnectionID:  st.ConnectionID,
			UserAName:     userAName,
			UserBName:     userBName,
			CurrentDay:    st.CurrentDay,
			StreakState:   string(st.StreakState),
			LastCheckInAt: &st.UpdatedAt,
			CreatedAt:     st.CreatedAt,
		}
	}

	totalPages := (total + pageSize - 1) / pageSize
	return &dto.AdminStreakListResponse{
		Streaks:    result,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetStreak returns streak details
func (s *StreakManagementService) GetStreak(ctx context.Context, id string) (*dto.AdminStreakResponse, error) {
	st, err := s.entClient.Streak.
		Query().
		Where(streak.IDEQ(id)).
		WithConnection().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("streak not found: %w", err)
	}

	userAName, userBName := "", ""
	if st.Edges.Connection != nil {
		userA, _ := s.entClient.User.Get(ctx, st.Edges.Connection.UserAID)
		userB, _ := s.entClient.User.Get(ctx, st.Edges.Connection.UserBID)
		if userA != nil && userA.FirstName != nil {
			userAName = *userA.FirstName
		}
		if userB != nil && userB.FirstName != nil {
			userBName = *userB.FirstName
		}
	}

	return &dto.AdminStreakResponse{
		ID:            st.ID,
		ConnectionID:  st.ConnectionID,
		UserAName:     userAName,
		UserBName:     userBName,
		CurrentDay:    st.CurrentDay,
		StreakState:   string(st.StreakState),
		LastCheckInAt: &st.UpdatedAt,
		CreatedAt:     st.CreatedAt,
	}, nil
}

// AdjustStreak adjusts a streak's day
func (s *StreakManagementService) AdjustStreak(ctx context.Context, id string, req *dto.AdjustStreakRequest) error {
	st, err := s.entClient.Streak.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("streak not found: %w", err)
	}

	newState := streak.StreakStateActive
	if req.NewDay >= 15 {
		newState = streak.StreakStateCompleted
	}

	_, err = st.Update().
		SetCurrentDay(req.NewDay).
		SetStreakState(newState).
		Save(ctx)
	return err
}

// ResetStreak resets a streak to day 0
func (s *StreakManagementService) ResetStreak(ctx context.Context, id string, req *dto.ResetStreakRequest) error {
	st, err := s.entClient.Streak.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("streak not found: %w", err)
	}

	_, err = st.Update().
		SetCurrentDay(0).
		SetStreakState(streak.StreakStateActive).
		Save(ctx)
	return err
}
