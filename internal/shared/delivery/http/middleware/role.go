package middleware

import (
	"errors"
	"net/http"
	"slices"

	"app/internal/shared/constants"
	"app/internal/shared/delivery/http/response"
	"app/internal/shared/domain/repository"
	"app/pkg/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RequireRole allows the request only when the authenticated user currently has
// one of the given roles. It must run after AuthMiddleware.
//
// The role is loaded from the database on every request using the UserID in the
// token claims, never from the token itself, so a role change (or deactivation)
// takes effect immediately without issuing a new token.
//
// This is a temporary authorization bridge (S0-06) until the permissions &
// roles module (PRD 03) is built.
func RequireRole(userRepo repository.UserRepository, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := GetLangFromGin(c)

		claimsVal, exists := c.Get(SESS)
		if !exists {
			abortWith(c, http.StatusUnauthorized, constants.GetErrorMessage(constants.Unauthorized, lang))
			return
		}
		claims, ok := claimsVal.(*jwt.Claims)
		if !ok || claims.UserID == "" {
			abortWith(c, http.StatusUnauthorized, constants.GetErrorMessage(constants.Unauthorized, lang))
			return
		}

		user, err := userRepo.GetByID(c.Request.Context(), claims.UserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				abortWith(c, http.StatusUnauthorized, constants.GetErrorMessage(constants.Unauthorized, lang))
				return
			}
			abortWith(c, http.StatusInternalServerError, constants.GetErrorMessage(constants.SomethingWentWrong, lang))
			return
		}

		if !user.IsActive || !slices.Contains(roles, user.Role) {
			abortWith(c, http.StatusForbidden, constants.GetErrorMessage(constants.Forbidden, lang))
			return
		}

		c.Next()
	}
}

func abortWith(c *gin.Context, status int, message string) {
	response.NewResponse(c, status, nil, message, nil)
	c.Abort()
}
