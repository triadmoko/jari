package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

func TestGormLogLevel(t *testing.T) {
	tests := []struct {
		env, level string
		want       logger.LogLevel
	}{
		{"development", "", logger.Info},
		{"production", "", logger.Warn},
		{"test", "", logger.Warn},
		{"production", "debug", logger.Info},
		{"production", "trace", logger.Info},
		{"development", "info", logger.Warn},
		{"development", "warn", logger.Warn},
		{"development", "error", logger.Error},
		{"development", "fatal", logger.Silent},
		{"development", "panic", logger.Silent},
	}
	for _, tt := range tests {
		t.Run(tt.env+"/"+tt.level, func(t *testing.T) {
			assert.Equal(t, tt.want, GormLogLevel(tt.env, tt.level))
		})
	}
}
