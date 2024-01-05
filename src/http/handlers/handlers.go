package handlers

import (
	"context"
	"reflect"
	"regexp"
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/safciplak/capila/src/apm"
	"github.com/safciplak/capila/src/http/response"
)

// Request is an interface that is used by GetHandlerFunc
type Request interface {
	Validate(ctx *gin.Context) error
}

// GetAPMSpanHandlerName gets the handler name to be used in APM
func GetAPMSpanHandlerName(doRequest func(requestCtx context.Context) (interface{}, error)) string {
	nameFull := runtime.FuncForPC(reflect.ValueOf(doRequest).Pointer()).Name()

	// @TODO: VP-5368: Regex is 10x slower than string replace, check for alternatives later?
	re := regexp.MustCompile(`handlers\.\(.*?(\w*?)\)\.`)
	results := re.FindAllStringSubmatch(nameFull, -1)

	if len(results) == 1 && len(results[0]) == 2 {
		return results[0][1]
	}

	// Unable to test
	return nameFull
}

// GetHandlerFunc returns a handler function that creates and validates requests
func GetHandlerFunc(ctx *gin.Context, request Request, doRequest func(requestCtx context.Context) (interface{}, error)) {
	var (
		requestCtx   = ctx.Request.Context()
		httpResponse = response.Create()
	)

	defer apm.End(apm.Start(requestCtx, GetAPMSpanHandlerName(doRequest), "handler"))

	validationErr := request.Validate(ctx)
	if validationErr != nil {
		httpResponse.HandleValidationError(requestCtx, validationErr)
		httpResponse.ReturnJSON(ctx)

		return
	}

	responseData, responseErr := doRequest(requestCtx)
	if responseErr != nil {
		httpResponse.HandleError(requestCtx, responseErr)
		httpResponse.ReturnJSON(ctx)

		return
	}

	httpResponse.Data = responseData
	httpResponse.ReturnJSON(ctx)
}
