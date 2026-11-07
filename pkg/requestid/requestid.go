// Package requestid carries a per-request correlation ID through context.Context
// so that middleware, use cases, and the logger can share it without importing
// each other.
package requestid

import (
	"context"

	"github.com/google/uuid"
)

// Header is the HTTP header used to receive and echo the request ID.
const Header = "X-Request-ID"

const maxLen = 128

type ctxKey struct{}

func New() string { return uuid.NewString() }

func NewContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKey{}, id)
}

// FromContext returns the request ID stored in ctx, or "" when there is none.
func FromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	id, _ := ctx.Value(ctxKey{}).(string)
	return id
}

// Valid reports whether id is safe to accept from a client and write to logs
// and response headers: 1-128 characters of [A-Za-z0-9._-] only. Anything else
// (control characters, spaces, very long values) could forge log lines.
func Valid(id string) bool {
	if id == "" || len(id) > maxLen {
		return false
	}
	for _, r := range id {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
		case r == '-', r == '_', r == '.':
		default:
			return false
		}
	}
	return true
}
