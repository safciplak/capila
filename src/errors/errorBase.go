package errors

// ErrorBase represents an incomplete base error which needs to be extended in order to function properly
// Introduced simply to reduce Code duplication
type ErrorBase struct {
	error
	Child   error
	Message string
	Code    string
	Status  int
}

// Is required for proper ErrorIs and ErrorAs functionality
func (err ErrorBase) Is(target error) bool {
	return err.Error() == target.Error()
}

// Error returns an error description to be compliant with the Error interface
func (err ErrorBase) Error() string {
	return err.Message
}

// GetStatusCode retrieves HTTP Status Code belonging to this error
func (err ErrorBase) GetStatusCode() int {
	return err.Status
}

// GetDetail retrieves additional details from the error, needs to be overridden to work
func (err ErrorBase) GetDetail() string {
	return ""
}

// GetCode retrieves a unique Code from the error, used to specify this particular error
func (err ErrorBase) GetCode() string {
	return err.Code
}

// Unwrap is used to unwrap the known error to get the underlying error
func (err ErrorBase) Unwrap() error {
	return err.Child
}
