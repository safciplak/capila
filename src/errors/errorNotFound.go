package errors

import "net/http"

// ErrorNotFound represents a resource not found error
type ErrorNotFound struct {
	ErrorBase
}

// NewErrorNotFound creates an ErrorNotFound with proper default values
func NewErrorNotFound() ErrorNotFound {
	return ErrorNotFound{
		ErrorBase: ErrorBase{
			Code:    ErrorCodeNotFound,
			Status:  http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		},
	}
}

// Wrap wraps a generic error into an InterfaceError
// Put here instead of the errorBase because it'll break type inference otherwise (see unit test)
func (err ErrorNotFound) Wrap(child error) InterfaceError {
	err.Child = child
	return err
}
