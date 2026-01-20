// internal/auth/handlers/providers.go
package handlers

import (
	"github.com/google/wire"
)

// ProviderSet defines the Wire provider set for auth handlers
var ProviderSet = wire.NewSet(
	NewAuthHandler,
)
