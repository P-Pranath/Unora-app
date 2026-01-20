// internal/monetization/services/credits_service.go
package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/creditpackage"
	"github.com/UnoraApp/be/ent/generated/credittransaction"
	"github.com/UnoraApp/be/internal/monetization/dto"
)

// CreditsService handles credit-related business logic
type CreditsService struct {
	entClient *ent.Client
}

// NewCreditsService creates a new credits service
func NewCreditsService(entClient *ent.Client) *CreditsService {
	return &CreditsService{
		entClient: entClient,
	}
}

// GetPackages returns all active credit packages
func (s *CreditsService) GetPackages(ctx context.Context) ([]*dto.CreditPackageResponse, error) {
	packages, err := s.entClient.CreditPackage.
		Query().
		Where(creditpackage.IsActiveEQ(true)).
		Order(ent.Asc(creditpackage.FieldSortOrder)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get packages: %w", err)
	}

	result := make([]*dto.CreditPackageResponse, len(packages))
	for i, p := range packages {
		totalCredits := p.CreditAmount + p.BonusCredits
		priceDisplay := fmt.Sprintf("₹%d", p.PriceAmount/100) // Convert paise to rupees

		result[i] = &dto.CreditPackageResponse{
			ID:              p.ID,
			Name:            p.Name,
			Description:     ptrToString(p.Description),
			CreditAmount:    p.CreditAmount,
			BonusCredits:    p.BonusCredits,
			TotalCredits:    totalCredits,
			PriceAmount:     p.PriceAmount,
			PriceDisplay:    priceDisplay,
			Currency:        p.Currency,
			DiscountPercent: p.DiscountPercent,
			BadgeText:       ptrToString(p.BadgeText),
			IsPopular:       p.IsPopular,
		}
	}

	return result, nil
}

// GetBalance returns user's credit balance and stats
func (s *CreditsService) GetBalance(ctx context.Context, userID string) (*dto.CreditBalanceResponse, error) {
	user, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Calculate lifetime stats
	var lifetimeEarned, lifetimeSpent int

	// Get all transactions for stats
	transactions, err := s.entClient.CreditTransaction.
		Query().
		Where(credittransaction.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	for _, t := range transactions {
		if t.CreditAmount > 0 {
			lifetimeEarned += t.CreditAmount
		} else {
			lifetimeSpent += -t.CreditAmount
		}
	}

	return &dto.CreditBalanceResponse{
		Balance:        user.CreditBalance,
		PendingCredits: 0, // Could be calculated from pending orders
		LifetimeEarned: lifetimeEarned,
		LifetimeSpent:  lifetimeSpent,
	}, nil
}

// GetTransactionHistory returns user's credit transaction history
func (s *CreditsService) GetTransactionHistory(ctx context.Context, userID string, limit int) ([]*dto.CreditTransactionResponse, error) {
	if limit <= 0 {
		limit = 50
	}

	transactions, err := s.entClient.CreditTransaction.
		Query().
		Where(credittransaction.UserIDEQ(userID)).
		Order(ent.Desc(credittransaction.FieldCreatedAt)).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	result := make([]*dto.CreditTransactionResponse, len(transactions))
	for i, t := range transactions {
		result[i] = &dto.CreditTransactionResponse{
			ID:              t.ID,
			TransactionType: string(t.TransactionType),
			CreditAmount:    t.CreditAmount,
			BalanceAfter:    t.BalanceAfter,
			Description:     ptrToString(t.Description),
			ReferenceType:   ptrToString(t.ReferenceType),
			ReferenceID:     ptrToString(t.ReferenceID),
			CreatedAt:       t.CreatedAt,
		}
	}

	return result, nil
}

// AddCredits adds credits to user's balance and records transaction
func (s *CreditsService) AddCredits(ctx context.Context, userID string, amount int, txType credittransaction.TransactionType, description, refType, refID string) (*ent.CreditTransaction, error) {
	// Get user
	user, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Update balance
	newBalance := user.CreditBalance + amount
	_, err = user.Update().
		SetCreditBalance(newBalance).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update balance: %w", err)
	}

	// Create transaction record
	txID := uuid.New().String()
	tx, err := s.entClient.CreditTransaction.
		Create().
		SetID(txID).
		SetUserID(userID).
		SetTransactionType(txType).
		SetCreditAmount(amount).
		SetBalanceAfter(newBalance).
		SetNillableDescription(strPtr(description)).
		SetNillableReferenceType(strPtr(refType)).
		SetNillableReferenceID(strPtr(refID)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return tx, nil
}

// DeductCredits deducts credits from user's balance
func (s *CreditsService) DeductCredits(ctx context.Context, userID string, amount int, txType credittransaction.TransactionType, description, refType, refID string) (*ent.CreditTransaction, error) {
	return s.AddCredits(ctx, userID, -amount, txType, description, refType, refID)
}

// Helper functions
func ptrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
