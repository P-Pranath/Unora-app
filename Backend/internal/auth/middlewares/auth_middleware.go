// internal/auth/middlewares/auth_middleware.go
package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/internal/auth/services"
	"github.com/UnoraApp/be/pkg/apperror"
)

// AuthMiddleware validates the access token and injects user info into context
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			apperror.AbortWithError(c, apperror.Unauthorized("Missing authorization header"))
			return
		}

		// Check Bearer prefix
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			apperror.AbortWithError(c, apperror.Unauthorized("Invalid authorization header format"))
			return
		}

		token := parts[1]

		// Validate token
		payload, err := authService.ValidateAccessToken(token)
		if err != nil {
			apperror.AbortWithError(c, apperror.Unauthorized("Invalid or expired token"))
			return
		}

		// Set user info in context for handlers to use
		c.Set("userID", payload.UserID)
		c.Set("userEmail", payload.Email)
		c.Set("userProvider", payload.Provider)
		c.Set("userPayload", payload)

		c.Next()
	}
}

// OptionalAuthMiddleware validates token if present but doesn't require it
func OptionalAuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.Next()
			return
		}

		token := parts[1]
		payload, err := authService.ValidateAccessToken(token)
		if err != nil {
			c.Next()
			return
		}

		c.Set("userID", payload.UserID)
		c.Set("userEmail", payload.Email)
		c.Set("userProvider", payload.Provider)
		c.Set("userPayload", payload)

		c.Next()
	}
}
