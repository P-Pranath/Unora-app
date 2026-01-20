// internal/admin/services/connection_management_service.go
package services

import (
	"context"
	"fmt"
	"time"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/connection"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// ConnectionManagementService handles admin connection management
type ConnectionManagementService struct {
	entClient *ent.Client
}

// NewConnectionManagementService creates a new connection management service
func NewConnectionManagementService(entClient *ent.Client) *ConnectionManagementService {
	return &ConnectionManagementService{
		entClient: entClient,
	}
}

// ListConnections returns paginated list of connections
func (s *ConnectionManagementService) ListConnections(ctx context.Context, page, pageSize int, status, search string) (*dto.AdminConnectionListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	query := s.entClient.Connection.Query()

	if status != "" {
		query = query.Where(connection.ConnectionStatusEQ(connection.ConnectionStatus(status)))
	}

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count connections: %w", err)
	}

	connections, err := query.
		WithStreak().
		Order(ent.Desc(connection.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get connections: %w", err)
	}

	result := make([]dto.AdminConnectionResponse, len(connections))
	for i, c := range connections {
		userA, _ := s.entClient.User.Get(ctx, c.UserAID)
		userB, _ := s.entClient.User.Get(ctx, c.UserBID)

		userAName, userBName := "", ""
		if userA != nil && userA.FirstName != nil {
			userAName = *userA.FirstName
		}
		if userB != nil && userB.FirstName != nil {
			userBName = *userB.FirstName
		}

		streakDay, streakState := 0, ""
		if c.Edges.Streak != nil {
			streakDay = c.Edges.Streak.CurrentDay
			streakState = string(c.Edges.Streak.StreakState)
		}

		result[i] = dto.AdminConnectionResponse{
			ID:               c.ID,
			UserAID:          c.UserAID,
			UserAName:        userAName,
			UserBID:          c.UserBID,
			UserBName:        userBName,
			ServerType:       string(c.ServerType),
			ConnectionStatus: string(c.ConnectionStatus),
			StreakDay:        streakDay,
			StreakState:      streakState,
			CreatedAt:        c.CreatedAt,
			TerminatedAt:     c.TerminatedAt,
		}
	}

	totalPages := (total + pageSize - 1) / pageSize
	return &dto.AdminConnectionListResponse{
		Connections: result,
		Total:       total,
		Page:        page,
		PageSize:    pageSize,
		TotalPages:  totalPages,
	}, nil
}

// GetConnection returns connection details
func (s *ConnectionManagementService) GetConnection(ctx context.Context, id string) (*dto.AdminConnectionResponse, error) {
	c, err := s.entClient.Connection.
		Query().
		Where(connection.IDEQ(id)).
		WithStreak().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("connection not found: %w", err)
	}

	userA, _ := s.entClient.User.Get(ctx, c.UserAID)
	userB, _ := s.entClient.User.Get(ctx, c.UserBID)

	userAName, userBName := "", ""
	if userA != nil && userA.FirstName != nil {
		userAName = *userA.FirstName
	}
	if userB != nil && userB.FirstName != nil {
		userBName = *userB.FirstName
	}

	streakDay, streakState := 0, ""
	if c.Edges.Streak != nil {
		streakDay = c.Edges.Streak.CurrentDay
		streakState = string(c.Edges.Streak.StreakState)
	}

	return &dto.AdminConnectionResponse{
		ID:               c.ID,
		UserAID:          c.UserAID,
		UserAName:        userAName,
		UserBID:          c.UserBID,
		UserBName:        userBName,
		ServerType:       string(c.ServerType),
		ConnectionStatus: string(c.ConnectionStatus),
		StreakDay:        streakDay,
		StreakState:      streakState,
		CreatedAt:        c.CreatedAt,
		TerminatedAt:     c.TerminatedAt,
	}, nil
}

// TerminateConnection terminates a connection
func (s *ConnectionManagementService) TerminateConnection(ctx context.Context, id string, req *dto.TerminateConnectionRequest) error {
	c, err := s.entClient.Connection.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("connection not found: %w", err)
	}

	now := time.Now()
	_, err = c.Update().
		SetConnectionStatus(connection.ConnectionStatusTerminated).
		SetTerminatedAt(now).
		Save(ctx)
	return err
}
