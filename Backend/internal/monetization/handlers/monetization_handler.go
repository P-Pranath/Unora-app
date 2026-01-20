// internal/monetization/handlers/monetization_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/monetization/dto"
	"github.com/UnoraApp/be/internal/monetization/services"
	"github.com/UnoraApp/be/pkg/apperror"
	"github.com/UnoraApp/be/pkg/response"
)

// MonetizationHandler handles monetization HTTP requests
type MonetizationHandler struct {
	creditsService *services.CreditsService
	paymentService *services.PaymentService
}

// NewMonetizationHandler creates a new monetization handler
func NewMonetizationHandler(creditsService *services.CreditsService, paymentService *services.PaymentService) *MonetizationHandler {
	return &MonetizationHandler{
		creditsService: creditsService,
		paymentService: paymentService,
	}
}

// ========== Credit Package Endpoints ==========

// GetPackages godoc
// @Summary      Get credit packages
// @Description  Get all available credit packages for purchase
// @Tags         credits
// @Accept       json
// @Produce      json
// @Success      200 {object} response.APIResponse{data=[]dto.CreditPackageResponse} "Credit packages"
// @Router       /credit-packages [get]
func (h *MonetizationHandler) GetPackages(c *gin.Context) {
	packages, err := h.creditsService.GetPackages(c.Request.Context())
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, packages)
}

// ========== Credit Balance Endpoints ==========

// GetBalance godoc
// @Summary      Get credit balance
// @Description  Get user's credit balance and lifetime stats
// @Tags         credits
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=dto.CreditBalanceResponse} "Credit balance"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /credits/balance [get]
func (h *MonetizationHandler) GetBalance(c *gin.Context) {
	userID, _ := c.Get("userID")

	balance, err := h.creditsService.GetBalance(c.Request.Context(), userID.(string))
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, balance)
}

// GetTransactionHistory godoc
// @Summary      Get transaction history
// @Description  Get user's credit transaction history
// @Tags         credits
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        limit query int false "Number of transactions (default 50)"
// @Success      200 {object} response.APIResponse{data=[]dto.CreditTransactionResponse} "Transactions"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /credits/transactions [get]
func (h *MonetizationHandler) GetTransactionHistory(c *gin.Context) {
	userID, _ := c.Get("userID")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	transactions, err := h.creditsService.GetTransactionHistory(c.Request.Context(), userID.(string), limit)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, transactions)
}

// ========== Payment Endpoints ==========

// CreateOrder godoc
// @Summary      Create payment order
// @Description  Create a Razorpay order for credit package purchase
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateOrderRequest true "Package to purchase"
// @Success      200 {object} response.APIResponse{data=dto.CreateOrderResponse} "Order details"
// @Failure      400 {object} response.APIResponse "Invalid request"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /payments/create-order [post]
func (h *MonetizationHandler) CreateOrder(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	order, err := h.paymentService.CreateOrder(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusOK, order)
}

// VerifyPayment godoc
// @Summary      Verify payment
// @Description  Verify Razorpay payment and add credits
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.VerifyPaymentRequest true "Payment verification data"
// @Success      200 {object} response.APIResponse{data=dto.VerifyPaymentResponse} "Verification result"
// @Failure      400 {object} response.APIResponse "Verification failed"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /payments/verify [post]
func (h *MonetizationHandler) VerifyPayment(c *gin.Context) {
	userID, _ := c.Get("userID")

	var req dto.VerifyPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	result, err := h.paymentService.VerifyPayment(c.Request.Context(), userID.(string), &req)
	if err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}
	response.JSON(c, http.StatusOK, result)
}

// GetPaymentOrders godoc
// @Summary      Get payment orders
// @Description  Get user's payment order history
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        limit query int false "Number of orders (default 20)"
// @Success      200 {object} response.APIResponse{data=[]dto.PaymentOrderResponse} "Payment orders"
// @Failure      401 {object} response.APIResponse "Not authenticated"
// @Router       /payments/orders [get]
func (h *MonetizationHandler) GetPaymentOrders(c *gin.Context) {
	userID, _ := c.Get("userID")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	orders, err := h.paymentService.GetPaymentOrders(c.Request.Context(), userID.(string), limit)
	if err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}
	response.JSON(c, http.StatusOK, orders)
}

// RazorpayWebhook godoc
// @Summary      Razorpay webhook
// @Description  Handle Razorpay webhook events
// @Tags         payments
// @Accept       json
// @Produce      json
// @Success      200 {object} response.APIResponse "Webhook processed"
// @Router       /payments/webhook [post]
func (h *MonetizationHandler) RazorpayWebhook(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		apperror.HandleError(c, apperror.BadRequest(err.Error()))
		return
	}

	event, _ := payload["event"].(string)
	payloadData, _ := payload["payload"].(map[string]interface{})

	if err := h.paymentService.HandleWebhook(c.Request.Context(), event, payloadData); err != nil {
		apperror.HandleError(c, apperror.InternalError(err))
		return
	}

	response.JSON(c, http.StatusOK, gin.H{"message": "Webhook processed"})
}
