//go:generate generate-interfaces.sh

package helpers

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// RequestHelper is a utility for sending requests.
type RequestHelper struct {
	httpClient *http.Client
}

// NewRequestHelper returns a new RequestHelper.
func NewRequestHelper(httpClient *http.Client) InterfaceRequestHelper {
	return &RequestHelper{
		httpClient: httpClient,
	}
}

// request is the private function for handling the request.
func (requestHelper *RequestHelper) request(ctx context.Context, method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	var (
		request  *http.Request
		response *http.Response
		err      error
	)

	request, err = http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return response, err
	}

	for headerKey, headerValue := range headers {
		request.Header.Set(headerKey, headerValue)
	}

	return requestHelper.httpClient.Do(request)
}

// Request handles a request to an external service / api.
func (requestHelper *RequestHelper) Request(ctx context.Context, method, url string, body io.Reader, headers map[string]string) ([]byte, error) {
	var (
		response     *http.Response
		err          error
		responseBody []byte = nil
	)

	response, err = requestHelper.request(ctx, method, url, body, headers)
	if err != nil {
		return responseBody, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return responseBody, errors.New(response.Status)
	}

	return io.ReadAll(response.Body)
}
