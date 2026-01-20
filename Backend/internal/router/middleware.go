// internal/router/middleware.go
package router

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/pkg/logger"
)

// SetupMiddlewares configures global middleware for the router
func SetupMiddlewares(router *gin.Engine, cfg *config.Config) {
	// Recovery middleware
	router.Use(gin.Recovery())

	// Logging middleware
	router.Use(requestLogger())

	// CORS middleware
	router.Use(corsMiddleware(cfg))
}

// requestLogger logs each request with request ID tracking
func requestLogger() gin.HandlerFunc {
	log := logger.GetLogger("http")

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Generate/get request ID
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		// Get user ID if authenticated
		userID, _ := c.Get("userID")
		userIDStr, _ := userID.(string)

		event := log.Info().
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", path).
			Int("status", status).
			Dur("latency", latency).
			Str("ip", c.ClientIP())

		if query != "" {
			event = event.Str("query", query)
		}
		if userIDStr != "" {
			event = event.Str("user_id", userIDStr)
		}

		// Log with status-based message
		if status >= 500 {
			event.Msg("Server error")
		} else if status >= 400 {
			event.Msg("Client error")
		} else {
			event.Msg("Request completed")
		}
	}
}

// generateRequestID creates a simple unique request ID
func generateRequestID() string {
	return strings.ReplaceAll(time.Now().Format("20060102150405.000000"), ".", "")
}

// corsMiddleware handles Cross-Origin Resource Sharing
func corsMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Parse allowed origins
		allowedOrigins := strings.Split(cfg.CORS.AllowedOrigins, ",")

		// Check if origin is allowed
		allowed := false
		for _, o := range allowedOrigins {
			o = strings.TrimSpace(o)
			if o == "*" || o == origin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		c.Header("Access-Control-Allow-Methods", cfg.CORS.AllowedMethods)
		c.Header("Access-Control-Allow-Headers", cfg.CORS.AllowedHeaders)
		c.Header("Access-Control-Expose-Headers", cfg.CORS.ExposedHeaders)
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
