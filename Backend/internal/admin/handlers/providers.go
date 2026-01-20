// internal/admin/handlers/providers.go
package handlers

import "github.com/google/wire"

// ProviderSet is the wire provider set for admin handlers
var ProviderSet = wire.NewSet(
	NewAdminHandler,
)
