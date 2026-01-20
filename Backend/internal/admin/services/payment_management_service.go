// internal/admin/services/payment_management_service.go
package services

import (
	"context"
	"fmt"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/ent/generated/paymentorder"
	"github.com/UnoraApp/be/internal/admin/dto"
)

// PaymentManagementService handles admin payment management
type PaymentManagementService struct {
	entClient *ent.Client
}

// NewPaymentManagementService creates a new payment management service
func NewPaymentManagementService(entClient *ent.Client) *PaymentManagementService {
	return &PaymentManagementService{
		entClient: entClient,
	}
}

// ListPaymentOrders returns paginated list of payment orders
func (s *PaymentManagementService) ListPaymentOrders(ctx context.Context, page, pageSize int, status string) (*dto.AdminPaymentListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	query := s.entClient.PaymentOrder.Query()

	if status != "" {
		query = query.Where(paymentorder.OrderStatusEQ(paymentorder.OrderStatus(status)))
	}

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count orders: %w", err)
	}

	orders, err := query.
		Order(ent.Desc(paymentorder.FieldCreatedAt)).
		Offset(offset).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}

	result := make([]dto.AdminPaymentOrderResponse, len(orders))
	for i, o := range orders {
		user, _ := s.entClient.User.Get(ctx, o.UserID)
		userEmail := ""
		if user != nil && user.Email != nil {
			userEmail = *user.Email
		}

		packageName := ""
		if o.PackageID != nil {
			pkg, _ := s.entClient.CreditPackage.Get(ctx, *o.PackageID)
			if pkg != nil {
				packageName = pkg.Name
			}
		}

		result[i] = dto.AdminPaymentOrderResponse{
			ID:                o.ID,
			UserID:            o.UserID,
			UserEmail:         userEmail,
			PackageID:         ptrToString(o.PackageID),
			PackageName:       packageName,
			RazorpayOrderID:   o.RazorpayOrderID,
			RazorpayPaymentID: ptrToString(o.RazorpayPaymentID),
			Amount:            o.Amount,
			Currency:          o.Currency,
			CreditsToAdd:      o.CreditsToAdd,
			OrderStatus:       string(o.OrderStatus),
			FailureReason:     ptrToString(o.FailureReason),
			CreatedAt:         o.CreatedAt,
			PaidAt:            o.PaidAt,
		}
	}

	totalPages := (total + pageSize - 1) / pageSize
	return &dto.AdminPaymentListResponse{
		Orders:     result,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// RefundPayment marks a payment as refunded (actual refund via Razorpay dashboard)
func (s *PaymentManagementService) RefundPayment(ctx context.Context, id string, req *dto.RefundPaymentRequest) error {
	order, err := s.entClient.PaymentOrder.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("order not found: %w", err)
	}

	if order.OrderStatus != paymentorder.OrderStatusPaid {
		return fmt.Errorf("can only refund paid orders")
	}

	// Mark as refunded
	_, err = order.Update().
		SetOrderStatus(paymentorder.OrderStatusRefunded).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}

	// Deduct credits from user (if they were added)
	user, _ := s.entClient.User.Get(ctx, order.UserID)
	if user != nil {
		newBalance := user.CreditBalance - order.CreditsToAdd
		if newBalance < 0 {
			newBalance = 0
		}
		user.Update().SetCreditBalance(newBalance).Save(ctx)
	}

	return nil
}
