package middleware

import (
	"app/internal/shared/constants"
	"context"

	"github.com/gin-gonic/gin"
)

// LangKey is the key under which the language is stored in the gin context.
const LangKey = "lang"

// ctxKey is a dedicated type for context.Context keys so they cannot collide
// with keys set by other packages (staticcheck SA1029).
type ctxKey string

// LangCtxKey is the key under which the language is stored in the request context.
const LangCtxKey ctxKey = "lang"

func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")

		switch lang {
		case "id", "ID":
			lang = string(constants.LangID)
		default:
			lang = string(constants.LangEN)
		}

		c.Set(LangKey, constants.Lang(lang))

		ctx := context.WithValue(c.Request.Context(), LangCtxKey, constants.Lang(lang))
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func GetLangFromContext(ctx context.Context) constants.Lang {
	if lang, ok := ctx.Value(LangCtxKey).(constants.Lang); ok {
		return lang
	}
	return constants.LangEN
}

func GetLangFromGin(c *gin.Context) constants.Lang {
	if lang, exists := c.Get(LangKey); exists {
		if l, ok := lang.(constants.Lang); ok {
			return l
		}
	}
	return constants.LangEN
}
