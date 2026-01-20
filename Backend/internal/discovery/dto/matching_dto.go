// internal/discovery/dto/matching_dto.go
package dto

import "time"

// ExpressInterestRequest is the request for expressing interest in a user
// @Description Express interest in a discovery card
type ExpressInterestRequest struct {
	DiscoveryCardID string `json:"discoveryCardId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// InterestResponse represents an interest (sent or received)
// @Description Interest expression data
type InterestResponse struct {
	ID         string          `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ServerType string          `json:"serverType" example:"partner"`
	Status     string          `json:"status" example:"pending"`
	User       InterestUserInfo `json:"user"`
	CreatedAt  time.Time       `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	MatchedAt  *time.Time      `json:"matchedAt,omitempty" example:"2024-01-02T00:00:00Z"`
}

// InterestUserInfo represents the other user in an interest
// @Description User info for interest display
type InterestUserInfo struct {
	ID        string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	FirstName string `json:"firstName" example:"John"`
	Age       int    `json:"age" example:"28"`
	City      string `json:"city" example:"Mumbai"`
	PhotoURL  string `json:"photoUrl" example:"https://storage.example.com/photo.jpg"`
}

// MatchNotification is sent when an interest becomes mutual
// @Description Notification when mutual interest forms a match
type MatchNotification struct {
	ConnectionID string          `json:"connectionId" example:"550e8400-e29b-41d4-a716-446655440000"`
	ServerType   string          `json:"serverType" example:"partner"`
	MatchedUser  InterestUserInfo `json:"matchedUser"`
	MatchedAt    time.Time       `json:"matchedAt" example:"2024-01-01T00:00:00Z"`
}

// ConnectionResponse represents a connection between two users
// @Description Active connection between users
type ConnectionResponse struct {
	ID           string           `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ServerType   string           `json:"serverType" example:"partner"`
	Status       string           `json:"status" example:"active"`
	Partner      ConnectionPartner `json:"partner"`
	Streak       *StreakSummary   `json:"streak,omitempty"`
	CreatedAt    time.Time        `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// ConnectionPartner represents the other user in a connection
// @Description Partner info for connection display
type ConnectionPartner struct {
	ID        string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	FirstName string `json:"firstName" example:"John"`
	Age       int    `json:"age" example:"28"`
	City      string `json:"city" example:"Mumbai"`
	PhotoURL  string `json:"photoUrl" example:"https://storage.example.com/photo.jpg"`
	Bio       string `json:"bio" example:"Love hiking and photography"`
}

// StreakSummary represents a brief streak status for connection list
// @Description Brief streak status
type StreakSummary struct {
	ID         string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CurrentDay int    `json:"currentDay" example:"5"`
	State      string `json:"state" example:"active"`
	NeedsCheckIn bool `json:"needsCheckIn" example:"true"`
}

// ConnectionDetailResponse represents full connection details
// @Description Full connection details with streak and reveals
type ConnectionDetailResponse struct {
	ConnectionResponse
	StreakDetails *StreakDetailResponse `json:"streakDetails,omitempty"`
	RevealsSummary []RevealSummary      `json:"reveals,omitempty"`
}

// StreakDetailResponse represents full streak details
// @Description Full streak information
type StreakDetailResponse struct {
	ID                string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CurrentDay        int        `json:"currentDay" example:"5"`
	State             string     `json:"state" example:"active"`
	ResetCount        int        `json:"resetCount" example:"0"`
	MyCheckInToday    bool       `json:"myCheckInToday" example:"true"`
	PartnerCheckInToday bool     `json:"partnerCheckInToday" example:"false"`
	RecoveryDeadline  *time.Time `json:"recoveryDeadline,omitempty"`
	CreatedAt         time.Time  `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// RevealSummary represents a brief reveal status
// @Description Brief reveal milestone status
type RevealSummary struct {
	Number     int    `json:"number" example:"1"`
	Type       string `json:"type" example:"personality"`
	Status     string `json:"status" example:"locked"`
	DayRequired int   `json:"dayRequired" example:"5"`
}
