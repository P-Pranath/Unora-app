// internal/admin/middlewares/api_key_middleware.go
package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/UnoraApp/be/pkg/apperror"
)

const (
	// AdminAPIKeyHeader is the header name for admin API key
	AdminAPIKeyHeader = "X-Admin-API-Key"
)

// AdminAPIKeyMiddleware validates the admin API key
func AdminAPIKeyMiddleware() gin.HandlerFunc {
	adminAPIKey := os.Getenv("ADMIN_API_KEY")

	return func(c *gin.Context) {
		// Check if admin key is configured
		if adminAPIKey == "" {
			apperror.AbortWithError(c, apperror.Unavailable("Admin API is not configured"))
			return
		}

		// Get API key from header
		providedKey := c.GetHeader(AdminAPIKeyHeader)
		if providedKey == "" {
			apperror.AbortWithError(c, apperror.Unauthorized("Admin API key is required"))
			return
		}

		// Validate API key
		if providedKey != adminAPIKey {
			apperror.AbortWithError(c, apperror.Forbidden("Invalid admin API key"))
			return
		}

		// Set admin context
		c.Set("isAdmin", true)
		c.Next()
	}
}
