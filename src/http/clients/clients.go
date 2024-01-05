package clients

import (
	"net/http"

	"go.elastic.co/apm/module/apmhttp/v2"
)

// NewAPMHTTPClient instantiates an APM wrapped client
func NewAPMHTTPClient() *http.Client {
	return apmhttp.WrapClient(http.DefaultClient)
}

// NewHTTPClient instantiates a client.
func NewHTTPClient() *http.Client {
	return http.DefaultClient
}
