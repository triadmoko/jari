package requestid

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextRoundTrip(t *testing.T) {
	assert.Equal(t, "", FromContext(context.Background()))
	assert.Equal(t, "", FromContext(nil)) //nolint:staticcheck // nil ctx must not panic
	assert.Equal(t, "abc-1", FromContext(NewContext(context.Background(), "abc-1")))
}

func TestNewIsValidAndUnique(t *testing.T) {
	a, b := New(), New()
	assert.NotEqual(t, a, b)
	assert.True(t, Valid(a))
}

func TestValid(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want bool
	}{
		{"uuid", "3f2b8c1e-9d4a-4c55-8a53-0d7f1e2b6a10", true},
		{"alnum with separators", "req_1.2-A", true},
		{"empty", "", false},
		{"space", "a b", false},
		{"newline (log forging)", "abc\nlevel=error", false},
		{"unicode", "réq", false},
		{"too long", strings.Repeat("a", 129), false},
		{"max length", strings.Repeat("a", 128), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Valid(tt.id))
		})
	}
}
