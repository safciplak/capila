// code generated by ifacemaker; DO NOT EDIT.

package response

import (
	"context"

	"github.com/safciplak/capila/src/http/presenters"
	"github.com/gin-gonic/gin"
)

// InterfaceResponse is the interface implemented by Response
type InterfaceResponse interface {
	// CheckForErrors handles the validation results for the response.
	// NOTE: validationErrors left in for backwards compatible reasons
	CheckForErrors(ctx context.Context, validationErrors map[string]string, err error)
	// SetError sets the error and the default StatusCode.
	SetError(ctx context.Context, err error)
	// HasErrors checks if there are any errors.
	HasErrors() bool
	// HasValidationErrors checks if the response has validation errors.
	HasValidationErrors() bool
	// HandleValidationError handles validator v10 errors.
	HandleValidationError(ctx context.Context, err error) *Response
	// HandleError handles custom errors and returns the correct http status code.
	HandleError(ctx context.Context, err error) *Response
	// ReturnJSON returns the response in JSON.
	ReturnJSON(ctx *gin.Context)
	// Output returns the response with the given presenter.
	Output(presenter presenters.InterfacePresenter)
}
