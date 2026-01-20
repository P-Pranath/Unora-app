// internal/reveal/services/reveal_service.go
package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/connection"
	"github.com/UnoraApp/be/ent/generated/reveal"
	"github.com/UnoraApp/be/ent/generated/revealmilestone"
	"github.com/UnoraApp/be/ent/generated/streak"
	"github.com/UnoraApp/be/internal/reveal/dto"
)

// RevealService handles reveal-related business logic
type RevealService struct {
	entClient *ent.Client
}

// NewRevealService creates a new reveal service
func NewRevealService(entClient *ent.Client) *RevealService {
	return &RevealService{
		entClient: entClient,
	}
}

// GetMilestones returns all reveal milestones
func (s *RevealService) GetMilestones(ctx context.Context) ([]*dto.RevealMilestoneResponse, error) {
	milestones, err := s.entClient.RevealMilestone.
		Query().
		Where(revealmilestone.IsActiveEQ(true)).
		Order(ent.Asc(revealmilestone.FieldRevealNumber)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get milestones: %w", err)
	}

	result := make([]*dto.RevealMilestoneResponse, len(milestones))
	for i, m := range milestones {
		result[i] = &dto.RevealMilestoneResponse{
			ID:           m.ID,
			RevealNumber: m.RevealNumber,
			DayRequired:  m.DayRequired,
			RevealType:   string(m.RevealType),
			Title:        m.Title,
			Description:  ptrToString(m.Description),
			IconName:     ptrToString(m.IconName),
			CreditCost:   m.CreditCost,
		}
	}

	return result, nil
}

// GetConnectionReveals returns all reveals for a connection
func (s *RevealService) GetConnectionReveals(ctx context.Context, userID, connectionID string) (*dto.ConnectionRevealsResponse, error) {
	// Verify user is part of connection
	conn, err := s.entClient.Connection.
		Query().
		Where(connection.IDEQ(connectionID)).
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		WithStreak().
		WithReveals(func(q *ent.RevealQuery) {
			q.WithMilestone().WithContent()
		}).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("connection not found")
		}
		return nil, fmt.Errorf("failed to get connection: %w", err)
	}

	// Get all milestones
	milestones, err := s.entClient.RevealMilestone.
		Query().
		Where(revealmilestone.IsActiveEQ(true)).
		Order(ent.Asc(revealmilestone.FieldRevealNumber)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get milestones: %w", err)
	}

	// Get current streak day
	currentDay := 0
	streakActive := false
	if conn.Edges.Streak != nil {
		currentDay = conn.Edges.Streak.CurrentDay
		streakActive = conn.Edges.Streak.StreakState == streak.StreakStateActive ||
			conn.Edges.Streak.StreakState == streak.StreakStateAtRisk
	}

	// Build reveal map from existing reveals
	revealMap := make(map[string]*ent.Reveal)
	for _, r := range conn.Edges.Reveals {
		revealMap[r.MilestoneID] = r
	}

	reveals := make([]dto.RevealResponse, len(milestones))
	var nextRevealDay *int

	for i, m := range milestones {
		existingReveal := revealMap[m.ID]

		revealResp := dto.RevealResponse{
			RevealNumber: m.RevealNumber,
			DayRequired:  m.DayRequired,
			RevealType:   string(m.RevealType),
			Title:        m.Title,
			CreditCost:   m.CreditCost,
		}

		if existingReveal != nil {
			revealResp.ID = existingReveal.ID
			revealResp.ConnectionID = existingReveal.ConnectionID
			revealResp.Status = string(existingReveal.RevealStatus)
			revealResp.UnlockMethod = string(existingReveal.UnlockMethod)
			revealResp.UnlockedAt = existingReveal.UnlockedAt
			revealResp.ViewedAt = existingReveal.ViewedAt
			revealResp.CanUnlock = false

			// Include content if unlocked
			if existingReveal.RevealStatus != reveal.RevealStatusLocked && existingReveal.Edges.Content != nil {
				content := existingReveal.Edges.Content
				revealResp.Content = &dto.RevealContentResponse{
					ID:                   content.ID,
					AISummary:            ptrToString(content.AiSummary),
					CompatibilityInsight: ptrToString(content.CompatibilityInsight),
					DimensionScores:      content.DimensionScores,
				}
				if content.ConversationStarters != nil {
					revealResp.Content.ConversationStarters = parseConversationStarters(*content.ConversationStarters)
				}
			}
		} else {
			revealResp.Status = "locked"
			revealResp.CanUnlock = currentDay >= m.DayRequired

			// Track next reveal day
			if currentDay < m.DayRequired && nextRevealDay == nil {
				nextRevealDay = &m.DayRequired
			}
		}

		reveals[i] = revealResp
	}

	return &dto.ConnectionRevealsResponse{
		ConnectionID:  connectionID,
		CurrentDay:    currentDay,
		StreakActive:  streakActive,
		Reveals:       reveals,
		NextRevealDay: nextRevealDay,
	}, nil
}

