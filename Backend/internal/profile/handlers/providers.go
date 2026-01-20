// internal/profile/handlers/providers.go
package handlers

import "github.com/google/wire"

// ProviderSet is the wire provider set for profile handlers
var ProviderSet = wire.NewSet(
	NewProfileHandler,
)
