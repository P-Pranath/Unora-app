// cmd/seeder/main.go
package main

import (
	"context"
	"log"

	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/pkg/database"
	"github.com/UnoraApp/be/pkg/logger"
)

func main() {
	logger.InitLogger("dev")
	logSeeder := logger.GetLogger("seeder")

	logSeeder.Info().Msg("Starting database seeder...")

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup database connection
	entClient, err := database.SetupEntClient(cfg)
	if err != nil {
		log.Fatalf("Failed to setup DB client: %v", err)
	}
	defer entClient.Close()

	ctx := context.Background()

	// Add your seeding logic here
	// Example:
	// err = seedUsers(ctx, entClient)
	// if err != nil {
	//     log.Fatalf("Failed to seed users: %v", err)
	// }
	_ = ctx // Remove this when you add seeding logic

	logSeeder.Info().Msg("Database seeding completed successfully")
}
