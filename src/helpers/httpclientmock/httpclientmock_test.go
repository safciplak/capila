package httpclientmock

import (
	"context"
	"io"
	"net/http"
	netUrl "net/url"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestSuite which encapsulate the tests for the service
type TestSuite struct {
	suite.Suite
	mocker *HTTPMocker
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestSuite))
}

// satisfying linter here
func getWithContext(client *http.Client, url string) (*http.Response, error) {
	r, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(r)
}
func postFormWithContext(client *http.Client, url string, data netUrl.Values) (*http.Response, error) {
	r, err := http.NewRequestWithContext(context.Background(), "GET", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return client.Do(r)
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	test.mocker = SetupMocker(
		"./",
		"v1",
		"v1",
		"examples",
		"success.json",
		200,
		test.T(),
	)
}

func (test *TestSuite) TestSetupMocker() {
	httpMocker := SetupMocker(
		"./",
		"v1",
		"v1",
		"examples",
		"success.json",
		200,
		test.T(),
	)
	client := httpMocker.GetClient()

	response, err := getWithContext(client, "/v1/test")
	test.Nil(err)
	test.NotNil(response)
	response.Body.Close()
}

func (test *TestSuite) TestGetMostRecentCall() {
	test.Run("it returns an error when no calls are made yet", func() {
		_, err := test.mocker.GetMostRecentCall()
		test.NotNil(err)
	})

	test.Run("it returns the last call, only one", func() {
		r, err := getWithContext(test.mocker.GetClient(), "/v1/other")
		test.Nil(err)
		defer r.Body.Close()

		call, err := test.mocker.GetMostRecentCall()

		calls := test.mocker.GetCalls()

		test.NotNil(call)
		test.Equal(calls[len(calls)-1], call)
		test.Nil(err)
	})

	test.Run("it returns the last call, multiple", func() {
		for i := 0; i < 3; i++ {
			func() {
				response, err := getWithContext(test.mocker.GetClient(), "/v1/next")
				if err == nil {
					defer response.Body.Close()
				}
			}()
		}
		call, err := test.mocker.GetMostRecentCall()

		calls := test.mocker.GetCalls()

		test.NotNil(call)
		test.Equal(calls[len(calls)-1], call)
		test.Nil(err)
	})
}

func (test *TestSuite) TestHTTPMocker_calls() {
	test.Run("it can handle multiple concurrent calls", func() {
		count := 100
		wg := sync.WaitGroup{}
		wg.Add(count)
		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				response, err := getWithContext(test.mocker.GetClient(), "/v1/next")
				if err == nil {
					defer response.Body.Close()
				}
			}()
		}

		wg.Wait()
		calls := test.mocker.GetCalls()

		test.Equal(count, len(calls))
	})
}

// nolint:funlen // Test functions should be allowed to be bigger than average when they contain multiple Runs to scope the cases
func (test *TestSuite) TestRecordCall() {
	client := test.mocker.GetClient()

	drainBody := func(body io.ReadCloser) string {
		buf := new(strings.Builder)

		defer body.Close()

		_, err := io.Copy(buf, body)
		test.Nil(err)

		s, err := netUrl.QueryUnescape(buf.String())
		test.Nil(err)

		return s
	}

	test.Run("it tracks the calls on the HTTPMocker", func() {
		response, err := getWithContext(client, "/v1/test")
		test.Nil(err)
		test.NotNil(response)
		defer response.Body.Close()

		calls := test.mocker.GetCalls()

		test.Len(calls, 1)

		r2, err := getWithContext(client, "/v1/test")
		test.Nil(err)
		defer r2.Body.Close()

		calls = test.mocker.GetCalls()

		test.Len(calls, 2)
	})

	test.Run("it captures the Request", func() {
		form := map[string][]string{
			"hello": {"world", "!"},
		}
		response, err := postFormWithContext(client, "/v1/hello", form)
		test.Nil(err)
		test.NotNil(response)
		defer response.Body.Close()

		call, err := test.mocker.GetMostRecentCall()
		test.Nil(err)

		buf := new(strings.Builder)

		body, err := call.Request.GetBody()
		test.Nil(err)

		defer body.Close()

		_, err = io.Copy(buf, body)
		test.Nil(err)

		s, err := netUrl.QueryUnescape(buf.String())
		test.Nil(err)

		test.Equal("hello=world&hello=!", s)
	})

	test.Run("it does not read the Request body", func() {
		form := map[string][]string{
			"hello": {"world", "!"},
		}
		response, err := postFormWithContext(client, "/v1/world", form)
		test.Nil(err)
		test.NotNil(response)
		defer response.Body.Close()

		call, err := test.mocker.GetMostRecentCall()
		test.Nil(err)

		// HTTPMocker did not read the body yet
		s := drainBody(call.Request.Body)
		test.Equal("hello=world&hello=!", s)

		// but we just did
		s = drainBody(call.Request.Body)
		test.Equal("", s)
	})

	test.Run("it captures the Response", func() {
		form := map[string][]string{
			"hello": {"world", "!"},
		}
		response, err := postFormWithContext(client, "/v1/test", form)
		test.Nil(err)
		test.NotNil(response)
		defer response.Body.Close()

		call, err := test.mocker.GetMostRecentCall()
		test.Nil(err)

		buf := new(strings.Builder)

		body := call.Response.Body

		defer body.Close()
		_, err = io.Copy(buf, body)
		test.Nil(err)

		test.Equal("{\n  \"greet\": \"Hello world\"\n}", buf.String())
	})

	test.Run("it provides the served ResponseBody", func() {
		form := map[string][]string{
			"hello": {"world", "!"},
		}
		response, err := postFormWithContext(client, "/v1/test", form)
		test.Nil(err)
		test.NotNil(response)
		defer response.Body.Close()

		call, err := test.mocker.GetMostRecentCall()
		test.Nil(err)

		body := call.ResponseBody

		test.Equal("{\n  \"greet\": \"Hello world\"\n}", string(body))
	})

	test.Run("it ResponseBody can be used when body has been read already", func() {
		form := map[string][]string{
			"hello": {"world", "!"},
		}
		response, err := postFormWithContext(client, "/v1/test", form)
		test.Nil(err)
		test.NotNil(response)
		defer response.Body.Close()

		// simulate SUT reading the body
		drainBody(response.Body)

		call, err := test.mocker.GetMostRecentCall()
		test.Nil(err)

		// request is already drained by SUT
		s := drainBody(call.Response.Body)
		test.Equal("", s)

		test.Equal("{\n  \"greet\": \"Hello world\"\n}", string(call.ResponseBody))
	})
}

func (test *TestSuite) TestUseStructDirectly() {
	httpMocker := &HTTPMocker{
		BasePath:   "./",
		URL:        "v1",
		APIVersion: "v1",
		Folder:     "examples",
		File:       "success.json",
		StatusCode: 200,
		Testing:    test.T(),
	}

	client := httpMocker.GetClient()

	response, err := getWithContext(client, "/v1/test")
	test.Nil(err)
	test.NotNil(response)

	defer response.Body.Close()

	buf := new(strings.Builder)

	_, err = io.Copy(buf, response.Body)
	test.Nil(err)

	test.Equal("{\n  \"greet\": \"Hello world\"\n}", buf.String())

	calls := httpMocker.GetCalls()
	test.Len(calls, 0)
}
