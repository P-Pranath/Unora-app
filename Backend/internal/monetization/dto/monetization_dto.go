// internal/monetization/dto/monetization_dto.go
package dto

import "time"

// CreditPackageResponse represents a purchasable credit package
// @Description Credit package available for purchase
type CreditPackageResponse struct {
	ID              string  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name            string  `json:"name" example:"Starter Pack"`
	Description     string  `json:"description" example:"Get started with 50 credits"`
	CreditAmount    int     `json:"creditAmount" example:"50"`
	BonusCredits    int     `json:"bonusCredits" example:"5"`
	TotalCredits    int     `json:"totalCredits" example:"55"`
	PriceAmount     int     `json:"priceAmount" example:"9900"`
	PriceDisplay    string  `json:"priceDisplay" example:"₹99"`
	Currency        string  `json:"currency" example:"INR"`
	DiscountPercent float64 `json:"discountPercent" example:"10"`
	BadgeText       string  `json:"badgeText,omitempty" example:"Most Popular"`
	IsPopular       bool    `json:"isPopular" example:"true"`
}

// CreditBalanceResponse returns user's credit balance
// @Description User's credit balance
type CreditBalanceResponse struct {
	Balance           int `json:"balance" example:"150"`
	PendingCredits    int `json:"pendingCredits" example:"0"`
	LifetimeEarned    int `json:"lifetimeEarned" example:"200"`
	LifetimeSpent     int `json:"lifetimeSpent" example:"50"`
}

// CreditTransactionResponse represents a credit transaction
// @Description Credit transaction history entry
type CreditTransactionResponse struct {
	ID              string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	TransactionType string    `json:"transactionType" example:"purchase"`
	CreditAmount    int       `json:"creditAmount" example:"50"`
	BalanceAfter    int       `json:"balanceAfter" example:"150"`
	Description     string    `json:"description" example:"Purchased Starter Pack"`
	ReferenceType   string    `json:"referenceType,omitempty" example:"payment_order"`
	ReferenceID     string    `json:"referenceId,omitempty" example:"order_123"`
	CreatedAt       time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
}

// CreateOrderRequest is the request to create a Razorpay order
// @Description Create payment order for credit package
type CreateOrderRequest struct {
	PackageID string `json:"packageId" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// CreateOrderResponse returns the Razorpay order details
// @Description Razorpay order details for payment
type CreateOrderResponse struct {
	OrderID       string `json:"orderId" example:"order_ABC123"`
	RazorpayKey   string `json:"razorpayKey" example:"rzp_test_xxx"`
	Amount        int    `json:"amount" example:"9900"`
	Currency      string `json:"currency" example:"INR"`
	PackageName   string `json:"packageName" example:"Starter Pack"`
	CreditsToAdd  int    `json:"creditsToAdd" example:"55"`
	UserEmail     string `json:"userEmail" example:"user@example.com"`
	UserPhone     string `json:"userPhone,omitempty" example:"9876543210"`
}

// VerifyPaymentRequest is the request to verify payment
// @Description Verify Razorpay payment after success
type VerifyPaymentRequest struct {
	RazorpayOrderID   string `json:"razorpayOrderId" validate:"required" example:"order_ABC123"`
	RazorpayPaymentID string `json:"razorpayPaymentId" validate:"required" example:"pay_XYZ789"`
	RazorpaySignature string `json:"razorpaySignature" validate:"required" example:"signature_hash"`
}

// VerifyPaymentResponse returns the payment verification result
// @Description Payment verification result
type VerifyPaymentResponse struct {
	Success       bool   `json:"success" example:"true"`
	CreditsAdded  int    `json:"creditsAdded" example:"55"`
	NewBalance    int    `json:"newBalance" example:"155"`
	TransactionID string `json:"transactionId" example:"550e8400-e29b-41d4-a716-446655440000"`
	Message       string `json:"message" example:"Payment successful! 55 credits added."`
}

// PaymentOrderResponse represents a payment order
// @Description Payment order status
type PaymentOrderResponse struct {
	ID                string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	RazorpayOrderID   string     `json:"razorpayOrderId" example:"order_ABC123"`
	Amount            int        `json:"amount" example:"9900"`
	Currency          string     `json:"currency" example:"INR"`
	CreditsToAdd      int        `json:"creditsToAdd" example:"55"`
	Status            string     `json:"status" example:"paid"`
	CreatedAt         time.Time  `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	PaidAt            *time.Time `json:"paidAt,omitempty" example:"2024-01-01T00:05:00Z"`
}
