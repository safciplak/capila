package errors

import "net/http"

// ErrorBadGateway represents a bad gateway HTTP error
type ErrorBadGateway struct {
	ErrorBase
}

// NewErrorBadGateway creates an ErrorBadGateway with proper default values
func NewErrorBadGateway() ErrorBadGateway {
	return ErrorBadGateway{
		ErrorBase: ErrorBase{
			Code:    ErrorCodeBadGateway,
			Status:  http.StatusBadGateway,
			Message: http.StatusText(http.StatusBadGateway),
		},
	}
}

// Wrap wraps a generic error into an InterfaceError
// Put here instead of the errorBase because it'll break type inference otherwise (see unit test)
func (err ErrorBadGateway) Wrap(child error) InterfaceError {
	err.Child = child
	return err
}
