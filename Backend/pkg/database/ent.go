// pkg/database/ent.go
package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	ent "github.com/UnoraApp/be/ent/generated"
	"github.com/UnoraApp/be/internal/config"
)

// SetupEntClient creates a new Ent client connected to MySQL
func SetupEntClient(cfg *config.Config) (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("ent.Open error: %w", err)
	}

	return client, nil
}
