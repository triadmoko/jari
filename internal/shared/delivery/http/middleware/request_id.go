package middleware

import (
	"app/pkg/requestid"

	"github.com/gin-gonic/gin"
)

const RequestIDKey = "request_id"

// RequestIDMiddleware assigns every request a correlation ID. It reuses a
// well-formed X-Request-ID sent by the caller (so a proxy or client can trace a
// request end to end) and generates a new one otherwise. The ID is stored in the
// gin and request contexts and echoed in the response header.
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(requestid.Header)
		if !requestid.Valid(id) {
			id = requestid.New()
		}

		c.Set(RequestIDKey, id)
		c.Request = c.Request.WithContext(requestid.NewContext(c.Request.Context(), id))
		c.Header(requestid.Header, id)

		c.Next()
	}
}
