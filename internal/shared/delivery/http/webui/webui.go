// Package webui serves the compiled SPA frontend (see /frontend) that is
// embedded into the binary via go:embed.
package webui

import (
	"embed"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed all:dist
var distFS embed.FS

const indexFile = "index.html"

// RegisterRoutes mounts the embedded SPA as the router's NoRoute handler.
// Existing files are served as-is; any other GET/HEAD path falls back to
// index.html so client-side routing works on a hard refresh.
//
// Paths under /api/ and /swagger never fall back to the SPA and keep
// returning a JSON 404.
func RegisterRoutes(r *gin.Engine) {
	sub, err := fs.Sub(distFS, "dist")
	if err != nil {
		// "dist" is always embedded, so this cannot happen.
		panic(err)
	}
	r.NoRoute(Handler(sub))
}

// Handler returns the SPA fallback handler for the given build output.
func Handler(root fs.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path

		if isReserved(reqPath) || (c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead) {
			c.JSON(http.StatusNotFound, gin.H{"message": "route not found"})
			return
		}

		name := strings.TrimPrefix(path.Clean("/"+reqPath), "/")
		if name == "" {
			name = indexFile
		}

		if name != indexFile && isFile(root, name) {
			serve(c, root, name)
			return
		}

		// A missing static asset must be a real 404, not index.html.
		if name != indexFile && (strings.HasPrefix(name, "_app/") || path.Ext(name) != "") {
			c.Status(http.StatusNotFound)
			return
		}

		if !isFile(root, indexFile) {
			c.JSON(http.StatusNotFound, gin.H{"message": "frontend is not built"})
			return
		}
		serve(c, root, indexFile)
	}
}

// serve writes the file directly instead of using http.FileServer, which
// would redirect /index.html to ./ and break the SPA fallback.
func serve(c *gin.Context, root fs.FS, name string) {
	data, err := fs.ReadFile(root, name)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	contentType := mime.TypeByExtension(path.Ext(name))
	if contentType == "" {
		contentType = http.DetectContentType(data)
	}
	setCacheHeader(c, name)
	c.Data(http.StatusOK, contentType, data)
}

func isReserved(p string) bool {
	return p == "/api" || strings.HasPrefix(p, "/api/") ||
		p == "/swagger" || strings.HasPrefix(p, "/swagger/")
}

func isFile(root fs.FS, name string) bool {
	info, err := fs.Stat(root, name)
	return err == nil && !info.IsDir()
}

func setCacheHeader(c *gin.Context, name string) {
	if strings.HasPrefix(name, "_app/immutable/") {
		c.Header("Cache-Control", "public, max-age=31536000, immutable")
		return
	}
	if name == indexFile {
		c.Header("Cache-Control", "no-cache")
	}
}
