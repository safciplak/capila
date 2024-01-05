package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/http/response"
)

// PanicMiddleware formats the response according to the standard defined in the various presenters instead of
// returning an empty response
func PanicMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(ctx *gin.Context, recovered interface{}) {
		var (
			requestCtx   = ctx.Request.Context()
			httpResponse = response.Create()
		)
		if err, ok := recovered.(string); ok {
			httpResponse.HandleError(requestCtx, errors.New(err))
			httpResponse.ReturnJSON(ctx)
		}

		httpResponse.HandleError(requestCtx, errors.New("panic encountered"))
		httpResponse.ReturnJSON(ctx)
	})
}
