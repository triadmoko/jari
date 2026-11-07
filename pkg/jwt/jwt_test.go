package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const secret = "test-secret-key"

func TestGenerateToken_Success(t *testing.T) {
	user := UserPayload{
		ID:       "user-123",
		Email:    "test@example.com",
		Username: "testuser",
	}

	token, err := GenerateToken(secret, user)

	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateTokenWithExpiry_Success(t *testing.T) {
	user := UserPayload{
		ID:       "user-123",
		Email:    "test@example.com",
		Username: "testuser",
	}
	expiry := 1 * time.Hour

	token, err := GenerateTokenWithExpiry(secret, user, expiry)

	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken_Success(t *testing.T) {
	secret := "test-secret-key"
	user := UserPayload{
		ID:       "user-123",
		Email:    "test@example.com",
		Username: "testuser",
	}

	token, err := GenerateToken(secret, user)
	require.NoError(t, err)

	claims, err := ValidateToken(secret, token)

	require.NoError(t, err)
	assert.Equal(t, user.ID, claims.UserID)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.Username, claims.Username)
}

func TestValidateToken_InvalidToken(t *testing.T) {
	secret := "test-secret-key"
	invalidToken := "invalid.token.here"

	claims, err := ValidateToken(secret, invalidToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateToken_WrongSecret(t *testing.T) {
	wrongSecret := "wrong-secret"
	user := UserPayload{
		ID:       "user-123",
		Email:    "test@example.com",
		Username: "testuser",
	}

	token, err := GenerateToken(secret, user)
	require.NoError(t, err)

	claims, err := ValidateToken(wrongSecret, token)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	secret := "test-secret-key"
	user := UserPayload{
		ID:       "user-123",
		Email:    "test@example.com",
		Username: "testuser",
	}

	token, err := GenerateTokenWithExpiry(secret, user, -1*time.Hour)
	require.NoError(t, err)

	claims, err := ValidateToken(secret, token)

	assert.Error(t, err)
	assert.Nil(t, claims)
}
