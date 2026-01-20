// pkg/database/migrate.go
package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"

	"github.com/UnoraApp/be/internal/config"
	"github.com/UnoraApp/be/pkg/logger"
)

// RunMigrations runs all pending Goose migrations
func RunMigrations(cfg *config.Config) error {
	log := logger.GetLogger("migration")

	db, err := NewMySQLDB(cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}
	migrationsDir := filepath.Join(wd, "migrations", "versions")
	log.Info().Str("path", migrationsDir).Msg("Using Goose migrations")

	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("failed to run goose migrations: %w", err)
	}

	log.Info().Msg("✅ Migrations applied successfully")
	return nil
}
