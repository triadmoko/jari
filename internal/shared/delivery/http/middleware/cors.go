package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	corsAllowHeaders = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
	corsAllowMethods = "POST, OPTIONS, GET, PUT, PATCH, DELETE"
	corsMaxAge       = "600"
)

// CORSMiddleware creates a CORS middleware that only allows the given origins.
//
// The request Origin is reflected only when it is in the allowlist (a wildcard
// is incompatible with credentials). Requests from any other origin get no
// CORS headers, so browsers block them; requests without an Origin header
// (same-origin SPA, curl) are unaffected. With an empty allowlist every
// cross-origin request is rejected.
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		allowed[origin] = struct{}{}
	}

	return func(c *gin.Context) {
		h := c.Writer.Header()
		h.Add("Vary", "Origin")

		origin := c.Request.Header.Get("Origin")
		if _, ok := allowed[origin]; ok && origin != "" {
			h.Set("Access-Control-Allow-Origin", origin)
			h.Set("Access-Control-Allow-Credentials", "true")
			h.Set("Access-Control-Allow-Headers", corsAllowHeaders)
			h.Set("Access-Control-Allow-Methods", corsAllowMethods)
			h.Set("Access-Control-Max-Age", corsMaxAge)
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
