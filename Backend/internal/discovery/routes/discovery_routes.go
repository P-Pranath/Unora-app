// internal/discovery/routes/discovery_routes.go
package routes

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/discovery/handlers"
	"github.com/UnoraApp/be/internal/discovery/services"
	"github.com/UnoraApp/be/pkg/storage"
)

// RegisterDiscoveryRoutes registers all discovery and matching routes
func RegisterDiscoveryRoutes(
	router *gin.RouterGroup,
	entClient *ent.Client,
	storageClient storage.Client,
	authMiddleware gin.HandlerFunc,
) {
	// Create services
	discoveryService := services.NewDiscoveryService(entClient, storageClient)
	matchingService := services.NewMatchingService(entClient, storageClient)

	// Create handler
	handler := handlers.NewDiscoveryHandler(discoveryService, matchingService)

	// Public routes
	router.GET("/servers", handler.GetServers)

	// Protected routes (require authentication)
	protected := router.Group("")
	protected.Use(authMiddleware)
	{
		// Filters
		protected.GET("/servers/:serverType/filters", handler.GetFilter)
		protected.PUT("/servers/:serverType/filters", handler.UpdateFilter)

		// Discovery
		protected.GET("/servers/:serverType/discover", handler.GetDiscovery)
		protected.POST("/servers/:serverType/discover/refresh", handler.RefreshDiscovery)
		protected.GET("/discover/refresh-status", handler.GetRefreshStatus)

		// Interests
		protected.POST("/interests", handler.ExpressInterest)
		protected.GET("/interests/sent", handler.GetSentInterests)
		protected.GET("/interests/received", handler.GetReceivedInterests)

		// Connections
		protected.GET("/connections", handler.GetConnections)
		protected.GET("/connections/:id", handler.GetConnection)
		protected.DELETE("/connections/:id", handler.TerminateConnection)
	}
}
