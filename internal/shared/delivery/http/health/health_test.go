package health

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newRouter(h *Handler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h.RegisterRoutes(r)
	return r
}

func get(r http.Handler, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
	return w
}

func decode(t *testing.T, w *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var body map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	return body
}

func ok(context.Context) error { return nil }

func failing(msg string) func(context.Context) error {
	return func(context.Context) error { return errors.New(msg) }
}

func TestLiveAndAliasIgnoreDependencies(t *testing.T) {
	h := New(logrus.New(), Check{Name: "database", Fn: failing("db down")})
	r := newRouter(h)

	for _, path := range []string{"/health", "/health/live"} {
		w := get(r, path)
		assert.Equal(t, http.StatusOK, w.Code, path)
		body := decode(t, w)
		assert.Equal(t, "ok", body["status"])
		assert.Equal(t, "Service is running", body["message"])
	}
}

func TestReadyAllChecksPass(t *testing.T) {
	h := New(logrus.New(), Check{Name: "database", Fn: ok}, Check{Name: "migrations", Fn: ok})

	w := get(newRouter(h), "/health/ready")

	assert.Equal(t, http.StatusOK, w.Code)
	body := decode(t, w)
	assert.Equal(t, "ok", body["status"])
	assert.Equal(t, map[string]any{"database": "ok", "migrations": "ok"}, body["checks"])
	assert.Equal(t, "no-store", w.Header().Get("Cache-Control"))
}

func TestReadyDatabaseDownReturns503WithoutLeakingError(t *testing.T) {
	h := New(logrus.New(),
		Check{Name: "database", Fn: failing("dial tcp 10.0.0.5:5432: connection refused")},
		Check{Name: "migrations", Fn: ok},
	)

	w := get(newRouter(h), "/health/ready")

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	body := decode(t, w)
	assert.Equal(t, "unavailable", body["status"])
	assert.Equal(t, map[string]any{"database": "fail", "migrations": "ok"}, body["checks"])
	assert.NotContains(t, w.Body.String(), "10.0.0.5")
	assert.NotContains(t, w.Body.String(), "connection refused")
}

func TestReadyTimesOutOnHungCheck(t *testing.T) {
	block := make(chan struct{})
	defer close(block)
	h := New(logrus.New(), Check{Name: "database", Fn: func(context.Context) error {
		<-block
		return nil
	}})
	h.timeout = 50 * time.Millisecond

	start := time.Now()
	w := get(newRouter(h), "/health/ready")

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	assert.Less(t, time.Since(start), time.Second)
}

func TestReadyWithNoChecksIsReady(t *testing.T) {
	w := get(newRouter(New(logrus.New())), "/health/ready")
	assert.Equal(t, http.StatusOK, w.Code)
}
