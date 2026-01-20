// internal/monetization/services/razorpay_client.go
package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// RazorpayClient handles Razorpay API interactions
type RazorpayClient struct {
	keyID     string
	keySecret string
	baseURL   string
	client    *http.Client
}

// RazorpayConfig holds Razorpay configuration
type RazorpayConfig struct {
	KeyID     string
	KeySecret string
}

// NewRazorpayClient creates a new Razorpay client
func NewRazorpayClient(cfg *RazorpayConfig) *RazorpayClient {
	return &RazorpayClient{
		keyID:     cfg.KeyID,
		keySecret: cfg.KeySecret,
		baseURL:   "https://api.razorpay.com/v1",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetKeyID returns the public key ID for frontend
func (c *RazorpayClient) GetKeyID() string {
	return c.keyID
}

// CreateOrderRequest is the request to Razorpay create order API
type CreateOrderInput struct {
	Amount         int               `json:"amount"`
	Currency       string            `json:"currency"`
	Receipt        string            `json:"receipt"`
	Notes          map[string]string `json:"notes,omitempty"`
	PartialPayment bool              `json:"partial_payment"`
}

// RazorpayOrder is the response from Razorpay create order API
type RazorpayOrder struct {
	ID            string `json:"id"`
	Entity        string `json:"entity"`
	Amount        int    `json:"amount"`
	AmountPaid    int    `json:"amount_paid"`
	AmountDue     int    `json:"amount_due"`
	Currency      string `json:"currency"`
	Receipt       string `json:"receipt"`
	Status        string `json:"status"`
	Attempts      int    `json:"attempts"`
	CreatedAt     int64  `json:"created_at"`
}

// CreateOrder creates a new Razorpay order
func (c *RazorpayClient) CreateOrder(input *CreateOrderInput) (*RazorpayOrder, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+"/orders", strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(c.keyID, c.keySecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("razorpay error: %s", string(respBody))
	}

	var order RazorpayOrder
	if err := json.Unmarshal(respBody, &order); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &order, nil
}

// VerifyPaymentSignature verifies the Razorpay payment signature
func (c *RazorpayClient) VerifyPaymentSignature(orderID, paymentID, signature string) bool {
	message := orderID + "|" + paymentID
	h := hmac.New(sha256.New, []byte(c.keySecret))
	h.Write([]byte(message))
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

// FetchPayment fetches payment details from Razorpay
func (c *RazorpayClient) FetchPayment(paymentID string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", c.baseURL+"/payments/"+paymentID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(c.keyID, c.keySecret)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("razorpay error: %s", string(respBody))
	}

	var payment map[string]interface{}
	if err := json.Unmarshal(respBody, &payment); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return payment, nil
}
