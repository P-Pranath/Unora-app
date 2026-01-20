// internal/safety/routes/safety_routes.go
package routes

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/safety/handlers"
	"github.com/UnoraApp/be/internal/safety/services"
)

// RegisterSafetyRoutes registers all safety-related routes
func RegisterSafetyRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
	authMiddleware gin.HandlerFunc,
) {
	// Create services
	safetyService := services.NewSafetyService(entClient)

	// Create handler
	handler := handlers.NewSafetyHandler(safetyService)

	// All safety routes require authentication
	protected := router.Group("")
	protected.Use(authMiddleware)
	{
		// Block operations
		protected.POST("/blocks", handler.BlockUser)
		protected.POST("/blocks/unblock", handler.UnblockUser)
		protected.GET("/blocks", handler.GetBlockedUsers)
		protected.GET("/blocks/status/:userId", handler.GetBlockStatus)

		// Report operations
		protected.POST("/reports", handler.ReportUser)
	}
}
