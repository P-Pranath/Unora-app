// internal/monetization/services/payment_service.go
package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/credittransaction"
	"github.com/UnoraApp/be/ent/generated/paymentorder"
	"github.com/UnoraApp/be/internal/monetization/dto"
)

// PaymentService handles payment-related business logic
type PaymentService struct {
	entClient       *ent.Client
	razorpay        *RazorpayClient
	creditsService  *CreditsService
}

// NewPaymentService creates a new payment service
func NewPaymentService(entClient *ent.Client, razorpay *RazorpayClient, creditsService *CreditsService) *PaymentService {
	return &PaymentService{
		entClient:      entClient,
		razorpay:       razorpay,
		creditsService: creditsService,
	}
}

// CreateOrder creates a Razorpay order for a credit package
func (s *PaymentService) CreateOrder(ctx context.Context, userID string, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	// Get user
	user, err := s.entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Get package
	pkg, err := s.entClient.CreditPackage.Get(ctx, req.PackageID)
	if err != nil {
		return nil, fmt.Errorf("package not found: %w", err)
	}

	if !pkg.IsActive {
		return nil, fmt.Errorf("package is not available")
	}

	// Create Razorpay order
	orderInput := &CreateOrderInput{
		Amount:         pkg.PriceAmount,
		Currency:       pkg.Currency,
		Receipt:        uuid.New().String(),
		PartialPayment: false,
		Notes: map[string]string{
			"user_id":    userID,
			"package_id": pkg.ID,
		},
	}

	razorpayOrder, err := s.razorpay.CreateOrder(orderInput)
	if err != nil {
		return nil, fmt.Errorf("failed to create razorpay order: %w", err)
	}

	// Store order in database
	orderID := uuid.New().String()
	totalCredits := pkg.CreditAmount + pkg.BonusCredits
	expiresAt := time.Now().Add(30 * time.Minute)

	_, err = s.entClient.PaymentOrder.
		Create().
		SetID(orderID).
		SetUserID(userID).
		SetPackageID(pkg.ID).
		SetRazorpayOrderID(razorpayOrder.ID).
		SetAmount(pkg.PriceAmount).
		SetCurrency(pkg.Currency).
		SetCreditsToAdd(totalCredits).
		SetOrderStatus(paymentorder.OrderStatusCreated).
		SetExpiresAt(expiresAt).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	// Get user contact info
	email := ""
	phone := ""
	if user.Email != nil {
		email = *user.Email
	}
	if user.PhoneNumber != nil {
		phone = *user.PhoneNumber
	}

	return &dto.CreateOrderResponse{
		OrderID:      razorpayOrder.ID,
		RazorpayKey:  s.razorpay.GetKeyID(),
		Amount:       pkg.PriceAmount,
		Currency:     pkg.Currency,
		PackageName:  pkg.Name,
		CreditsToAdd: totalCredits,
		UserEmail:    email,
		UserPhone:    phone,
	}, nil
}

// VerifyPayment verifies Razorpay payment and adds credits
func (s *PaymentService) VerifyPayment(ctx context.Context, userID string, req *dto.VerifyPaymentRequest) (*dto.VerifyPaymentResponse, error) {
	// Verify signature
	if !s.razorpay.VerifyPaymentSignature(req.RazorpayOrderID, req.RazorpayPaymentID, req.RazorpaySignature) {
		return nil, fmt.Errorf("invalid payment signature")
	}

	// Get payment order
	order, err := s.entClient.PaymentOrder.
		Query().
		Where(paymentorder.RazorpayOrderIDEQ(req.RazorpayOrderID)).
		Where(paymentorder.UserIDEQ(userID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	// Check if already processed
	if order.OrderStatus == paymentorder.OrderStatusPaid {
		return nil, fmt.Errorf("payment already processed")
	}

	// Check if expired
	if order.ExpiresAt != nil && time.Now().After(*order.ExpiresAt) {
		return nil, fmt.Errorf("order has expired")
	}

	// Update order status
	now := time.Now()
	_, err = order.Update().
		SetOrderStatus(paymentorder.OrderStatusPaid).
		SetRazorpayPaymentID(req.RazorpayPaymentID).
		SetRazorpaySignature(req.RazorpaySignature).
		SetPaidAt(now).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	// Add credits to user
	description := fmt.Sprintf("Purchased credits via Razorpay")
	tx, err := s.creditsService.AddCredits(
		ctx,
		userID,
		order.CreditsToAdd,
		credittransaction.TransactionTypePurchase,
		description,
		"payment_order",
		order.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add credits: %w", err)
	}

	return &dto.VerifyPaymentResponse{
		Success:       true,
		CreditsAdded:  order.CreditsToAdd,
		NewBalance:    tx.BalanceAfter,
		TransactionID: tx.ID,
		Message:       fmt.Sprintf("Payment successful! %d credits added.", order.CreditsToAdd),
	}, nil
}

// GetPaymentOrders returns user's payment order history
func (s *PaymentService) GetPaymentOrders(ctx context.Context, userID string, limit int) ([]*dto.PaymentOrderResponse, error) {
	if limit <= 0 {
		limit = 20
	}

	orders, err := s.entClient.PaymentOrder.
		Query().
		Where(paymentorder.UserIDEQ(userID)).
		Order(ent.Desc(paymentorder.FieldCreatedAt)).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}

	result := make([]*dto.PaymentOrderResponse, len(orders))
	for i, o := range orders {
		result[i] = &dto.PaymentOrderResponse{
			ID:              o.ID,
			RazorpayOrderID: o.RazorpayOrderID,
			Amount:          o.Amount,
			Currency:        o.Currency,
			CreditsToAdd:    o.CreditsToAdd,
			Status:          string(o.OrderStatus),
			CreatedAt:       o.CreatedAt,
			PaidAt:          o.PaidAt,
		}
	}

	return result, nil
}

// HandleWebhook handles Razorpay webhook events
func (s *PaymentService) HandleWebhook(ctx context.Context, event string, payload map[string]interface{}) error {
	switch event {
	case "payment.captured":
		// Payment was captured successfully
		// This is handled by VerifyPayment, but webhook provides backup
		return nil
	case "payment.failed":
		// Payment failed
		paymentEntity, ok := payload["payment"].(map[string]interface{})
		if !ok {
			return nil
		}
		orderID, ok := paymentEntity["order_id"].(string)
		if !ok {
			return nil
		}
		reason, _ := paymentEntity["error_description"].(string)

		_, err := s.entClient.PaymentOrder.
			Update().
			Where(paymentorder.RazorpayOrderIDEQ(orderID)).
			SetOrderStatus(paymentorder.OrderStatusFailed).
			SetFailureReason(reason).
			Save(ctx)
		return err
	default:
		return nil
	}
}
