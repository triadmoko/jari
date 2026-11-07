// Package health exposes liveness and readiness probes for orchestrators
// (Docker, Kubernetes) and load balancers.
//
//   - GET /health/live  : the process is up. Never touches dependencies.
//   - GET /health/ready : dependencies work (database reachable, schema at the
//     expected migration version). 503 when any check fails.
//   - GET /health       : backwards-compatible alias of /health/live.
//
// The endpoints are public, so responses only name each check and say ok/fail;
// the underlying error goes to the log, never to the caller.
package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// DefaultTimeout bounds one readiness evaluation so a hung dependency cannot
// make the probe hang.
const DefaultTimeout = 3 * time.Second

// Check is one named readiness check. Fn returns nil when the dependency is healthy.
type Check struct {
	Name string
	Fn   func(ctx context.Context) error
}

// Handler serves the health endpoints.
type Handler struct {
	checks  []Check
	timeout time.Duration
	log     *logrus.Logger
}

// New creates a Handler that evaluates checks for readiness.
func New(log *logrus.Logger, checks ...Check) *Handler {
	return &Handler{checks: checks, timeout: DefaultTimeout, log: log}
}

// RegisterRoutes mounts /health, /health/live and /health/ready on r.
func (h *Handler) RegisterRoutes(r gin.IRoutes) {
	r.GET("/health", h.Live)
	r.GET("/health/live", h.Live)
	r.GET("/health/ready", h.Ready)
}

// Live reports that the process is running.
func (h *Handler) Live(c *gin.Context) {
	c.Header("Cache-Control", "no-store")
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Service is running",
	})
}

// Ready runs every check and answers 200 when all pass, 503 otherwise.
func (h *Handler) Ready(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.timeout)
	defer cancel()

	results := make(map[string]string, len(h.checks))
	ready := true
	for _, check := range h.checks {
		if err := run(ctx, check.Fn); err != nil {
			ready = false
			results[check.Name] = "fail"
			h.log.WithContext(c.Request.Context()).
				WithField("check", check.Name).
				Warnf("readiness check failed: %v", err)
			continue
		}
		results[check.Name] = "ok"
	}

	status, label := http.StatusOK, "ok"
	if !ready {
		status, label = http.StatusServiceUnavailable, "unavailable"
	}
	c.Header("Cache-Control", "no-store")
	c.JSON(status, gin.H{"status": label, "checks": results})
}

// run calls fn but returns as soon as ctx ends, even if fn is still blocked
// (database drivers do not always honor the context while connecting).
func run(ctx context.Context, fn func(context.Context) error) error {
	done := make(chan error, 1)
	go func() { done <- fn(ctx) }()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
