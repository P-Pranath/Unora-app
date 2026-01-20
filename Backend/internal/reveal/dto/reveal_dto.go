// internal/reveal/dto/reveal_dto.go
package dto

import "time"

// RevealMilestoneResponse represents a reveal milestone (master data)
// @Description Reveal milestone configuration
type RevealMilestoneResponse struct {
	ID           string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	RevealNumber int    `json:"revealNumber" example:"1"`
	DayRequired  int    `json:"dayRequired" example:"5"`
	RevealType   string `json:"revealType" example:"personality"`
	Title        string `json:"title" example:"Personality Reveal"`
	Description  string `json:"description" example:"Discover your match's personality traits"`
	IconName     string `json:"iconName" example:"sparkles"`
	CreditCost   int    `json:"creditCost" example:"0"`
}

// RevealResponse represents a user's reveal for a connection
// @Description User's reveal status for a connection milestone
type RevealResponse struct {
	ID            string                   `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ConnectionID  string                   `json:"connectionId" example:"550e8400-e29b-41d4-a716-446655440000"`
	RevealNumber  int                      `json:"revealNumber" example:"1"`
	DayRequired   int                      `json:"dayRequired" example:"5"`
	RevealType    string                   `json:"revealType" example:"personality"`
	Title         string                   `json:"title" example:"Personality Reveal"`
	Status        string                   `json:"status" example:"unlocked"`
	UnlockMethod  string                   `json:"unlockMethod" example:"earned"`
	CanUnlock     bool                     `json:"canUnlock" example:"true"`
	CreditCost    int                      `json:"creditCost" example:"0"`
	Content       *RevealContentResponse   `json:"content,omitempty"`
	UnlockedAt    *time.Time               `json:"unlockedAt,omitempty" example:"2024-01-05T00:00:00Z"`
	ViewedAt      *time.Time               `json:"viewedAt,omitempty"`
}

// RevealContentResponse represents the AI-generated reveal content
// @Description AI-generated reveal content with insights
type RevealContentResponse struct {
	ID                    string                 `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	AISummary             string                 `json:"aiSummary" example:"Based on your conversations, you share a deep appreciation for..."`
	CompatibilityInsight  string                 `json:"compatibilityInsight" example:"Your communication styles complement each other well..."`
	ConversationStarters  []string               `json:"conversationStarters" example:"[\"What made you interested in photography?\"]"`
	DimensionScores       map[string]interface{} `json:"dimensionScores,omitempty"`
}

// ConnectionRevealsResponse represents all reveals for a connection
// @Description All reveals status for a connection
type ConnectionRevealsResponse struct {
	ConnectionID   string           `json:"connectionId" example:"550e8400-e29b-41d4-a716-446655440000"`
	CurrentDay     int              `json:"currentDay" example:"5"`
	StreakActive   bool             `json:"streakActive" example:"true"`
	Reveals        []RevealResponse `json:"reveals"`
	NextRevealDay  *int             `json:"nextRevealDay,omitempty" example:"10"`
}

// UnlockRevealRequest is the request to unlock a reveal
// @Description Request to unlock a reveal
type UnlockRevealRequest struct {
	UseCredits bool `json:"useCredits" example:"false"`
}

// UnlockRevealResponse is the response after unlocking
// @Description Reveal unlock result
type UnlockRevealResponse struct {
	Success         bool            `json:"success" example:"true"`
	Reveal          *RevealResponse `json:"reveal"`
	CreditsUsed     int             `json:"creditsUsed" example:"0"`
	RemainingCredits int            `json:"remainingCredits" example:"100"`
	Message         string          `json:"message" example:"Reveal unlocked! Check out your new insights."`
}
