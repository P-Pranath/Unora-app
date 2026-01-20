//go:build wireinject
// +build wireinject

// internal/di/wire.go
package di

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"

	ent "github.com/UnoraApp/be/ent/generated"
	authhandlers "github.com/UnoraApp/be/internal/auth/handlers"
	authservices "github.com/UnoraApp/be/internal/auth/services"
	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/pkg/storage"
)

// Container holds all initialized dependencies for the application
type Container struct {
	Config        *config.Config
	EntClient     *ent.Client
	RedisClient   *redis.Client
	StorageClient storage.Client

	// Auth
	AuthService *authservices.AuthService
	AuthHandler *authhandlers.AuthHandler

	// Add more modules here as you build them:
	// UserService *userservices.UserService
	// UserHandler *userhandlers.UserHandler
}

// AuthSet is the provider set for auth module
var AuthSet = wire.NewSet(
	authservices.ProviderSet,
	authhandlers.ProviderSet,
)

// ProviderSet is the complete application provider set
var ProviderSet = wire.NewSet(
	AuthSet,
	// Add more module sets here:
	// UserSet,
	// ProductSet,
)

// InitializeContainer creates a fully wired Container
func InitializeContainer(
	cfg *config.Config,
	entClient *ent.Client,
	redisClient *redis.Client,
	storageClient storage.Client,
) (*Container, error) {
	wire.Build(
		ProviderSet,
		wire.Struct(new(Container), "*"),
	)
	return nil, nil
}

// InitializeAuthHandler creates just the auth handler (useful for testing)
func InitializeAuthHandler(
	entClient *ent.Client,
	redisClient *redis.Client,
	cfg *config.Config,
) *authhandlers.AuthHandler {
	wire.Build(AuthSet)
	return nil
}

// InitializeAuthService creates just the auth service (useful for testing)
func InitializeAuthService(
	entClient *ent.Client,
	redisClient *redis.Client,
	cfg *config.Config,
) *authservices.AuthService {
	wire.Build(authservices.ProviderSet)
	return nil
}
