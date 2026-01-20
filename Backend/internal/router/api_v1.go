// internal/router/api_v1.go
package router

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	ent "github.com/UnoraApp/be/ent/generated"
	adminroutes "github.com/UnoraApp/be/internal/admin/routes"
	"github.com/UnoraApp/be/internal/auth/middlewares"
	authroutes "github.com/UnoraApp/be/internal/auth/routes"
	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/internal/di"
	discoveryroutes "github.com/UnoraApp/be/internal/discovery/routes"
	"github.com/UnoraApp/be/internal/health"
	monetizationroutes "github.com/UnoraApp/be/internal/monetization/routes"
	monetizationservices "github.com/UnoraApp/be/internal/monetization/services"
	profileroutes "github.com/UnoraApp/be/internal/profile/routes"
	revealroutes "github.com/UnoraApp/be/internal/reveal/routes"
	safetyroutes "github.com/UnoraApp/be/internal/safety/routes"
	streakroutes "github.com/UnoraApp/be/internal/streak/routes"
	"github.com/UnoraApp/be/pkg/storage"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(router *gin.Engine, entClient *ent.Client, redisClient *redis.Client, cfg *config.Config, storageClient storage.Client) {
	ctx := context.Background()

	log.Println("Registering API routes...")

	// Health check endpoints
	router.GET("/health/liveness", gin.WrapF(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		health.LivenessHandler(w)
	})))

	router.GET("/health/readiness", gin.WrapF(health.ReadinessHandler(
		func() error {
			// Check database connection
			tx, err := entClient.Tx(ctx)
			if err != nil {
				return err
			}
			return tx.Rollback()
		},
		func() error {
			// Check Redis connection
			_, err := redisClient.Ping(ctx).Result()
			return err
		},
	)))

	// Group routes under /api/v1
	api := router.Group("/api/v1")

	// Root endpoint
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "API v1 is running",
			"version": "1.0.0",
		})
	})

	// ==========================================================
	// Register module routes
	// ==========================================================

	// Auth routes (Google OAuth, token refresh, logout)
	authroutes.RegisterAuthRoutes(api, entClient, redisClient, cfg)

	// Profile routes (profile, photos, hobbies)
	authService := di.InitializeAuthService(entClient, redisClient, cfg)
	authMiddleware := middlewares.AuthMiddleware(authService)
	profileroutes.RegisterProfileRoutes(api, entClient, storageClient, authService, authMiddleware)

	// Discovery and Matching routes (servers, discover, interests, connections)
	discoveryroutes.RegisterDiscoveryRoutes(api, entClient, storageClient, authMiddleware)

	// Streak routes (check-ins, nudges, recovery)
	streakroutes.RegisterStreakRoutes(api, entClient, storageClient, authMiddleware)

	// Reveal routes (milestones, unlock, content)
	revealroutes.RegisterRevealRoutes(api, entClient, authMiddleware)

	// Monetization routes (credits, payments, Razorpay)
	razorpayConfig := &monetizationservices.RazorpayConfig{
		KeyID:     os.Getenv("RAZORPAY_KEY_ID"),
		KeySecret: os.Getenv("RAZORPAY_KEY_SECRET"),
	}
	monetizationroutes.RegisterMonetizationRoutes(api, entClient, razorpayConfig, authMiddleware)

	// Safety routes (block, report)
	safetyroutes.RegisterSafetyRoutes(api, entClient, authMiddleware)

	// Admin routes (dashboard, user mgmt, reports, content)
	adminroutes.RegisterAdminRoutes(api, entClient)

	// Storage presigned URL endpoints (for frontend file uploads)
	storageGroup := api.Group("/storage")
	{
		// Get presigned upload URL
		storageGroup.POST("/presigned-upload", func(c *gin.Context) {
			var req struct {
				FileName    string `json:"fileName" binding:"required"`
				ContentType string `json:"contentType" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
				return
			}

			presignedURL, err := storageClient.GetPresignedUploadURL(c.Request.Context(), req.FileName, req.ContentType)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success":      true,
				"presignedUrl": presignedURL,
				"fileName":     req.FileName,
			})
		})

		// Get presigned download URL
		storageGroup.GET("/presigned-download/:fileName", func(c *gin.Context) {
			fileName := c.Param("fileName")

			presignedURL, err := storageClient.GetPresignedDownloadURL(c.Request.Context(), fileName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success":      true,
				"presignedUrl": presignedURL,
				"fileName":     fileName,
			})
		})

		// Get presigned preview URL
		storageGroup.GET("/presigned-preview/:fileName", func(c *gin.Context) {
			fileName := c.Param("fileName")

			presignedURL, err := storageClient.GetPresignedPreviewURL(c.Request.Context(), fileName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success":      true,
				"presignedUrl": presignedURL,
				"fileName":     fileName,
			})
		})
	}
}
