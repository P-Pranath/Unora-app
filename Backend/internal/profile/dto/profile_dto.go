// internal/profile/dto/profile_dto.go
package dto

import "time"

// CreateProfileRequest is the request body for creating a user profile
// @Description Create profile request
type CreateProfileRequest struct {
	FirstName       string `json:"firstName" validate:"required,max=50" example:"John"`
	LastName        string `json:"lastName" validate:"max=50" example:"Doe"`
	DateOfBirth     string `json:"dateOfBirth" validate:"required" example:"1995-06-15"`
	Gender          string `json:"gender" validate:"required,oneof=male female non-binary prefer_not_to_say" example:"male"`
	City            string `json:"city" validate:"required,max=100" example:"Mumbai"`
	Bio             string `json:"bio" validate:"max=500" example:"Passionate about hiking and photography"`
	IntentStatement string `json:"intentStatement" validate:"max=200" example:"Looking for meaningful connections"`
}

// UpdateProfileRequest is the request body for updating a user profile
// @Description Update profile request
type UpdateProfileRequest struct {
	FirstName       *string `json:"firstName,omitempty" validate:"omitempty,max=50" example:"John"`
	LastName        *string `json:"lastName,omitempty" validate:"omitempty,max=50" example:"Doe"`
	DateOfBirth     *string `json:"dateOfBirth,omitempty" example:"1995-06-15"`
	Gender          *string `json:"gender,omitempty" validate:"omitempty,oneof=male female non-binary prefer_not_to_say" example:"male"`
	City            *string `json:"city,omitempty" validate:"omitempty,max=100" example:"Mumbai"`
	Bio             *string `json:"bio,omitempty" validate:"omitempty,max=500" example:"Updated bio"`
	IntentStatement *string `json:"intentStatement,omitempty" validate:"omitempty,max=200" example:"Updated intent"`
}

// ProfileResponse is the response body for profile data
// @Description User profile information
type ProfileResponse struct {
	ID              string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID          string     `json:"userId" example:"550e8400-e29b-41d4-a716-446655440001"`
	FirstName       string     `json:"firstName" example:"John"`
	LastName        string     `json:"lastName" example:"Doe"`
	DateOfBirth     *time.Time `json:"dateOfBirth" example:"1995-06-15T00:00:00Z"`
	Gender          string     `json:"gender" example:"male"`
	City            string     `json:"city" example:"Mumbai"`
	Bio             string     `json:"bio" example:"Passionate about hiking"`
	IntentStatement string     `json:"intentStatement" example:"Looking for meaningful connections"`
	CreatedAt       time.Time  `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UpdatedAt       time.Time  `json:"updatedAt" example:"2024-01-01T00:00:00Z"`
}

// OnboardingStatusResponse returns the onboarding completion status
// @Description Onboarding completion status for the user
type OnboardingStatusResponse struct {
	ProfileComplete bool `json:"profileComplete" example:"true"`
	PhotosComplete  bool `json:"photosComplete" example:"false"`
	PhotoCount      int  `json:"photoCount" example:"2"`
	PhotosRequired  int  `json:"photosRequired" example:"3"`
	HobbiesComplete bool `json:"hobbiesComplete" example:"true"`
	HobbyCount      int  `json:"hobbyCount" example:"3"`
	HobbiesRequired int  `json:"hobbiesRequired" example:"2"`
	IsComplete      bool `json:"isComplete" example:"false"`
}
