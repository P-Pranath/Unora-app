// internal/admin/services/server_management_service.go
package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/server"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// ServerManagementService handles admin server management
type ServerManagementService struct {
	entClient *ent.Client
}

// NewServerManagementService creates a new server management service
func NewServerManagementService(entClient *ent.Client) *ServerManagementService {
	return &ServerManagementService{
		entClient: entClient,
	}
}

// ListServers returns all servers
func (s *ServerManagementService) ListServers(ctx context.Context) ([]*ent.Server, error) {
	return s.entClient.Server.Query().
		Order(ent.Asc(server.FieldSortOrder)).
		All(ctx)
}

// GetServer returns a server by ID
func (s *ServerManagementService) GetServer(ctx context.Context, id string) (*ent.Server, error) {
	return s.entClient.Server.Get(ctx, id)
}

// CreateServer creates a new server
func (s *ServerManagementService) CreateServer(ctx context.Context, req *dto.CreateServerRequest) (*ent.Server, error) {
	id := uuid.New().String()
	return s.entClient.Server.
		Create().
		SetID(id).
		SetServerType(server.ServerType(req.ServerType)).
		SetDisplayName(req.DisplayName).
		SetIconName(req.IconName).
		SetSortOrder(req.SortOrder).
		Save(ctx)
}

// UpdateServer updates a server
func (s *ServerManagementService) UpdateServer(ctx context.Context, id string, req *dto.UpdateServerRequest) error {
	srv, err := s.entClient.Server.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("server not found: %w", err)
	}

	update := srv.Update()
	if req.DisplayName != nil {
		update.SetDisplayName(*req.DisplayName)
	}
	if req.IconName != nil {
		update.SetIconName(*req.IconName)
	}
	if req.SortOrder != nil {
		update.SetSortOrder(*req.SortOrder)
	}

	_, err = update.Save(ctx)
	return err
}

// DeleteServer deletes a server
func (s *ServerManagementService) DeleteServer(ctx context.Context, id string) error {
	return s.entClient.Server.DeleteOneID(id).Exec(ctx)
}