// UnlockReveal unlocks a reveal for a connection
func (s *RevealService) UnlockReveal(ctx context.Context, userID, connectionID, milestoneID string, req *dto.UnlockRevealRequest) (*dto.UnlockRevealResponse, error) {
	// Verify user is part of connection
	conn, err := s.entClient.Connection.
		Query().
		Where(connection.IDEQ(connectionID)).
		Where(
			connection.Or(
				connection.UserAIDEQ(userID),
				connection.UserBIDEQ(userID),
			),
		).
		WithStreak().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("connection not found: %w", err)
	}

	// Get milestone
	milestone, err := s.entClient.RevealMilestone.Get(ctx, milestoneID)
	if err != nil {
		return nil, fmt.Errorf("milestone not found: %w", err)
	}

	// Check if already unlocked
	existingReveal, _ := s.entClient.Reveal.
		Query().
		Where(reveal.ConnectionIDEQ(connectionID)).
		Where(reveal.MilestoneIDEQ(milestoneID)).
		Only(ctx)
	if existingReveal != nil && existingReveal.RevealStatus != reveal.RevealStatusLocked {
		return nil, fmt.Errorf("reveal already unlocked")
	}

	// Get current streak day
	currentDay := 0
	if conn.Edges.Streak != nil {
		currentDay = conn.Edges.Streak.CurrentDay
	}

	// Get user for credits
	user, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Determine unlock method
	unlockMethod := reveal.UnlockMethodEarned
	creditsUsed := 0

	if currentDay < milestone.DayRequired {
		// Need to purchase
		if !req.UseCredits {
			return nil, fmt.Errorf("not enough streak days to unlock, use credits to unlock early")
		}
		if user.CreditBalance < milestone.CreditCost {
			return nil, fmt.Errorf("insufficient credits")
		}
		unlockMethod = reveal.UnlockMethodPurchased
		creditsUsed = milestone.CreditCost

		// Deduct credits
		_, err = user.Update().
			SetCreditBalance(user.CreditBalance - creditsUsed).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to deduct credits: %w", err)
		}
	}

	// Create or update reveal
	now := time.Now()
	var r *ent.Reveal

	if existingReveal != nil {
		r, err = existingReveal.Update().
			SetRevealStatus(reveal.RevealStatusUnlocked).
			SetUnlockMethod(unlockMethod).
			SetUnlockedAt(now).
			Save(ctx)
	} else {
		revealID := uuid.New().String()
		r, err = s.entClient.Reveal.
			Create().
			SetID(revealID).
			SetConnectionID(connectionID).
			SetMilestoneID(milestoneID).
			SetUnlockMethod(unlockMethod).
			SetRevealStatus(reveal.RevealStatusUnlocked).
			SetUnlockedAt(now).
			Save(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to unlock reveal: %w", err)
	}

	// Generate content (placeholder - would use AI in production)
	contentID := uuid.New().String()
	content, err := s.entClient.RevealContent.
		Create().
		SetID(contentID).
		SetRevealID(r.ID).
		SetAiSummary("Based on your interactions, you both share a genuine appreciation for meaningful connections. Your communication styles complement each other beautifully.").
		SetCompatibilityInsight("You both value authenticity and deep conversations. This shared foundation suggests strong potential for a lasting connection.").
		SetConversationStarters("What's a hobby you've always wanted to try?\nWhat's the most memorable trip you've taken?\nWhat does your ideal weekend look like?").
		Save(ctx)
	if err != nil {
		// Log error but don't fail the unlock
	}

	// Get updated credits
	user, _ = s.entClient.User.Get(ctx, userID)

	revealResp := &dto.RevealResponse{
		ID:           r.ID,
		ConnectionID: r.ConnectionID,
		RevealNumber: milestone.RevealNumber,
		DayRequired:  milestone.DayRequired,
		RevealType:   string(milestone.RevealType),
		Title:        milestone.Title,
		Status:       string(r.RevealStatus),
		UnlockMethod: string(r.UnlockMethod),
		CanUnlock:    false,
		CreditCost:   milestone.CreditCost,
		UnlockedAt:   r.UnlockedAt,
	}

	if content != nil {
		revealResp.Content = &dto.RevealContentResponse{
			ID:                   content.ID,
			AISummary:            ptrToString(content.AiSummary),
			CompatibilityInsight: ptrToString(content.CompatibilityInsight),
			ConversationStarters: parseConversationStarters(ptrToString(content.ConversationStarters)),
		}
	}

	return &dto.UnlockRevealResponse{
		Success:          true,
		Reveal:           revealResp,
		CreditsUsed:      creditsUsed,
		RemainingCredits: user.CreditBalance,
		Message:          "Reveal unlocked! Check out your new insights.",
	}, nil
}

// MarkRevealViewed marks a reveal as viewed
func (s *RevealService) MarkRevealViewed(ctx context.Context, userID, revealID string) error {
	// Get reveal and verify access
	r, err := s.entClient.Reveal.
		Query().
		Where(reveal.IDEQ(revealID)).
		WithConnection().
		Only(ctx)
	if err != nil {
		return fmt.Errorf("reveal not found: %w", err)
	}

	// Verify user is part of connection
	conn := r.Edges.Connection
	if conn.UserAID != userID && conn.UserBID != userID {
		return fmt.Errorf("unauthorized")
	}

	if r.RevealStatus == reveal.RevealStatusLocked {
		return fmt.Errorf("reveal is locked")
	}

	if r.ViewedAt != nil {
		return nil // Already viewed
	}

	_, err = r.Update().
		SetRevealStatus(reveal.RevealStatusViewed).
		SetViewedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to mark as viewed: %w", err)
	}

	return nil
}

// Helper functions
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func parseConversationStarters(s string) []string {
	if s == "" {
		return []string{}
	}
	starters := strings.Split(s, "\n")
	result := make([]string, 0, len(starters))
	for _, starter := range starters {
		starter = strings.TrimSpace(starter)
		if starter != "" {
			result = append(result, starter)
		}
	}
	return result
}
