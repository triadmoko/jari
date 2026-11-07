package main

import (
	"fmt"
	"os"

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

	db, err := database.NewPostgresDB()
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
