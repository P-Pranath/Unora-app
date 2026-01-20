//go:build wireinject
// +build wireinject

// internal/auth/wire.go
package auth

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/auth/handlers"
	"github.com/UnoraApp/be/internal/auth/services"
	"github.com/UnoraApp/be/internal/config"
)

// ProviderSet is the wire provider set for auth module
var ProviderSet = wire.NewSet(
	// Services
	services.NewGoogleService,
	services.NewTokenService,
	services.NewAuthService,
	// Handlers
	handlers.NewAuthHandler,
)

// InitializeAuthHandler wires up the auth handler with all dependencies
func InitializeAuthHandler(
	entClient *ent.Client,
	redisClient *redis.Client,
	cfg *config.Config,
) *handlers.AuthHandler {
	wire.Build(ProviderSet)
	return nil
}

// InitializeAuthService wires up the auth service with all dependencies
func InitializeAuthService(
	entClient *ent.Client,
	redisClient *redis.Client,
	cfg *config.Config,
) *services.AuthService {
	wire.Build(
		services.NewGoogleService,
		services.NewTokenService,
		services.NewAuthService,
	)
	return nil
}
