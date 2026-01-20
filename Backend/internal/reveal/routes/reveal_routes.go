// internal/reveal/routes/reveal_routes.go
package routes

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/reveal/handlers"
	"github.com/UnoraApp/be/internal/reveal/services"
)

// RegisterRevealRoutes registers all reveal-related routes
func RegisterRevealRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
	authMiddleware gin.HandlerFunc,
) {
	// Create services
	revealService := services.NewRevealService(entClient)

	// Create handler
	handler := handlers.NewRevealHandler(revealService)

	// Public routes
	router.GET("/reveal-milestones", handler.GetMilestones)

	// Protected routes
	protected := router.Group("")
	protected.Use(authMiddleware)
	{
		// Connection reveals
		protected.GET("/connections/:connectionId/reveals", handler.GetConnectionReveals)
		protected.POST("/connections/:connectionId/reveals/:milestoneId/unlock", handler.UnlockReveal)

		// Reveal actions
		protected.POST("/reveals/:revealId/viewed", handler.MarkRevealViewed)
	}
}
