package middleware

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware writes one structured log line per request through logrus,
// tagged with the request ID. Install it after RequestIDMiddleware. The level
// follows the status: 5xx error, 4xx warn, otherwise info.
func LoggerMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		// FullPath is the route template (no IDs or query values), "" when unmatched.
		route := c.FullPath()
		if route == "" {
			route = "unmatched"
		}

		entry := log.WithContext(c.Request.Context()).WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"path":       path,
			"route":      route,
			"status":     c.Writer.Status(),
			"latency_ms": float64(time.Since(start).Microseconds()) / 1000,
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
			"size":       c.Writer.Size(),
		})
		if errs := c.Errors.ByType(gin.ErrorTypePrivate).String(); errs != "" {
			entry = entry.WithField("errors", errs)
		}

		switch status := c.Writer.Status(); {
		case status >= http.StatusInternalServerError:
			entry.Error("request completed")
		case status >= http.StatusBadRequest:
			entry.Warn("request completed")
		default:
			entry.Info("request completed")
		}
	}
}

// RecoveryMiddleware turns a handler panic into a 500 response and logs it with
// the request ID, instead of gin's default plain-text dump on stderr.
func RecoveryMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return gin.CustomRecoveryWithWriter(io.Discard, func(c *gin.Context, recovered any) {
		log.WithContext(c.Request.Context()).Errorf("panic recovered: %v", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
