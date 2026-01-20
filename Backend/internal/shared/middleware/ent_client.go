// internal/shared/middleware/ent_client.go
package middleware

import (
	"github.com/gin-gonic/gin"

	ent "github.com/UnoraApp/be/ent/generated"
)

// EntClientKey is the context key for the Ent client
const EntClientKey = "entClient"

// EntClient middleware injects the Ent client into the request context
func EntClient(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(EntClientKey, client)
		c.Next()
	}
}

// GetEntClient retrieves the Ent client from the context
func GetEntClient(c *gin.Context) *ent.Client {
	if client, exists := c.Get(EntClientKey); exists {
		return client.(*ent.Client)
	}
	return nil
}
