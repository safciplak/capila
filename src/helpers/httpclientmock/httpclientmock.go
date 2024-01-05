package httpclientmock

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// roundTripFunc accepts a func with http.Request and returns a http.Response.
type roundTripFunc func(req *http.Request) *http.Response

// RoundTrip simply executes the func passed.
func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
//
//nolint:interfacer //ignore for now
func NewTestClient(roundTrip roundTripFunc) *http.Client {
	return &http.Client{
		Transport: roundTrip,
	}
}

type Call struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
}

// HTTPMocker helps mocking HTTP calls
type HTTPMocker struct {
	Testing    *testing.T
	calls      chan []*Call
	BasePath   string
	URL        string
	APIVersion string
	Folder     string
	File       string
	StatusCode int
}

// SetupMocker is a quick setter for the values
func SetupMocker(basePath, url, apiVersion, folder, file string, statusCode int, t *testing.T) *HTTPMocker {
	httpMocker := &HTTPMocker{
		BasePath:   basePath,
		URL:        url,
		APIVersion: apiVersion,
		Folder:     folder,
		File:       file,
		StatusCode: statusCode,
		Testing:    t,
		calls:      make(chan []*Call, 1),
	}
	httpMocker.calls <- []*Call{}

	return httpMocker
}

// GetClient returns a mock client
func (httpMocker *HTTPMocker) GetClient() *http.Client {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Contains(httpMocker.Testing, req.URL.String(), httpMocker.URL)
		httpHeader := make(http.Header)
		httpHeader.Add("Content-Type", "application/json")

		responseBody := httpMocker.getJSONResponse()
		// without Content-Type the client will crash
		response := &http.Response{
			StatusCode: httpMocker.StatusCode,
			Status:     http.StatusText(httpMocker.StatusCode),
			Body:       io.NopCloser(bytes.NewBuffer(httpMocker.getJSONResponse())),
			Header:     httpHeader,
		}
		if httpMocker.calls != nil {
			calls := <-httpMocker.calls
			calls = append(calls, &Call{req, response, responseBody})
			httpMocker.calls <- calls
		}
		return response
	})

	return client
}

// getJSONResponse returns a json Response from the given folder
func (httpMocker *HTTPMocker) getJSONResponse() []byte {
	content, err := os.ReadFile(
		filepath.Join(httpMocker.BasePath, "mock", httpMocker.APIVersion, httpMocker.Folder, httpMocker.File))

	assert.Nil(httpMocker.Testing, err)

	return content
}

// GetCalls returns all calls made during the lifetime of this HTTPMocker
func (httpMocker *HTTPMocker) GetCalls() []*Call {
	if httpMocker.calls == nil {
		return []*Call{}
	}

	calls := <-httpMocker.calls
	httpMocker.calls <- calls

	return calls
}

// GetMostRecentCall returns the most recent call
func (httpMocker *HTTPMocker) GetMostRecentCall() (*Call, error) {
	calls := httpMocker.GetCalls()
	if len(calls) == 0 {
		return nil, errors.New("no call available")
	}

	call := calls[len(calls)-1]

	return call, nil
}
