package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Supported values for the ENV variable.
const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
	EnvTest        = "test"
)

// developmentJWTSecret is only used when ENV=development and JWT_SECRET is empty,
// so local runs work out of the box. It is rejected in every other environment.
const developmentJWTSecret = "dev-only-insecure-jwt-secret"

// developmentCORSOrigins is the Vite dev server origin, used only when
// ENV=development and CORS_ALLOWED_ORIGINS is empty.
var developmentCORSOrigins = []string{"http://localhost:5173", "http://127.0.0.1:5173"}

// insecureJWTSecrets are placeholder values that must never sign tokens outside
// development (the old built-in default and the sample values from the repo).
var insecureJWTSecrets = map[string]struct{}{
	"your-secret-key": {},
	"your-super-secret-jwt-key-change-this-in-production": {},
	"change-me":          {},
	"changeme":           {},
	"secret":             {},
	developmentJWTSecret: {},
}

// Config holds all configuration for our application
type Config struct {
	Env      string
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	URL     string
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	SSLMode string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret string
}

// CORSConfig holds the cross-origin allowlist.
type CORSConfig struct {
	// AllowedOrigins lists the origins permitted to make cross-origin requests.
	// A "*" entry allows every origin (without credentials).
	// Empty means every cross-origin request is rejected (same-origin still works).
	AllowedOrigins []string
}

// IsDevelopment reports whether the app runs with ENV=development.
func (c *Config) IsDevelopment() bool { return c.Env == EnvDevelopment }

// IsProduction reports whether the app runs with ENV=production.
func (c *Config) IsProduction() bool { return c.Env == EnvProduction }

// Parse reads configuration from environment variables without validating it.
// Use Load for the server; Parse is for callers that only need a subset
// (e.g. the migrate subcommand, which has no JWT_SECRET).
func Parse() *Config {
	env := strings.ToLower(strings.TrimSpace(getEnv("ENV", EnvDevelopment)))

	return &Config{
		Env: env,
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
		},
		Database: DatabaseConfig{
			URL:     getEnv("DATABASE_URL", ""),
			Host:    getEnv("DB_HOST", "localhost"),
			Port:    getEnv("DB_PORT", "5432"),
			User:    getEnv("DB_USER", "postgres"),
			Pass:    getEnv("DB_PASS", "password"),
			Name:    getEnv("DB_NAME", "app"),
			SSLMode: getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", ""),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvList("CORS_ALLOWED_ORIGINS"),
		},
	}
}

// Load reads configuration from environment variables and validates it.
// It returns an error (instead of silently falling back to a weak default)
// when the configuration is unsafe, so the process can exit with a clear message.
func Load() (*Config, error) {
	cfg := Parse()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

// Validate checks that the configuration is safe to start with. In
// development an empty JWT_SECRET is replaced by a well-known dev secret.
func (c *Config) Validate() error {
	var errs []error

	switch c.Env {
	case EnvDevelopment, EnvProduction, EnvTest:
	default:
		errs = append(errs, fmt.Errorf("ENV %q is invalid: must be one of %s, %s, %s",
			c.Env, EnvDevelopment, EnvProduction, EnvTest))
	}

	if err := c.validateJWT(); err != nil {
		errs = append(errs, err)
	}

	if err := c.validateCORS(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (c *Config) validateCORS() error {
	if len(c.CORS.AllowedOrigins) == 0 && c.IsDevelopment() {
		c.CORS.AllowedOrigins = append([]string(nil), developmentCORSOrigins...)
	}

	return nil
}

func (c *Config) validateJWT() error {
	if c.IsDevelopment() {
		if c.JWT.Secret == "" {
			c.JWT.Secret = developmentJWTSecret
		}
		return nil
	}

	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET is required when ENV=%s (generate one with: openssl rand -hex 32)", c.Env)
	}
	if _, bad := insecureJWTSecrets[c.JWT.Secret]; bad {
		return fmt.Errorf("JWT_SECRET must not be a default/example value when ENV=%s (generate one with: openssl rand -hex 32)", c.Env)
	}
	return nil
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getEnvList reads a comma-separated environment variable, trimming spaces and
// dropping empty entries. It returns nil when the variable is unset or empty.
func getEnvList(key string) []string {
	var out []string
	for item := range strings.SplitSeq(os.Getenv(key), ",") {
		if item = strings.TrimSpace(item); item != "" {
			out = append(out, item)
		}
	}
	return out
}
