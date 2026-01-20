// pkg/logger/middleware.go
package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestLogger returns a Gin middleware for request logging
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate request ID
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)

		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)
		status := c.Writer.Status()

		// Get user ID if authenticated
		userID, _ := c.Get("userID")
		userIDStr, _ := userID.(string)

		// Build log event
		event := Info().
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", path).
			Int("status", status).
			Dur("latency", latency).
			Str("ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent())

		if raw != "" {
			event = event.Str("query", raw)
		}
		if userIDStr != "" {
			event = event.Str("user_id", userIDStr)
		}

		// Log based on status code
		if status >= 500 {
			event.Msg("Server error")
		} else if status >= 400 {
			event.Msg("Client error")
		} else {
			event.Msg("Request completed")
		}
	}
}

// Recovery returns a Gin middleware for panic recovery with logging
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				requestID, _ := c.Get("request_id")
				Error().
					Interface("panic", r).
					Str("request_id", requestID.(string)).
					Str("path", c.Request.URL.Path).
					Msg("Panic recovered")

				c.AbortWithStatusJSON(500, gin.H{
					"success": false,
					"error": gin.H{
						"code":    "INTERNAL_ERROR",
						"message": "An unexpected error occurred",
					},
				})
			}
		}()
		c.Next()
	}
}
