// internal/profile/dto/photo_dto.go
package dto

import "time"

// UploadPhotoURLRequest is the request for getting a presigned upload URL
// @Description Request presigned URL for photo upload
type UploadPhotoURLRequest struct {
	ContentType string `json:"contentType" validate:"required" example:"image/jpeg"`
	IsFacePhoto bool   `json:"isFacePhoto" example:"true"`
}

// UploadPhotoURLResponse is the response with presigned upload URL
// @Description Presigned URL for uploading photo to storage
type UploadPhotoURLResponse struct {
	UploadURL   string `json:"uploadUrl" example:"https://storage.example.com/presigned-url..."`
	StorageKey  string `json:"storageKey" example:"photos/user-id/photo-1.jpg"`
	ExpiresIn   int    `json:"expiresIn" example:"1800"`
}

// RegisterPhotoRequest is the request for registering an uploaded photo
// @Description Register photo after successful upload
type RegisterPhotoRequest struct {
	StorageKey   string `json:"storageKey" validate:"required" example:"photos/user-id/photo-1.jpg"`
	DisplayOrder int    `json:"displayOrder" validate:"required,min=1,max=6" example:"1"`
	IsFacePhoto  bool   `json:"isFacePhoto" example:"true"`
}

// UpdatePhotoOrderRequest is the request for reordering photos
// @Description Update photo display order
type UpdatePhotoOrderRequest struct {
	DisplayOrder int `json:"displayOrder" validate:"required,min=1,max=6" example:"2"`
}

// PhotoResponse is the response body for photo data
// @Description User photo information
type PhotoResponse struct {
	ID                 string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StorageKey         string     `json:"storageKey" example:"photos/user-id/photo-1.jpg"`
	URL                string     `json:"url" example:"https://storage.example.com/photos/user-id/photo-1.jpg"`
	DisplayOrder       int        `json:"displayOrder" example:"1"`
	IsFacePhoto        bool       `json:"isFacePhoto" example:"true"`
	AIValidatedAt      *time.Time `json:"aiValidatedAt,omitempty" example:"2024-01-01T00:00:00Z"`
	AIValidationPassed *bool      `json:"aiValidationPassed,omitempty" example:"true"`
	CreatedAt          time.Time  `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}
