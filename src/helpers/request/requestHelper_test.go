//go:generate generate-interfaces.sh

package helpers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/safciplak/capila/src/http/response"
)

// Test Suite which encapsulate the tests for the handler!
type TestSuite struct {
	suite.Suite
	ctx           context.Context
	router        *gin.Engine
	recorder      *httptest.ResponseRecorder
	requestHelper InterfaceRequestHelper
	headers       map[string]string
	response      *response.Response
	server        *httptest.Server
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)

	test.ctx = context.TODO()
	test.router = gin.Default()
	test.recorder = httptest.NewRecorder()
	test.requestHelper = NewRequestHelper(http.DefaultClient)
	test.response = response.Create()

	test.headers = map[string]string{
		"Content-Type": "application/json",
	}
}

func (test *TestSuite) setupServer(status int) {
	test.server = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(status)
		_, err := res.Write([]byte(`{"data":"response"}`))
		test.Nil(err)
	}))
}

// TearDownTest asserts whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	test.server.CloseClientConnections()
	test.server.Close()
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (test *TestSuite) TestRequest() {
	test.setupServer(http.StatusOK)

	requestBody := []byte(`{"Example": "Body"}`)
	resp, err := test.requestHelper.Request(test.ctx, http.MethodGet, test.server.URL, bytes.NewBuffer(requestBody), test.headers)

	test.Nil(err)
	test.Equal(`{"data":"response"}`, string(resp))
}

func (test *TestSuite) TestRequestNotFound() {
	test.setupServer(http.StatusNotFound)

	requestBody := []byte(`{"Example": "Body"}`)
	resp, err := test.requestHelper.Request(test.ctx, http.MethodGet, test.server.URL, bytes.NewBuffer(requestBody), test.headers)

	test.NotNil(err)
	test.Equal("", string(resp))
}

func (test *TestSuite) TestRequestServerError() {
	test.setupServer(http.StatusInternalServerError)

	requestBody := []byte(`{"Example": "Body"}`)
	resp, err := test.requestHelper.Request(test.ctx, http.MethodGet, test.server.URL, bytes.NewBuffer(requestBody), test.headers)

	test.NotNil(err)
	test.Equal("", string(resp))
}

func (test *TestSuite) TestRequestBrokenContext() {
	test.setupServer(http.StatusOK)

	//nolint:staticcheck //Testing nil content here on purpose
	resp, err := test.requestHelper.Request(nil, http.MethodGet, test.server.URL, nil, test.headers)

	test.NotNil(err)
	test.Equal("", string(resp))
}
