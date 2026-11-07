package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/pkg/logger"
	"app/pkg/requestid"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newEngine(buf *bytes.Buffer) *gin.Engine {
	gin.SetMode(gin.TestMode)
	log := logger.NewLogger("info", true)
	log.SetOutput(buf)

	r := gin.New()
	r.Use(RequestIDMiddleware(), LoggerMiddleware(log), RecoveryMiddleware(log))
	r.GET("/ping", func(c *gin.Context) {
		log.WithContext(c.Request.Context()).Info("inside handler")
		c.JSON(http.StatusOK, gin.H{"id": requestid.FromContext(c.Request.Context())})
	})
	r.GET("/boom", func(c *gin.Context) { panic("kaboom") })
	return r
}

func logLines(t *testing.T, buf *bytes.Buffer) []map[string]any {
	t.Helper()
	var lines []map[string]any
	for _, raw := range bytes.Split(bytes.TrimSpace(buf.Bytes()), []byte("\n")) {
		var m map[string]any
		require.NoError(t, json.Unmarshal(raw, &m), string(raw))
		lines = append(lines, m)
	}
	return lines
}

func do(r http.Handler, path, rid string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	if rid != "" {
		req.Header.Set(requestid.Header, rid)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRequestIDGeneratedWhenMissing(t *testing.T) {
	var buf bytes.Buffer
	w := do(newEngine(&buf), "/ping", "")

	id := w.Header().Get(requestid.Header)
	assert.True(t, requestid.Valid(id))
	assert.Contains(t, w.Body.String(), id, "handler sees the same ID via request context")
}

func TestRequestIDReusedWhenValid(t *testing.T) {
	var buf bytes.Buffer
	w := do(newEngine(&buf), "/ping", "client-abc.123")
	assert.Equal(t, "client-abc.123", w.Header().Get(requestid.Header))
}

func TestRequestIDReplacedWhenInvalid(t *testing.T) {
	var buf bytes.Buffer
	w := do(newEngine(&buf), "/ping", "bad id\"; level=error")

	id := w.Header().Get(requestid.Header)
	assert.NotEqual(t, "bad id\"; level=error", id)
	assert.True(t, requestid.Valid(id))
}

func TestEveryLogLineCarriesRequestID(t *testing.T) {
	var buf bytes.Buffer
	w := do(newEngine(&buf), "/ping", "rid-42")
	require.Equal(t, http.StatusOK, w.Code)

	lines := logLines(t, &buf)
	require.Len(t, lines, 2) // handler line + request line
	for _, l := range lines {
		assert.Equal(t, "rid-42", l["request_id"])
	}
	req := lines[1]
	assert.Equal(t, "request completed", req["message"])
	assert.EqualValues(t, 200, req["status"])
	assert.Equal(t, "/ping", req["route"])
	assert.Equal(t, "GET", req["method"])
}

func TestPanicIsRecoveredAndLoggedWithRequestID(t *testing.T) {
	var buf bytes.Buffer
	w := do(newEngine(&buf), "/boom", "rid-panic")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "rid-panic", w.Header().Get(requestid.Header))

	lines := logLines(t, &buf)
	require.NotEmpty(t, lines)
	assert.Contains(t, lines[0]["message"], "kaboom")
	for _, l := range lines {
		assert.Equal(t, "rid-panic", l["request_id"])
	}
	assert.Equal(t, "error", lines[len(lines)-1]["severity"])
}

func TestUnmatchedRouteLogsWithoutRawPath(t *testing.T) {
	var buf bytes.Buffer
	do(newEngine(&buf), "/nope", "rid-404")

	lines := logLines(t, &buf)
	require.Len(t, lines, 1)
	assert.Equal(t, "unmatched", lines[0]["route"])
	assert.Equal(t, "warning", lines[0]["severity"])
}
