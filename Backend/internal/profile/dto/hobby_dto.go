// internal/profile/dto/hobby_dto.go
package dto

import "time"

// HobbyOptionResponse is the response for a hobby option (from master list)
// @Description Available hobby option
type HobbyOptionResponse struct {
	ID                string   `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name              string   `json:"name" example:"Photography"`
	Category          string   `json:"category" example:"Creative"`
	IconName          string   `json:"iconName" example:"camera"`
	MicroDescriptions []string `json:"microDescriptions" example:"[\"Amateur\",\"Professional\",\"Landscape\",\"Portrait\"]"`
}

// AddHobbyRequest is the request for adding a hobby to user's profile
// @Description Add hobby to profile
type AddHobbyRequest struct {
	HobbyOptionID    string `json:"hobbyOptionId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	MicroDescription string `json:"microDescription" validate:"max=200" example:"I love landscape photography"`
	DisplayOrder     int    `json:"displayOrder" validate:"min=0" example:"1"`
}

// UpdateHobbyRequest is the request for updating a hobby
// @Description Update hobby details
type UpdateHobbyRequest struct {
	MicroDescription *string `json:"microDescription,omitempty" validate:"omitempty,max=200" example:"Updated description"`
	DisplayOrder     *int    `json:"displayOrder,omitempty" validate:"omitempty,min=0" example:"2"`
}

// HobbyResponse is the response body for user's hobby
// @Description User's selected hobby
type HobbyResponse struct {
	ID               string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	HobbyOption      HobbyOptionResponse `json:"hobbyOption"`
	MicroDescription string    `json:"microDescription" example:"I love landscape photography"`
	DisplayOrder     int       `json:"displayOrder" example:"1"`
	CreatedAt        time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}
