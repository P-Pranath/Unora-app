// internal/profile/services/photo_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/photo"
	"github.com/UnoraApp/be/internal/profile/dto"
	"github.com/UnoraApp/be/pkg/storage"
)

// PhotoService handles photo-related business logic
type PhotoService struct {
	entClient     *ent.Client
	storageClient storage.Client
}

// NewPhotoService creates a new photo service
func NewPhotoService(entClient *ent.Client, storageClient storage.Client) *PhotoService {
	return &PhotoService{
		entClient:     entClient,
		storageClient: storageClient,
	}
}

// GetUploadURL returns a presigned URL for photo upload
func (s *PhotoService) GetUploadURL(ctx context.Context, userID string, req *dto.UploadPhotoURLRequest) (*dto.UploadPhotoURLResponse, error) {
	// Generate unique storage key
	photoID := uuid.New().String()
	storageKey := fmt.Sprintf("photos/%s/%s", userID, photoID)

	// Get presigned upload URL
	uploadURL, err := s.storageClient.GetPresignedUploadURL(ctx, storageKey, req.ContentType)
	if err != nil {
		return nil, fmt.Errorf("failed to get upload URL: %w", err)
	}

	return &dto.UploadPhotoURLResponse{
		UploadURL:  uploadURL,
		StorageKey: storageKey,
		ExpiresIn:  1800, // 30 minutes
	}, nil
}

// RegisterPhoto registers a photo after successful upload
func (s *PhotoService) RegisterPhoto(ctx context.Context, userID string, req *dto.RegisterPhotoRequest) (*dto.PhotoResponse, error) {
	// Count existing photos
	count, err := s.entClient.Photo.
		Query().
		Where(photo.UserIDEQ(userID)).
		Where(photo.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count photos: %w", err)
	}
	if count >= 6 {
		return nil, fmt.Errorf("maximum of 6 photos allowed")
	}

	// Create photo record
	photoID := uuid.New().String()
	p, err := s.entClient.Photo.
		Create().
		SetID(photoID).
		SetUserID(userID).
		SetStorageKey(req.StorageKey).
		SetDisplayOrder(req.DisplayOrder).
		SetIsFacePhoto(req.IsFacePhoto).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to register photo: %w", err)
	}

	return s.photoToResponse(ctx, p), nil
}

// GetUserPhotos returns all photos for a user
func (s *PhotoService) GetUserPhotos(ctx context.Context, userID string) ([]*dto.PhotoResponse, error) {
	photos, err := s.entClient.Photo.
		Query().
		Where(photo.UserIDEQ(userID)).
		Where(photo.DeletedAtIsNil()).
		Order(ent.Asc(photo.FieldDisplayOrder)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get photos: %w", err)
	}

	result := make([]*dto.PhotoResponse, len(photos))
	for i, p := range photos {
		result[i] = s.photoToResponse(ctx, p)
	}

	return result, nil
}

// UpdatePhotoOrder updates the display order of a photo
func (s *PhotoService) UpdatePhotoOrder(ctx context.Context, userID, photoID string, req *dto.UpdatePhotoOrderRequest) (*dto.PhotoResponse, error) {
	p, err := s.entClient.Photo.
		Query().
		Where(photo.IDEQ(photoID)).
		Where(photo.UserIDEQ(userID)).
		Where(photo.DeletedAtIsNil()).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("photo not found")
		}
		return nil, fmt.Errorf("failed to get photo: %w", err)
	}

	updated, err := p.Update().
		SetDisplayOrder(req.DisplayOrder).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update photo order: %w", err)
	}

	return s.photoToResponse(ctx, updated), nil
}

// DeletePhoto soft deletes a photo
func (s *PhotoService) DeletePhoto(ctx context.Context, userID, photoID string) error {
	result, err := s.entClient.Photo.
		Update().
		Where(photo.IDEQ(photoID)).
		Where(photo.UserIDEQ(userID)).
		Where(photo.DeletedAtIsNil()).
		SetDeletedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete photo: %w", err)
	}
	if result == 0 {
		return fmt.Errorf("photo not found")
	}
	return nil
}

func (s *PhotoService) photoToResponse(ctx context.Context, p *ent.Photo) *dto.PhotoResponse {
	// Get presigned download URL
	photoURL, _ := s.storageClient.GetPresignedDownloadURL(ctx, p.StorageKey)

	return &dto.PhotoResponse{
		ID:                 p.ID,
		StorageKey:         p.StorageKey,
		URL:                photoURL,
		DisplayOrder:       p.DisplayOrder,
		IsFacePhoto:        p.IsFacePhoto,
		AIValidatedAt:      p.AiValidatedAt,
		AIValidationPassed: p.AiValidationPassed,
		CreatedAt:          p.CreatedAt,
	}
}
