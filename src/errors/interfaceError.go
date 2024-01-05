package errors

// InterfaceError abstracts the underlying errors away
type InterfaceError interface {
	// Error only method of the base error struct
	Error() string

	// Wrap wraps a generic error into an InterfaceError
	Wrap(err error) InterfaceError
	// Unwrap is used to unwrap the InterfaceError to get the generic error
	Unwrap() error

	// GetDetail retrieves additional details from the error
	GetDetail() string
	// GetCode retrieves a unique ErrorCode from the error, used to specify this particular error
	GetCode() ErrorCode
	// GetStatusCode Retrieves HTTP Status Code belonging to this error
	GetStatusCode() int
}
