// internal/discovery/handlers/providers.go
package handlers

import "github.com/google/wire"

// ProviderSet is the wire provider set for discovery handlers
var ProviderSet = wire.NewSet(
	NewDiscoveryHandler,
)
