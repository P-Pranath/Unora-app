// internal/admin/services/reveal_milestone_service.go
package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/revealmilestone"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// RevealMilestoneService handles admin reveal milestone management
type RevealMilestoneService struct {
	entClient *ent.Client
}

// NewRevealMilestoneService creates a new reveal milestone service
func NewRevealMilestoneService(entClient *ent.Client) *RevealMilestoneService {
	return &RevealMilestoneService{
		entClient: entClient,
	}
}

// ListRevealMilestones returns all reveal milestones
func (s *RevealMilestoneService) ListRevealMilestones(ctx context.Context, includeInactive bool) ([]*ent.RevealMilestone, error) {
	query := s.entClient.RevealMilestone.Query()
	if !includeInactive {
		query = query.Where(revealmilestone.IsActiveEQ(true))
	}
	return query.Order(ent.Asc(revealmilestone.FieldRevealNumber)).All(ctx)
}

// GetRevealMilestone returns a milestone by ID
func (s *RevealMilestoneService) GetRevealMilestone(ctx context.Context, id string) (*ent.RevealMilestone, error) {
	return s.entClient.RevealMilestone.Get(ctx, id)
}

// CreateRevealMilestone creates a new reveal milestone
func (s *RevealMilestoneService) CreateRevealMilestone(ctx context.Context, req *dto.CreateRevealMilestoneRequest) (*ent.RevealMilestone, error) {
	id := uuid.New().String()
	return s.entClient.RevealMilestone.
		Create().
		SetID(id).
		SetRevealNumber(req.RevealNumber).
		SetDayRequired(req.DayRequired).
		SetRevealType(revealmilestone.RevealType(req.RevealType)).
		SetTitle(req.Title).
		SetNillableDescription(strPtr(req.Description)).
		SetNillableIconName(strPtr(req.IconName)).
		SetCreditCost(req.CreditCost).
		SetIsActive(true).
		Save(ctx)
}

// UpdateRevealMilestone updates a reveal milestone
func (s *RevealMilestoneService) UpdateRevealMilestone(ctx context.Context, id string, req *dto.UpdateRevealMilestoneRequest) error {
	milestone, err := s.entClient.RevealMilestone.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("milestone not found: %w", err)
	}

	update := milestone.Update()
	if req.DayRequired != nil {
		update.SetDayRequired(*req.DayRequired)
	}
	if req.Title != nil {
		update.SetTitle(*req.Title)
	}
	if req.Description != nil {
		update.SetDescription(*req.Description)
	}
	if req.IconName != nil {
		update.SetIconName(*req.IconName)
	}
	if req.CreditCost != nil {
		update.SetCreditCost(*req.CreditCost)
	}
	if req.IsActive != nil {
		update.SetIsActive(*req.IsActive)
	}

	_, err = update.Save(ctx)
	return err
}

// DeleteRevealMilestone soft deletes a reveal milestone
func (s *RevealMilestoneService) DeleteRevealMilestone(ctx context.Context, id string) error {
	milestone, err := s.entClient.RevealMilestone.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("milestone not found: %w", err)
	}

	_, err = milestone.Update().SetIsActive(false).Save(ctx)
	return err
}
