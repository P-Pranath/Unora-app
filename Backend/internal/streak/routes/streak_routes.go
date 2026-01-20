// internal/streak/routes/streak_routes.go
package routes

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/streak/handlers"
	"github.com/UnoraApp/be/internal/streak/services"
	"github.com/UnoraApp/be/pkg/storage"
)

// RegisterStreakRoutes registers all streak-related routes
func RegisterStreakRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
	storageClient storage.Client,
	authMiddleware gin.HandlerFunc,
) {
	// Create services
	streakService := services.NewStreakService(entClient, storageClient)

	// Create handler
	handler := handlers.NewStreakHandler(streakService)

	// All streak routes require authentication
	protected := router.Group("")
	protected.Use(authMiddleware)
	{
		// Streak operations on connections
		protected.GET("/connections/:connectionId/streak", handler.GetStreak)
		protected.POST("/connections/:connectionId/streak/check-in", handler.CheckIn)
		protected.POST("/connections/:connectionId/nudge", handler.SendNudge)

		// Global streak views
		protected.GET("/streaks/today", handler.GetTodayStreaks)

		// Nudges
		protected.GET("/nudges/received", handler.GetReceivedNudges)

		// Recovery
		protected.GET("/streaks/:streakId/recovery-options", handler.GetRecoveryOptions)
		protected.POST("/streaks/:streakId/recover", handler.RecoverStreak)
	}
}
