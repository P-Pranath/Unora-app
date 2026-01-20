// internal/reveal/services/providers.go
package services

import "github.com/google/wire"

// ProviderSet is the wire provider set for reveal services
var ProviderSet = wire.NewSet(
	NewRevealService,
)
