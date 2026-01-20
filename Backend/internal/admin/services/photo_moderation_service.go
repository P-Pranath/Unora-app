// internal/admin/services/photo_moderation_service.go
package services

import (
	"context"
	"fmt"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/photo"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// PhotoModerationService handles admin photo moderation
type PhotoModerationService struct {
	entClient *ent.Client
}

// NewPhotoModerationService creates a new photo moderation service
func NewPhotoModerationService(entClient *ent.Client) *PhotoModerationService {
	return &PhotoModerationService{
		entClient: entClient,
	}
}

// ListPendingPhotos returns photos needing moderation
func (s *PhotoModerationService) ListPendingPhotos(ctx context.Context, page, pageSize int) (*dto.AdminPhotoListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	// Get unverified photos
	query := s.entClient.Photo.Query().
		Where(photo.IsVerified(false))

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count photos: %w", err)
	}

	photos, err := query.
		Order(ent.Desc(photo.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get photos: %w", err)
	}

	result := make([]dto.AdminPhotoResponse, len(photos))
	for i, p := range photos {
		user, _ := s.entClient.User.Get(ctx, p.UserID)
		userName := ""
		if user != nil && user.FirstName != nil {
			userName = *user.FirstName
		}

		result[i] = dto.AdminPhotoResponse{
			ID:         p.ID,
			UserID:     p.UserID,
			UserName:   userName,
			PhotoURL:   p.StorageKey, // Use storage_key as the URL/path
			SlotNumber: p.DisplayOrder,
			IsVerified: p.IsVerified,
			IsFlagged:  p.IsFlagged,
			FlagReason: ptrToString(p.FlagReason),
			UploadedAt: p.CreatedAt,
		}
	}

	totalPages := (total + pageSize - 1) / pageSize
	return &dto.AdminPhotoListResponse{
		Photos:     result,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// ListFlaggedPhotos returns flagged photos
func (s *PhotoModerationService) ListFlaggedPhotos(ctx context.Context, page, pageSize int) (*dto.AdminPhotoListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	query := s.entClient.Photo.Query().
		Where(photo.IsFlagged(true))

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count photos: %w", err)
	}

	photos, err := query.
		Order(ent.Desc(photo.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get photos: %w", err)
	}

	result := make([]dto.AdminPhotoResponse, len(photos))
	for i, p := range photos {
		user, _ := s.entClient.User.Get(ctx, p.UserID)
		userName := ""
		if user != nil && user.FirstName != nil {
			userName = *user.FirstName
		}

		result[i] = dto.AdminPhotoResponse{
			ID:         p.ID,
			UserID:     p.UserID,
			UserName:   userName,
			PhotoURL:   p.StorageKey,
			SlotNumber: p.DisplayOrder,
			IsVerified: p.IsVerified,
			IsFlagged:  p.IsFlagged,
			FlagReason: ptrToString(p.FlagReason),
			UploadedAt: p.CreatedAt,
		}
	}

	totalPages := (total + pageSize - 1) / pageSize
	return &dto.AdminPhotoListResponse{
		Photos:     result,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// ApprovePhoto approves a photo
func (s *PhotoModerationService) ApprovePhoto(ctx context.Context, id string) error {
	p, err := s.entClient.Photo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("photo not found: %w", err)
	}

	_, err = p.Update().
		SetIsVerified(true).
		SetIsFlagged(false).
		ClearFlagReason().
		Save(ctx)
	return err
}

// RejectPhoto rejects and deletes a photo
func (s *PhotoModerationService) RejectPhoto(ctx context.Context, id string, req *dto.RejectPhotoRequest) error {
	p, err := s.entClient.Photo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("photo not found: %w", err)
	}

	// Mark as flagged then delete
	p.Update().SetIsFlagged(true).SetFlagReason(req.Reason).Save(ctx)

	return s.entClient.Photo.DeleteOneID(id).Exec(ctx)
}

// FlagPhoto flags a photo for review
func (s *PhotoModerationService) FlagPhoto(ctx context.Context, id, reason string) error {
	p, err := s.entClient.Photo.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("photo not found: %w", err)
	}

	_, err = p.Update().
		SetIsFlagged(true).
		SetFlagReason(reason).
		Save(ctx)
	return err
}
