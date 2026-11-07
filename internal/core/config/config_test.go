package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// clearEnv resets every variable Parse reads so tests do not depend on the host.
func clearEnv(t *testing.T) {
	t.Helper()
	for _, k := range []string{
		"ENV", "SERVER_PORT", "SERVER_HOST", "DATABASE_URL",
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASS", "DB_NAME", "DB_SSLMODE", "JWT_SECRET",
		"CORS_ALLOWED_ORIGINS", "LOG_LEVEL",
	} {
		t.Setenv(k, "")
	}
}

func TestLoad_ProductionWithoutJWTSecretFails(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "production")

	cfg, err := Load()

	require.Error(t, err)
	assert.Nil(t, cfg)
	assert.Contains(t, err.Error(), "JWT_SECRET is required")
	assert.Contains(t, err.Error(), "ENV=production")
}

func TestLoad_RejectsDefaultAndExampleSecretsOutsideDevelopment(t *testing.T) {
	for _, env := range []string{EnvProduction, EnvTest} {
		for secret := range insecureJWTSecrets {
			t.Run(env+"/"+secret, func(t *testing.T) {
				clearEnv(t)
				t.Setenv("ENV", env)
				t.Setenv("JWT_SECRET", secret)

				_, err := Load()

				require.Error(t, err)
				assert.Contains(t, err.Error(), "default/example")
			})
		}
	}
}

func TestLoad_ProductionWithStrongSecret(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "production")
	t.Setenv("JWT_SECRET", "0f1e2d3c4b5a69788796a5b4c3d2e1f00f1e2d3c4b5a69788796a5b4c3d2e1f0")

	cfg, err := Load()

	require.NoError(t, err)
	assert.True(t, cfg.IsProduction())
	assert.Equal(t, "0f1e2d3c4b5a69788796a5b4c3d2e1f00f1e2d3c4b5a69788796a5b4c3d2e1f0", cfg.JWT.Secret)
}

func TestLoad_DevelopmentFallsBackToDevSecret(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "development")

	cfg, err := Load()

	require.NoError(t, err)
	assert.True(t, cfg.IsDevelopment())
	assert.Equal(t, developmentJWTSecret, cfg.JWT.Secret)
}

func TestLoad_DevelopmentKeepsExplicitSecret(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "development")
	t.Setenv("JWT_SECRET", "your-secret-key")

	cfg, err := Load()

	require.NoError(t, err)
	assert.Equal(t, "your-secret-key", cfg.JWT.Secret)
}

func TestLoad_EnvDefaultsToDevelopmentAndIsNormalized(t *testing.T) {
	clearEnv(t)
	cfg, err := Load()
	require.NoError(t, err)
	assert.Equal(t, EnvDevelopment, cfg.Env)

	t.Setenv("ENV", "  Test ")
	t.Setenv("JWT_SECRET", "a-long-random-test-secret-value")
	cfg, err = Load()
	require.NoError(t, err)
	assert.Equal(t, EnvTest, cfg.Env)
}

func TestLoad_InvalidEnvFails(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "staging")
	t.Setenv("JWT_SECRET", "a-long-random-test-secret-value")

	_, err := Load()

	require.Error(t, err)
	assert.Contains(t, err.Error(), `ENV "staging" is invalid`)
}

func TestLoad_ReportsAllProblemsAtOnce(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "staging")

	_, err := Load()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ENV")
	assert.Contains(t, err.Error(), "JWT_SECRET")
}

func TestParse_ReadsDBPassFromDB_PASS(t *testing.T) {
	clearEnv(t)
	t.Setenv("DB_PASS", "s3cret-from-env")
	t.Setenv("DB_PASSWORD", "ignored")

	cfg := Parse()

	assert.Equal(t, "s3cret-from-env", cfg.Database.Pass)
}

func TestParse_Defaults(t *testing.T) {
	clearEnv(t)

	cfg := Parse()

	assert.Equal(t, "8080", cfg.Server.Port)
	assert.Equal(t, "0.0.0.0", cfg.Server.Host)
	assert.Equal(t, "localhost", cfg.Database.Host)
	assert.Equal(t, "5432", cfg.Database.Port)
	assert.Equal(t, "disable", cfg.Database.SSLMode)
	assert.Empty(t, cfg.JWT.Secret, "Parse must not invent a JWT secret")
}

func TestParse_CORSAllowedOriginsIsTrimmedList(t *testing.T) {
	clearEnv(t)
	t.Setenv("CORS_ALLOWED_ORIGINS", " https://app.example.com , ,https://admin.example.com")

	cfg := Parse()

	assert.Equal(t, []string{"https://app.example.com", "https://admin.example.com"}, cfg.CORS.AllowedOrigins)
}

func TestLoad_DevelopmentDefaultsToViteOrigins(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "development")

	cfg, err := Load()

	require.NoError(t, err)
	assert.Equal(t, developmentCORSOrigins, cfg.CORS.AllowedOrigins)
}

func TestLoad_DevelopmentKeepsExplicitOrigins(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "development")
	t.Setenv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")

	cfg, err := Load()

	require.NoError(t, err)
	assert.Equal(t, []string{"http://localhost:3000"}, cfg.CORS.AllowedOrigins)
}

func TestLoad_ProductionWithoutOriginsRejectsCrossOrigin(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "production")
	t.Setenv("JWT_SECRET", "a-long-random-test-secret-value")

	cfg, err := Load()

	require.NoError(t, err)
	assert.Empty(t, cfg.CORS.AllowedOrigins)
}

func TestLoad_WildcardOriginIsAccepted(t *testing.T) {
	clearEnv(t)
	t.Setenv("ENV", "development")
	t.Setenv("CORS_ALLOWED_ORIGINS", "*")

	cfg, err := Load()

	require.NoError(t, err)
	assert.Equal(t, []string{"*"}, cfg.CORS.AllowedOrigins)
}

func TestLoad_LogLevel(t *testing.T) {
	t.Run("defaults to empty", func(t *testing.T) {
		clearEnv(t)
		cfg, err := Load()
		require.NoError(t, err)
		assert.Empty(t, cfg.Log.Level)
	})

	t.Run("is normalized", func(t *testing.T) {
		clearEnv(t)
		t.Setenv("LOG_LEVEL", " DEBUG ")
		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, "debug", cfg.Log.Level)
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		clearEnv(t)
		t.Setenv("LOG_LEVEL", "verbose")
		cfg, err := Load()
		require.Error(t, err)
		assert.Nil(t, cfg)
		assert.Contains(t, err.Error(), "LOG_LEVEL")
	})
}

func TestJSONLogs(t *testing.T) {
	assert.True(t, (&Config{Env: EnvProduction}).JSONLogs())
	assert.False(t, (&Config{Env: EnvDevelopment}).JSONLogs())
	assert.False(t, (&Config{Env: EnvTest}).JSONLogs())
}
