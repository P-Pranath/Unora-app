// internal/router/router.go
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/pkg/storage"
)

// SetupRouter configures the Gin router with all routes and middleware
func SetupRouter(ent *ent.Client, redisClient *redis.Client, cfg *config.Config, storageClient storage.Client) *gin.Engine {
	// Set Gin mode
	if cfg.Server.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize router
	router := gin.New()

	// Add middleware
	SetupMiddlewares(router, cfg)

	// Register all API routes
	RegisterRoutes(router, ent, redisClient, cfg, storageClient)

	return router
}
