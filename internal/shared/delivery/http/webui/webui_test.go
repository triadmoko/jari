package webui

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func newRouter(files fstest.MapFS) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	r.NoRoute(Handler(files))
	return r
}

func do(r *gin.Engine, method, target string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(method, target, nil))
	return w
}

func TestHandler(t *testing.T) {
	files := fstest.MapFS{
		"index.html":                  {Data: []byte("<html>spa</html>")},
		"robots.txt":                  {Data: []byte("User-agent: *")},
		"_app/immutable/entry/app.js": {Data: []byte("console.log(1)")},
	}
	r := newRouter(files)

	t.Run("root serves index with no-cache", func(t *testing.T) {
		w := do(r, http.MethodGet, "/")
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "spa")
		assert.Equal(t, "no-cache", w.Header().Get("Cache-Control"))
	})

	t.Run("unknown client route falls back to index", func(t *testing.T) {
		w := do(r, http.MethodGet, "/admin/users")
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "spa")
	})

	t.Run("existing file is served", func(t *testing.T) {
		w := do(r, http.MethodGet, "/robots.txt")
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "User-agent")
	})

	t.Run("immutable asset has long cache", func(t *testing.T) {
		w := do(r, http.MethodGet, "/_app/immutable/entry/app.js")
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Header().Get("Cache-Control"), "immutable")
	})

	t.Run("missing asset is 404 not index", func(t *testing.T) {
		assert.Equal(t, http.StatusNotFound, do(r, http.MethodGet, "/_app/immutable/missing.js").Code)
		assert.Equal(t, http.StatusNotFound, do(r, http.MethodGet, "/favicon.png").Code)
	})

	t.Run("api and swagger never fall back", func(t *testing.T) {
		for _, p := range []string{"/api/v1/nope", "/api", "/swagger/nope"} {
			w := do(r, http.MethodGet, p)
			assert.Equal(t, http.StatusNotFound, w.Code, p)
			assert.Contains(t, w.Header().Get("Content-Type"), "application/json", p)
		}
	})

	t.Run("non-GET does not fall back", func(t *testing.T) {
		assert.Equal(t, http.StatusNotFound, do(r, http.MethodPost, "/admin").Code)
	})

	t.Run("registered routes win", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, do(r, http.MethodGet, "/health").Code)
	})

	t.Run("path traversal is contained", func(t *testing.T) {
		w := do(r, http.MethodGet, "/../../etc/passwd")
		assert.NotContains(t, w.Body.String(), "root:")
	})
}

func TestHandler_NotBuilt(t *testing.T) {
	r := newRouter(fstest.MapFS{})
	w := do(r, http.MethodGet, "/")
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRegisterRoutes_EmbedCompiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	assert.NotPanics(t, func() { RegisterRoutes(r) })
}

func TestRegisterRoutes_RealBuild(t *testing.T) {
	if _, err := distFS.ReadFile("dist/index.html"); err != nil {
		t.Skip("frontend not built (run `make web-build`)")
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	RegisterRoutes(r)

	for _, p := range []string{"/", "/admin", "/auth/login?redirectTo=%2Fadmin"} {
		w := do(r, http.MethodGet, p)
		assert.Equal(t, http.StatusOK, w.Code, p)
		assert.Contains(t, w.Body.String(), "__sveltekit_", p)
	}
	assert.Equal(t, http.StatusOK, do(r, http.MethodGet, "/health").Code)
	assert.Equal(t, http.StatusNotFound, do(r, http.MethodGet, "/api/v1/nope").Code)
}
