package errors

import "net/http"

// ErrorBadRequest represents a bad request HTTP error
type ErrorBadRequest struct {
	ErrorBase
}

// NewErrorBadRequest creates an ErrorBadRequest with proper default values
func NewErrorBadRequest() ErrorBadRequest {
	return ErrorBadRequest{
		ErrorBase: ErrorBase{
			Code:    ErrorCodeBadRequest,
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		},
	}
}

// Wrap wraps a generic error into an InterfaceError
// Put here instead of the errorBase because it'll break type inference otherwise (see unit test)
func (err ErrorBadRequest) Wrap(child error) InterfaceError {
	err.Child = child
	return err
}
