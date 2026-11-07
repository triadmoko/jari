package middleware

import (
	"app/internal/shared/constants"
	"app/internal/shared/delivery/http/response"
	"app/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	SESS = "sess"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := GetLangFromGin(c)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.NewResponse(c, http.StatusUnauthorized, nil, constants.GetErrorMessage(constants.Unauthorized, lang), nil)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.NewResponse(c, http.StatusUnauthorized, nil, constants.GetErrorMessage(constants.Unauthorized, lang), nil)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			response.NewResponse(c, http.StatusUnauthorized, nil, constants.GetErrorMessage(constants.Unauthorized, lang), nil)
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(jwtSecret, token)
		if err != nil {
			response.NewResponse(c, http.StatusUnauthorized, nil, constants.GetErrorMessage(constants.Unauthorized, lang), nil)
			c.Abort()
			return
		}

		c.Set(SESS, claims)
		c.Next()
	}
}
