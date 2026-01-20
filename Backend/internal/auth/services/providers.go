// internal/auth/services/providers.go
package services

import (
	"github.com/google/wire"
)

// ProviderSet defines the Wire provider set for auth services
var ProviderSet = wire.NewSet(
	NewGoogleService,
	NewTokenService,
	NewAuthService,
)
