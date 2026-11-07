package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func newCORSRouter(origins []string) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(CORSMiddleware(origins))
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	r.PATCH("/ping", func(c *gin.Context) { c.String(http.StatusOK, "patched") })
	return r
}

func doCORS(r *gin.Engine, method, origin string, extra map[string]string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, "/ping", nil)
	if origin != "" {
		req.Header.Set("Origin", origin)
	}
	for k, v := range extra {
		req.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestCORS_AllowedOriginIsReflected(t *testing.T) {
	r := newCORSRouter([]string{"https://app.example.com", "https://admin.example.com"})

	w := doCORS(r, http.MethodGet, "https://admin.example.com", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "https://admin.example.com", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Contains(t, w.Header().Values("Vary"), "Origin")
}

func TestCORS_UnknownOriginGetsNoCORSHeaders(t *testing.T) {
	r := newCORSRouter([]string{"https://app.example.com"})

	w := doCORS(r, http.MethodGet, "https://evil.example", nil)

	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Methods"))
	assert.Contains(t, w.Header().Values("Vary"), "Origin")
}

func TestCORS_NoOriginHeaderIsUnaffected(t *testing.T) {
	r := newCORSRouter([]string{"https://app.example.com"})

	w := doCORS(r, http.MethodGet, "", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
}

func TestCORS_EmptyAllowlistRejectsAllCrossOrigin(t *testing.T) {
	r := newCORSRouter(nil)

	w := doCORS(r, http.MethodGet, "https://app.example.com", nil)

	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
}

func TestCORS_PreflightForPatchPasses(t *testing.T) {
	r := newCORSRouter([]string{"https://app.example.com"})

	w := doCORS(r, http.MethodOptions, "https://app.example.com", map[string]string{
		"Access-Control-Request-Method":  http.MethodPatch,
		"Access-Control-Request-Headers": "authorization,content-type",
	})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "https://app.example.com", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "PATCH")
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Authorization")
}

func TestCORS_PreflightFromUnknownOriginHasNoAllowHeaders(t *testing.T) {
	r := newCORSRouter([]string{"https://app.example.com"})

	w := doCORS(r, http.MethodOptions, "https://evil.example", map[string]string{
		"Access-Control-Request-Method": http.MethodPatch,
	})

	assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Methods"))
}

func TestCORS_WildcardAllowsAnyOriginWithoutCredentials(t *testing.T) {
	r := newCORSRouter([]string{"*"})

	w := doCORS(r, http.MethodGet, "https://anything.example", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Empty(t, w.Header().Get("Access-Control-Allow-Credentials"))
}

func TestCORS_WildcardPreflightForPatchPasses(t *testing.T) {
	r := newCORSRouter([]string{"*"})

	w := doCORS(r, http.MethodOptions, "https://anything.example", map[string]string{
		"Access-Control-Request-Method": http.MethodPatch,
	})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Methods"), "PATCH")
	assert.Contains(t, w.Header().Get("Access-Control-Allow-Headers"), "Authorization")
}
