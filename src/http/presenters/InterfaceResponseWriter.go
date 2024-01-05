package presenters

import "net/http"

// InterfaceResponseWriter allows the response writer to be mocked
type InterfaceResponseWriter interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}
