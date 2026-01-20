// internal/monetization/routes/monetization_routes.go
package routes

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/monetization/handlers"
	"github.com/UnoraApp/be/internal/monetization/services"
)

// RegisterMonetizationRoutes registers all monetization routes
func RegisterMonetizationRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
	razorpayConfig *services.RazorpayConfig,
	authMiddleware gin.HandlerFunc,
) {
	// Create services
	creditsService := services.NewCreditsService(entClient)
	razorpayClient := services.NewRazorpayClient(razorpayConfig)
	paymentService := services.NewPaymentService(entClient, razorpayClient, creditsService)

	// Create handler
	handler := handlers.NewMonetizationHandler(creditsService, paymentService)

	// Public routes
	router.GET("/credit-packages", handler.GetPackages)

	// Webhook (no auth, but Razorpay validates via signature)
	router.POST("/payments/webhook", handler.RazorpayWebhook)

	// Protected routes
	protected := router.Group("")
	protected.Use(authMiddleware)
	{
		// Credits
		protected.GET("/credits/balance", handler.GetBalance)
		protected.GET("/credits/transactions", handler.GetTransactionHistory)

		// Payments
		protected.POST("/payments/create-order", handler.CreateOrder)
		protected.POST("/payments/verify", handler.VerifyPayment)
		protected.GET("/payments/orders", handler.GetPaymentOrders)
	}
}
