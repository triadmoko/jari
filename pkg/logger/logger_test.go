package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"app/pkg/requestid"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLevel(t *testing.T) {
	assert.Equal(t, logrus.DebugLevel, ParseLevel("DEBUG"))
	assert.Equal(t, logrus.WarnLevel, ParseLevel(" warning "))
	assert.Equal(t, logrus.InfoLevel, ParseLevel(""))
	assert.Equal(t, logrus.InfoLevel, ParseLevel("bogus"))
}

func TestJSONLoggerIncludesRequestID(t *testing.T) {
	var buf bytes.Buffer
	l := NewLogger("info", true)
	l.SetOutput(&buf)

	ctx := requestid.NewContext(context.Background(), "rid-123")
	l.WithContext(ctx).WithField("status", 200).Info("hello")

	var line map[string]any
	require.NoError(t, json.Unmarshal(buf.Bytes(), &line), buf.String())
	assert.Equal(t, "rid-123", line["request_id"])
	assert.Equal(t, "hello", line["message"])
	assert.Equal(t, "info", line["severity"])
	assert.EqualValues(t, 200, line["status"])
	assert.Contains(t, line, "timestamp")
}

func TestJSONLoggerWithoutRequestID(t *testing.T) {
	var buf bytes.Buffer
	l := NewLogger("info", true)
	l.SetOutput(&buf)

	l.Info("boot")

	var line map[string]any
	require.NoError(t, json.Unmarshal(buf.Bytes(), &line))
	assert.NotContains(t, line, "request_id")
}

func TestTextLoggerPrintsFields(t *testing.T) {
	var buf bytes.Buffer
	l := NewLogger("info", false)
	l.SetOutput(&buf)

	ctx := requestid.NewContext(context.Background(), "rid-9")
	l.WithContext(ctx).Info("hello")

	assert.Contains(t, buf.String(), "hello")
	assert.Contains(t, buf.String(), "request_id=rid-9")
}

func TestLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	l := NewLogger("warn", true)
	l.SetOutput(&buf)

	l.Info("quiet")
	assert.Empty(t, buf.String())
	l.Warn("loud")
	assert.Contains(t, buf.String(), "loud")
}
