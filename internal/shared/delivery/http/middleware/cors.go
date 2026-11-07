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
// The request Origin is reflected, together with credentials, only when it is
// in the allowlist. Requests from any other origin get no CORS headers, so
// browsers block them; requests without an Origin header (same-origin SPA,
// curl) are unaffected. With an empty allowlist every cross-origin request is
// rejected.
//
// If the allowlist contains "*", every origin is allowed with
// "Access-Control-Allow-Origin: *" and without credentials (browsers reject a
// wildcard combined with credentials). This works for Bearer-token auth, which
// does not rely on cookies.
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		allowed[origin] = struct{}{}
	}
	_, allowAll := allowed["*"]

	return func(c *gin.Context) {
		h := c.Writer.Header()
		h.Add("Vary", "Origin")

		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			_, listed := allowed[origin]
			switch {
			case allowAll:
				h.Set("Access-Control-Allow-Origin", "*")
			case listed:
				h.Set("Access-Control-Allow-Origin", origin)
				h.Set("Access-Control-Allow-Credentials", "true")
			}
			if allowAll || listed {
				h.Set("Access-Control-Allow-Headers", corsAllowHeaders)
				h.Set("Access-Control-Allow-Methods", corsAllowMethods)
				h.Set("Access-Control-Max-Age", corsMaxAge)
			}
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
