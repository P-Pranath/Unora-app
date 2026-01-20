// internal/discovery/dto/discovery_dto.go
package dto

import "time"

// ServerResponse represents a server type (partner/friend/growth)
// @Description Server information for intent-based discovery
type ServerResponse struct {
	ID          string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ServerType  string `json:"serverType" example:"partner"`
	DisplayName string `json:"displayName" example:"Looking for Partner"`
	IconName    string `json:"iconName" example:"HeartStraight"`
	SortOrder   int    `json:"sortOrder" example:"1"`
}

// FilterConfig represents the filter settings for discovery
// @Description Filter configuration for discovery matching
type FilterConfig struct {
	AgeMin   *int     `json:"ageMin,omitempty" example:"21"`
	AgeMax   *int     `json:"ageMax,omitempty" example:"35"`
	Genders  []string `json:"genders,omitempty" example:"[\"male\", \"female\"]"`
	Cities   []string `json:"cities,omitempty" example:"[\"Mumbai\", \"Delhi\"]"`
	Hobbies  []string `json:"hobbies,omitempty" example:"[\"hobby-option-id-1\"]"`
	Distance *int     `json:"distance,omitempty" example:"50"`
}

// UpdateFilterRequest is the request for updating filter config
// @Description Update filter settings for a server
type UpdateFilterRequest struct {
	Config FilterConfig `json:"config" validate:"required"`
}

// FilterResponse represents user's filter for a server
// @Description User's filter settings for a server
type FilterResponse struct {
	ID         string       `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ServerType string       `json:"serverType" example:"partner"`
	Config     FilterConfig `json:"config"`
	IsPending  bool         `json:"isPending" example:"false"`
	UpdatedAt  time.Time    `json:"updatedAt" example:"2024-01-01T00:00:00Z"`
}

// DiscoveryCardResponse represents a single discovery card (a potential match)
// @Description Discovery card showing a potential match
type DiscoveryCardResponse struct {
	ID           string        `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	DisplayOrder int           `json:"displayOrder" example:"1"`
	User         CardUserInfo  `json:"user"`
	Photos       []CardPhoto   `json:"photos"`
	Hobbies      []CardHobby   `json:"hobbies"`
	CreatedAt    time.Time     `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// CardUserInfo represents basic user info on a discovery card
// @Description Basic user info displayed on discovery card
type CardUserInfo struct {
	ID              string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	FirstName       string `json:"firstName" example:"John"`
	Age             int    `json:"age" example:"28"`
	City            string `json:"city" example:"Mumbai"`
	Bio             string `json:"bio" example:"Love hiking and photography"`
	IntentStatement string `json:"intentStatement" example:"Looking for genuine connections"`
}

// CardPhoto represents a photo on a discovery card
// @Description Photo displayed on discovery card
type CardPhoto struct {
	URL          string `json:"url" example:"https://storage.example.com/photo.jpg"`
	DisplayOrder int    `json:"displayOrder" example:"1"`
	IsFacePhoto  bool   `json:"isFacePhoto" example:"true"`
}

// CardHobby represents a hobby on a discovery card
// @Description Hobby displayed on discovery card
type CardHobby struct {
	Name             string `json:"name" example:"Photography"`
	Category         string `json:"category" example:"Creative"`
	IconName         string `json:"iconName" example:"camera"`
	MicroDescription string `json:"microDescription" example:"Landscape photography"`
}

// DiscoveryBatchResponse represents a batch of discovery cards
// @Description Batch of 5 discovery cards
type DiscoveryBatchResponse struct {
	ID          string                  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ServerType  string                  `json:"serverType" example:"partner"`
	BatchStatus string                  `json:"batchStatus" example:"active"`
	Cards       []DiscoveryCardResponse `json:"cards"`
	ExpiresAt   *time.Time              `json:"expiresAt,omitempty" example:"2024-01-02T00:00:00Z"`
	CreatedAt   time.Time               `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// RefreshStatusResponse returns the refresh availability status
// @Description Status of discovery refresh availability
type RefreshStatusResponse struct {
	CanRefresh         bool       `json:"canRefresh" example:"true"`
	RefreshAvailableAt *time.Time `json:"refreshAvailableAt,omitempty" example:"2024-01-01T12:00:00Z"`
	SecondsUntilRefresh *int      `json:"secondsUntilRefresh,omitempty" example:"3600"`
}
