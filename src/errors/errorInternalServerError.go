package errors

import "net/http"

// ErrorInternalServerError represents an internal server error
type ErrorInternalServerError struct {
	ErrorBase
}

// NewErrorInternalServerError creates an ErrorInternalServerError with proper default values
func NewErrorInternalServerError() ErrorInternalServerError {
	return ErrorInternalServerError{
		ErrorBase: ErrorBase{
			Code:    ErrorCodeInternalServerError,
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		},
	}
}

// Wrap wraps a generic error into an InterfaceError
// Put here instead of the errorBase because it'll break type inference otherwise (see unit test)
func (err ErrorInternalServerError) Wrap(child error) InterfaceError {
	err.Child = child
	return err
}
