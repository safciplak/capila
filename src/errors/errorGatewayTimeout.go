package errors

import "net/http"

// ErrorGatewayTimeout represents a gateway timeout HTTP error
type ErrorGatewayTimeout struct {
	ErrorBase
}

// NewErrorGatewayTimeout creates an ErrorGatewayTimeout with proper default values
func NewErrorGatewayTimeout() ErrorGatewayTimeout {
	return ErrorGatewayTimeout{
		ErrorBase: ErrorBase{
			Code:    ErrorCodeGatewayTimeout,
			Status:  http.StatusGatewayTimeout,
			Message: http.StatusText(http.StatusGatewayTimeout),
		},
	}
}

// Wrap wraps a generic error into an InterfaceError
// Put here instead of the errorBase because it'll break type inference otherwise (see unit test)
func (err ErrorGatewayTimeout) Wrap(child error) InterfaceError {
	err.Child = child
	return err
}
