package errors

import "net/http"

// ErrorServiceUnavailable represents a gateway timeout HTTP error
type ErrorServiceUnavailable struct {
	ErrorBase
}

// NewErrorServiceUnavailable creates an ErrorServiceUnavailable with proper default values
func NewErrorServiceUnavailable() ErrorServiceUnavailable {
	return ErrorServiceUnavailable{
		ErrorBase: ErrorBase{
			Code:    ErrorCodeServiceUnavailable,
			Status:  http.StatusServiceUnavailable,
			Message: http.StatusText(http.StatusServiceUnavailable),
		},
	}
}

// Wrap wraps a generic error into an InterfaceError
// Put here instead of the errorBase because it'll break type inference otherwise (see unit test)
func (err ErrorServiceUnavailable) Wrap(child error) InterfaceError {
	err.Child = child
	return err
}
