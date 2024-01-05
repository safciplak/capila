package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

// TestSuite which encapsulate the tests for the handler!
type TestSuite struct {
	suite.Suite
	ctx          *gin.Context
	router       *gin.Engine
	recorder     *httptest.ResponseRecorder
	request      *testRequest
	errorRequest *testRequest
	engine       *gin.Engine
}

const exampleData = "exampleData"

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	test.router = gin.New()
	test.recorder = httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(test.recorder)
	test.ctx = ctx
	test.engine = r

	test.request = &testRequest{}
	test.errorRequest = &testRequest{ShowError: true}
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TestSuite))
}

// testRequest is a struct that is used by GetHandlerFunc tests
type testRequest struct {
	GUID         string `json:"guid" form:"guid" binding:"required,uuid"`
	NestedObject nestedObject
	Name         string
	ShowError    bool
}

type nestedObject struct {
	GUID string `json:"guid" form:"guid" binding:"required,uuid"`
	Name string `json:"name" form:"name" binding:"required"`
}

func postRequestEmptyGUID() *bytes.Buffer {
	return bytes.NewBufferString(`{
    "guid": ""
}`)
}

func getValidationResponseString() string {
	return `{"errors":[` +
		`{"code":"ERROR_BAD_REQUEST","title":"errors.ErrorBadRequest","status":400},` +
		`{"code":"ERROR_INPUT_VALIDATION","title":"validator.fieldError","detail":"Key: 'testRequest.GUID' Error:Field validation for 'GUID' failed on the 'required' tag"},` +
		`{"code":"ERROR_INPUT_VALIDATION","title":"validator.fieldError","detail":"Key: 'testRequest.NestedObject.GUID' Error:Field validation for 'GUID' failed on the 'required' tag"},` +
		`{"code":"ERROR_INPUT_VALIDATION","title":"validator.fieldError","detail":"Key: 'testRequest.NestedObject.Name' Error:Field validation for 'Name' failed on the 'required' tag"}],` +
		`"_links":{"self":{"href":"/test"}},"meta":{"query":{}}}`
}

func getErrorResponseString() string {
	return `{"errors":[` +
		`{"code":"ERROR_INTERNAL_SERVER_ERROR","title":"errors.ErrorInternalServerError","status":500},` +
		`{"code":"ERROR_UNKNOWN","title":"errors.errorString","detail":"an error"}],` +
		`"_links":{"self":{"href":"/test"}},"meta":{"query":{}}}`
}

// Validate is an example validate implementation
func (request *testRequest) Validate(ctx *gin.Context) error {
	request.Name = "dummy"

	if request.ShowError {
		return ctx.ShouldBindJSON(request)
	}

	return nil
}

// BenchmarkGetAPMSpanName checks the benchmark of reflection for the APM Span Name
func BenchmarkGetAPMSpanName(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAPMSpanHandlerName(func(requestCtx context.Context) (interface{}, error) {
			return exampleData, nil
		})
	}
}

// TestGetHandlerFunc tests the success flow of the handlerFunc
func (test *TestSuite) TestGetHandlerFunc() {
	var (
		err    error
		hasRun = false
	)

	test.engine.POST("/test", func(ctx *gin.Context) {
		GetHandlerFunc(test.ctx, test.request, func(requestCtx context.Context) (interface{}, error) {
			hasRun = true

			return exampleData, nil
		})
	})

	test.ctx.Request, err = http.NewRequest(http.MethodPost, "/test", nil)
	test.Nil(err)

	test.engine.ServeHTTP(test.recorder, test.ctx.Request)

	test.Equal(`{"data":"exampleData","_links":{"self":{"href":"/test"}},"meta":{"query":{}}}`,
		test.recorder.Body.String())
	test.Equal(http.StatusOK, test.recorder.Code)
	test.Equal("dummy", test.request.Name)
	test.Equal(true, hasRun)
}

// TestGetHandlerFunc tests the error handling when the Request validation throws an error
func (test *TestSuite) TestGetHandlerFunc_validationErr() {
	var (
		err    error
		hasRun = false
	)

	test.engine.POST("/test", func(ctx *gin.Context) {
		GetHandlerFunc(test.ctx, test.errorRequest, func(requestCtx context.Context) (interface{}, error) {
			hasRun = true

			return exampleData, nil
		})
	})

	test.ctx.Request, err = http.NewRequest(http.MethodPost, "/test", postRequestEmptyGUID())
	test.Nil(err)

	test.engine.ServeHTTP(test.recorder, test.ctx.Request)

	test.Equal(getValidationResponseString(), test.recorder.Body.String())
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("dummy", test.errorRequest.Name)
	test.Equal(true, test.errorRequest.ShowError)
	test.Equal(false, hasRun)
}

// TestGetHandlerFunc tests the HandleError when an error is returned from the service
func (test *TestSuite) TestGetHandlerFunc_responseErr() {
	var (
		err    error
		hasRun = false
	)

	test.engine.POST("/test", func(ctx *gin.Context) {
		GetHandlerFunc(test.ctx, test.request, func(requestCtx context.Context) (interface{}, error) {
			hasRun = true

			return nil, errors.New("an error")
		})
	})

	test.ctx.Request, err = http.NewRequest(http.MethodPost, "/test", nil)
	test.Nil(err)

	test.engine.ServeHTTP(test.recorder, test.ctx.Request)

	test.Equal(getErrorResponseString(), test.recorder.Body.String())
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("dummy", test.request.Name)
	test.Equal(true, hasRun)
}
