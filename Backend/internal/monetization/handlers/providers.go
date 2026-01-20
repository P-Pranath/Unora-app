// internal/monetization/handlers/providers.go
package handlers

import "github.com/google/wire"

// ProviderSet is the wire provider set for monetization handlers
var ProviderSet = wire.NewSet(
	NewMonetizationHandler,
)
