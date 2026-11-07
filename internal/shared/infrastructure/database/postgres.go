package database

import (
	"app/internal/core/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgresDB represents a PostgreSQL database connection using GORM
type PostgresDB struct {
	DB *gorm.DB
}

// GormLogLevel picks the GORM logger verbosity. An explicit LOG_LEVEL wins:
// debug/trace log every SQL statement, info/warn log slow queries and errors,
// error logs errors only, fatal/panic are silent. Without LOG_LEVEL, development
// logs every statement and every other environment logs slow queries and errors.
func GormLogLevel(env, logLevel string) logger.LogLevel {
	switch logLevel {
	case "trace", "debug":
		return logger.Info
	case "info", "warn", "warning":
		return logger.Warn
	case "error":
		return logger.Error
	case "fatal", "panic":
		return logger.Silent
	}
	if env == config.EnvDevelopment {
		return logger.Info
	}
	return logger.Warn
}

// NewPostgresDB creates a new PostgreSQL database connection using GORM.
// gormLevel sets the verbosity of GORM's own logger (see GormLogLevel).
func NewPostgresDB(cfg config.DatabaseConfig, gormLevel logger.LogLevel) (*PostgresDB, error) {
	dsn := cfg.URL
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name, cfg.SSLMode)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Get underlying sql.DB for connection pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database with GORM")

	return &PostgresDB{DB: db}, nil
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Ping checks the database connection
func (p *PostgresDB) Ping() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// GetDB returns the GORM database instance
func (p *PostgresDB) GetDB() *gorm.DB {
	return p.DB
}
