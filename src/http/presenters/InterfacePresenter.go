package presenters

// InterfacePresenter is the interface presenters should comply to.
type InterfacePresenter interface {
	Error(statusCode int, err error)
	Success(statusCode int, data interface{})
}
