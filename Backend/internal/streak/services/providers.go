// internal/streak/services/providers.go
package services

import "github.com/google/wire"

// ProviderSet is the wire provider set for streak services
var ProviderSet = wire.NewSet(
	NewStreakService,
)
