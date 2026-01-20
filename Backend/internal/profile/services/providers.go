// internal/profile/services/providers.go
package services

import "github.com/google/wire"

// ProviderSet is the wire provider set for profile services
var ProviderSet = wire.NewSet(
	NewProfileService,
	NewPhotoService,
)
