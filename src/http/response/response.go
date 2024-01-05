//go:generate generate-interfaces.sh

package response

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/safciplak/capila/src/apm"
	capilaErrors "github.com/safciplak/capila/src/errors"
	"github.com/safciplak/capila/src/http/presenters"
	"github.com/safciplak/capila/src/http/presenters/JSON"
)

// Response holds the data the generic response gives
type Response struct {
	Error      capilaErrors.InterfaceError
	Data       interface{}
	Links      []string
	StatusCode int
}

// Create initiates a fresh response struct.
func Create() *Response {
	var response = new(Response)
	response.StatusCode = 200

	return response
}

// CheckForErrors handles the validation results for the response.
// NOTE: validationErrors left in for backwards compatible reasons
func (response *Response) CheckForErrors(ctx context.Context, _ map[string]string, err error) {
	if err != nil {
		response.SetError(ctx, err)
	}
}

// SetError sets the error and the default StatusCode.
func (response *Response) SetError(ctx context.Context, err error) {
	var internalErr capilaErrors.InterfaceError

	// Trace the generic error
	err = apm.TraceError(ctx, err)

	switch value := err.(type) {
	case capilaErrors.InterfaceError:
		internalErr = value
	case validator.ValidationErrors:
		// Wrap request validation errors in an bad request error
		internalErr = capilaErrors.NewErrorBadRequest().Wrap(value)
	default:
		// If the error which is set explicitly is not a known error, wrap it in a server error
		internalErr = capilaErrors.NewErrorInternalServerError().Wrap(value)
	}

	response.Error = internalErr
	response.StatusCode = internalErr.GetStatusCode()
}

// HasErrors checks if there are any errors.
func (response *Response) HasErrors() bool {
	return response.Error != nil
}

// HasValidationErrors checks if the response has validation errors.
func (response *Response) HasValidationErrors() bool {
	return response.Error != nil && errors.Is(response.Error, capilaErrors.NewErrorBadRequest())
}

// HandleValidationError handles validator v10 errors.
func (response *Response) HandleValidationError(ctx context.Context, err error) *Response {
	response.SetError(ctx, capilaErrors.NewErrorBadRequest().Wrap(err))
	return response
}

// HandleError handles custom errors and returns the correct http status code.
func (response *Response) HandleError(ctx context.Context, err error) *Response {
	response.SetError(ctx, err)
	return response
}

// ReturnJSON returns the response in JSON.
func (response *Response) ReturnJSON(ctx *gin.Context) {
	response.Output(JSON.Present(ctx))
}

// Output returns the response with the given presenter.
func (response *Response) Output(presenter presenters.InterfacePresenter) {
	if response.HasErrors() {
		presenter.Error(response.StatusCode, response.Error)
		return
	}

	presenter.Success(response.StatusCode, response.Data)
}
