// internal/discovery/services/providers.go
package services

import "github.com/google/wire"

// ProviderSet is the wire provider set for discovery services
var ProviderSet = wire.NewSet(
	NewDiscoveryService,
	NewMatchingService,
)
