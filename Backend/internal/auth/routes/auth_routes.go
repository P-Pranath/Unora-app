// internal/auth/routes/auth_routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/auth/middlewares"
	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/internal/di"
)

// RegisterAuthRoutes registers all authentication routes
// Uses Wire-generated dependency injection
func RegisterAuthRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
	redisClient *redis.Client,
	cfg *config.Config,
) {
	// Use Wire-generated injector to create handler with all dependencies
	authHandler := di.InitializeAuthHandler(entClient, redisClient, cfg)
	authService := di.InitializeAuthService(entClient, redisClient, cfg)

	// Public auth routes (no authentication required)
	authRoutes := router.Group("/auth")
	{
		// Google OAuth login - mobile app sends ID token
		authRoutes.POST("/google", authHandler.GoogleLogin)

		// Refresh token
		authRoutes.POST("/refresh", authHandler.RefreshToken)
	}

	// Protected auth routes (requires valid access token)
	protectedRoutes := router.Group("/auth")
	protectedRoutes.Use(middlewares.AuthMiddleware(authService))
	{
		// Logout - revokes refresh token
		protectedRoutes.POST("/logout", authHandler.Logout)

		// Get current user info from token
		protectedRoutes.GET("/me", authHandler.GetMe)
	}
}
