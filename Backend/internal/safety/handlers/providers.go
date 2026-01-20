// internal/safety/handlers/providers.go
package handlers

import "github.com/google/wire"

// ProviderSet is the wire provider set for safety handlers
var ProviderSet = wire.NewSet(
	NewSafetyHandler,
)
