// internal/safety/services/providers.go
package services

import "github.com/google/wire"

// ProviderSet is the wire provider set for safety services
var ProviderSet = wire.NewSet(
	NewSafetyService,
)
