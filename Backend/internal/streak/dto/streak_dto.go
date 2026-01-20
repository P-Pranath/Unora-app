// internal/streak/dto/streak_dto.go
package dto

import "time"

// StreakResponse represents detailed streak information
// @Description Streak details for a connection
type StreakResponse struct {
	ID                 string             `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ConnectionID       string             `json:"connectionId" example:"550e8400-e29b-41d4-a716-446655440000"`
	State              string             `json:"state" example:"active"`
	CurrentDay         int                `json:"currentDay" example:"5"`
	ResetCount         int                `json:"resetCount" example:"0"`
	MyCheckInToday     bool               `json:"myCheckInToday" example:"true"`
	PartnerCheckInToday bool              `json:"partnerCheckInToday" example:"false"`
	HealthScore        *float64           `json:"healthScore,omitempty" example:"0.85"`
	RecoveryDeadline   *time.Time         `json:"recoveryDeadline,omitempty"`
	CheckIns           []CheckInResponse  `json:"checkIns,omitempty"`
	CreatedAt          time.Time          `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UpdatedAt          time.Time          `json:"updatedAt" example:"2024-01-05T00:00:00Z"`
	CompletedAt        *time.Time         `json:"completedAt,omitempty"`
}

// CheckInRequest is the request for daily check-in
// @Description Daily check-in request
type CheckInRequest struct {
	Activity string `json:"activity,omitempty" validate:"max=200" example:"Had a great video call"`
}

// CheckInResponse represents a check-in record
// @Description Check-in record
type CheckInResponse struct {
	ID          string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	DayNumber   int       `json:"dayNumber" example:"5"`
	UserID      string    `json:"userId" example:"550e8400-e29b-41d4-a716-446655440000"`
	CheckInType string    `json:"checkInType" example:"manual"`
	Activity    string    `json:"activity,omitempty" example:"Had a video call"`
	CreatedAt   time.Time `json:"createdAt" example:"2024-01-05T00:00:00Z"`
}

// TodayStreaksResponse returns all streaks needing check-in today
// @Description Today's streaks requiring check-in
type TodayStreaksResponse struct {
	Streaks         []TodayStreakItem `json:"streaks"`
	TotalPending    int               `json:"totalPending" example:"2"`
	TotalCompleted  int               `json:"totalCompleted" example:"1"`
}

// TodayStreakItem represents a streak in the today view
// @Description Streak item for today's check-in list
type TodayStreakItem struct {
	StreakID        string `json:"streakId" example:"550e8400-e29b-41d4-a716-446655440000"`
	ConnectionID    string `json:"connectionId" example:"550e8400-e29b-41d4-a716-446655440000"`
	State           string `json:"state" example:"active"`
	CurrentDay      int    `json:"currentDay" example:"5"`
	PartnerName     string `json:"partnerName" example:"John"`
	PartnerPhotoURL string `json:"partnerPhotoUrl" example:"https://storage.example.com/photo.jpg"`
	NeedsCheckIn    bool   `json:"needsCheckIn" example:"true"`
	PartnerCheckedIn bool  `json:"partnerCheckedIn" example:"false"`
}

// SendNudgeRequest is the request for sending a nudge
// @Description Send nudge to streak partner
type SendNudgeRequest struct {
	Message string `json:"message,omitempty" validate:"max=200" example:"Hey! Don't forget to check in today 💪"`
}

// NudgeResponse represents a nudge
// @Description Nudge between streak partners
type NudgeResponse struct {
	ID         string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StreakID   string     `json:"streakId" example:"550e8400-e29b-41d4-a716-446655440000"`
	DayNumber  int        `json:"dayNumber" example:"5"`
	Status     string     `json:"status" example:"sent"`
	Message    string     `json:"message,omitempty" example:"Don't forget to check in!"`
	SenderID   string     `json:"senderId" example:"550e8400-e29b-41d4-a716-446655440000"`
	SenderName string     `json:"senderName" example:"John"`
	CreatedAt  time.Time  `json:"createdAt" example:"2024-01-05T10:00:00Z"`
	SeenAt     *time.Time `json:"seenAt,omitempty"`
}

// RecoveryOptionsResponse returns available recovery options
// @Description Streak recovery pricing options
type RecoveryOptionsResponse struct {
	CanRecover       bool       `json:"canRecover" example:"true"`
	RecoveryDeadline *time.Time `json:"recoveryDeadline,omitempty" example:"2024-01-06T00:00:00Z"`
	CreditCost       int        `json:"creditCost" example:"50"`
	UserCredits      int        `json:"userCredits" example:"100"`
}

// RecoverStreakRequest is the request for streak recovery
// @Description Request to recover a broken streak
type RecoverStreakRequest struct {
	PayWithCredits bool `json:"payWithCredits" example:"true"`
}

// RecoverStreakResponse is the response after recovery
// @Description Streak recovery result
type RecoverStreakResponse struct {
	Success         bool   `json:"success" example:"true"`
	NewCredits      int    `json:"newCredits" example:"50"`
	StreakRestored  bool   `json:"streakRestored" example:"true"`
	Message         string `json:"message" example:"Streak recovered successfully!"`
}
