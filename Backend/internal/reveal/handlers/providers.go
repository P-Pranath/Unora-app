// internal/reveal/handlers/providers.go
package handlers

import "github.com/google/wire"

// ProviderSet is the wire provider set for reveal handlers
var ProviderSet = wire.NewSet(
	NewRevealHandler,
)
