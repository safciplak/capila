package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/http/context"
)

// LanguageMiddleware attempts to extract the language from the headers and set it in the context ( defaults to en ).
func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			language   = strings.ToLower(c.Request.Header.Get("Accept-Language"))
			currentCtx = c.Request.Context()
		)

		if language != "" {
			c.Request = c.Request.Clone(context.SetLanguage(currentCtx, language))
		}
	}
}
