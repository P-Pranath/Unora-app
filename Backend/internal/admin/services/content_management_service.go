// internal/admin/services/content_management_service.go
package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/creditpackage"
	"github.com/UnoraApp/be/ent/generated/hobbyoption"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// ContentManagementService handles admin content management
type ContentManagementService struct {
	entClient *ent.Client
}

// NewContentManagementService creates a new content management service
func NewContentManagementService(entClient *ent.Client) *ContentManagementService {
	return &ContentManagementService{
		entClient: entClient,
	}
}

// CreateHobbyOption creates a new hobby option
func (s *ContentManagementService) CreateHobbyOption(ctx context.Context, req *dto.CreateHobbyOptionRequest) (*ent.HobbyOption, error) {
	id := uuid.New().String()

	hobbyOption, err := s.entClient.HobbyOption.
		Create().
		SetID(id).
		SetCategory(req.Category).
		SetName(req.Name).
		SetNillableIconName(strPtr(req.IconName)).
		SetIsActive(true).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create hobby option: %w", err)
	}

	return hobbyOption, nil
}

// UpdateHobbyOption updates a hobby option
func (s *ContentManagementService) UpdateHobbyOption(ctx context.Context, id string, updates map[string]interface{}) error {
	ho, err := s.entClient.HobbyOption.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("hobby option not found: %w", err)
	}

	update := ho.Update()

	if name, ok := updates["name"].(string); ok && name != "" {
		update.SetName(name)
	}
	if category, ok := updates["category"].(string); ok && category != "" {
		update.SetCategory(category)
	}
	if iconName, ok := updates["iconName"].(string); ok {
		update.SetIconName(iconName)
	}
	if isActive, ok := updates["isActive"].(bool); ok {
		update.SetIsActive(isActive)
	}

	_, err = update.Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update hobby option: %w", err)
	}

	return nil
}

// DeleteHobbyOption soft deletes a hobby option
func (s *ContentManagementService) DeleteHobbyOption(ctx context.Context, id string) error {
	ho, err := s.entClient.HobbyOption.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("hobby option not found: %w", err)
	}

	_, err = ho.Update().SetIsActive(false).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete hobby option: %w", err)
	}

	return nil
}

// ListHobbyOptions returns all hobby options
func (s *ContentManagementService) ListHobbyOptions(ctx context.Context, includeInactive bool) ([]*ent.HobbyOption, error) {
	query := s.entClient.HobbyOption.Query()
	if !includeInactive {
		query = query.Where(hobbyoption.IsActiveEQ(true))
	}
	return query.Order(ent.Asc(hobbyoption.FieldCategory), ent.Asc(hobbyoption.FieldName)).All(ctx)
}

// UpdateCreditPackage updates a credit package
func (s *ContentManagementService) UpdateCreditPackage(ctx context.Context, id string, req *dto.UpdateCreditPackageRequest) error {
	pkg, err := s.entClient.CreditPackage.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("package not found: %w", err)
	}

	update := pkg.Update()

	if req.Name != "" {
		update.SetName(req.Name)
	}
	if req.Description != "" {
		update.SetDescription(req.Description)
	}
	if req.CreditAmount != nil {
		update.SetCreditAmount(*req.CreditAmount)
	}
	if req.BonusCredits != nil {
		update.SetBonusCredits(*req.BonusCredits)
	}
	if req.PriceAmount != nil {
		update.SetPriceAmount(*req.PriceAmount)
	}
	if req.DiscountPercent != nil {
		update.SetDiscountPercent(*req.DiscountPercent)
	}
	if req.BadgeText != nil {
		update.SetBadgeText(*req.BadgeText)
	}
	if req.IsPopular != nil {
		update.SetIsPopular(*req.IsPopular)
	}
	if req.IsActive != nil {
		update.SetIsActive(*req.IsActive)
	}
	if req.SortOrder != nil {
		update.SetSortOrder(*req.SortOrder)
	}

	_, err = update.Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update package: %w", err)
	}

	return nil
}

// ListCreditPackages returns all credit packages
func (s *ContentManagementService) ListCreditPackages(ctx context.Context, includeInactive bool) ([]*ent.CreditPackage, error) {
	query := s.entClient.CreditPackage.Query()
	if !includeInactive {
		query = query.Where(creditpackage.IsActiveEQ(true))
	}
	return query.Order(ent.Asc(creditpackage.FieldSortOrder)).All(ctx)
}
