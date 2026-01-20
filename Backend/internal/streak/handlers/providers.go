// internal/streak/handlers/providers.go
package handlers

import "github.com/google/wire"

// ProviderSet is the wire provider set for streak handlers
var ProviderSet = wire.NewSet(
	NewStreakHandler,
)
