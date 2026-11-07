package main

import (
	"fmt"
	"os"

	"app/internal/core/config"
	"app/internal/shared/infrastructure/database"
	"app/internal/shared/infrastructure/database/migrator"
	"app/migrations"
)

// runMigrate implements `main migrate <command>` using the embedded migrations
// and the same database configuration as the API server.
func runMigrate(args []string) error {
	if err := migrator.ValidateArgs(args); err != nil {
		return err
	}

	cfg := config.Parse()
	db, err := database.NewPostgresDB(cfg.Database, database.GormLogLevel(cfg.Env, cfg.Log.Level))
	if err != nil {
		return err
	}
	sqlDB, err := db.GetDB().DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// migrator.Run takes ownership of sqlDB and closes it.
	return migrator.Run(sqlDB, migrations.FS, args, os.Stdout)
}
