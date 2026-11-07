// Package brand is the single place that defines the product identity on the
// backend. "Jari" is the final product name; if it ever changes, change Name/Slug here.
//
// Not covered: the Swagger annotations in cmd/api/main.go (@title and
// @description) are comments, so swag cannot read these constants; update them by hand.
package brand

import "net/http"

const (
	Name = "Jari"
	// Slug is the lowercase identifier for image names, labels, and storage keys.
	Slug = "jari"
	// Description is a one-line product summary (OCI image label, API description).
	Description = "Open-source work and project management platform"
)

// Version is the build version. Override at build time:
//
//	go build -ldflags "-X app/pkg/brand.Version=1.2.3"
var Version = "dev"

// UserAgent returns the User-Agent value for outgoing HTTP requests, e.g. "Jari/dev".
func UserAgent() string {
	return Name + "/" + Version
}

type userAgentTransport struct {
	base http.RoundTripper
}

// RoundTrip sets the User-Agent unless the caller already set one. The request
// is cloned first because a RoundTripper must not modify the one it receives.
func (t userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("User-Agent") == "" {
		req = req.Clone(req.Context())
		req.Header.Set("User-Agent", UserAgent())
	}
	return t.base.RoundTrip(req)
}

// NewTransport wraps base (http.DefaultTransport when nil) so every outgoing
// request carries UserAgent(). Use it for every outbound http.Client
// (webhook, email transport, integrations).
func NewTransport(base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return userAgentTransport{base: base}
}
