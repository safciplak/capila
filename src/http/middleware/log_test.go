package middlewares

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/safciplak/capila/src/logger"
)

const LogLevel = "Info"

const LogKeyRequest = "request"
const LogKeyResponse = "response"

const LogKeyURL = "url"
const LogKeyStatus = "status"
const LogKeyMethod = "method"
const LogKeyHeaders = "headers"
const LogKeyBody = "body"

// LogTestSuite contains all the helper test specific properties
type LogTestSuite struct {
	suite.Suite

	router   *gin.Engine
	recorder *httptest.ResponseRecorder
	log      *logger.MockInterfaceLogger
	zapLog   *logger.MockInterfaceZapLogger

	request *http.Request
}

// TestLanguageTestSuite runs the testsuite
func TestLogTestSuite(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}

func (test *LogTestSuite) SetupTest() {
	test.log = &logger.MockInterfaceLogger{}
	test.zapLog = &logger.MockInterfaceZapLogger{}

	test.router = gin.New()
	test.router.Use(LogMiddleware(test.log))
	test.router.POST("/reservations/:guid", func(c *gin.Context) {
		c.Header("X-Custom-Response-Header", "Custom value")
		c.JSON(http.StatusOK, getJSONResponse())
	})

	test.recorder = httptest.NewRecorder()

	test.request, _ = http.NewRequestWithContext(
		context.TODO(),
		http.MethodPost,
		"/reservations/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b",
		bytes.NewBufferString(getRequestBody()),
	)
	test.request.Header.Add("X-Custom-Request-Header", "Value")
}

// getRequestBody Returns a simple json object with a couple of values
func getRequestBody() string {
	return `
	{
		"stringProperty": "value",
		"floatProperty": 1.1,
		"integerProperty": 1,
		"nullProperty": null,
		"emptyObjectProperty": {},
		"emptyArrayProperty": [],
		"stringArray": ["a", "b"]
	}`
}

// getJSONResponse Returns a rather simple JSON object
func getJSONResponse() interface{} {
	return struct {
		A string `json:"a,omitempty"`
		B string `json:"b,omitempty"`
	}{
		"A Value",
		"B Value",
	}
}

// Returns a string representation of the Simple header values
func getHeaderStrings() []string {
	return []string{"Content-Type: application/json; charset=utf-8", "X-Custom-Response-Header: Custom value"}
}

// TearDownTest removes all side effects after the suite has been completed
func (test *LogTestSuite) TearDownTest() {
	test.log.AssertExpectations(test.T())
	test.zapLog.AssertExpectations(test.T())
}

// TestHappy tests
func (test *LogTestSuite) TestHappy() {
	test.log.On("Log", test.request.Context()).Return(test.zapLog).Once()

	test.zapLog.On(LogLevel, LogKeyRequest,
		zap.String(LogKeyURL, "/reservations/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b"),
		zap.String(LogKeyMethod, http.MethodPost),
		zap.Strings(LogKeyHeaders, []string{"X-Custom-Request-Header: Value"}),
		zap.String(LogKeyBody, getRequestBody()),
	).Return(test.zapLog).Once()
	test.zapLog.On(LogLevel, LogKeyResponse,
		zap.Int(LogKeyStatus, http.StatusOK),
		zap.Strings(LogKeyHeaders, getHeaderStrings()),
		zap.String(LogKeyBody, `{"a":"A Value","b":"B Value"}`),
	).Return(test.zapLog).Once()

	test.router.ServeHTTP(test.recorder, test.request)

	test.Equal(http.StatusOK, test.recorder.Code)
}

func (test *LogTestSuite) TestEmptyRequestBodyLog() {
	test.request, _ = http.NewRequestWithContext(
		context.TODO(),
		http.MethodPost,
		"/reservations/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b",
		bytes.NewBufferString(""),
	)

	test.log.On("Log", test.request.Context()).Return(test.zapLog).Once()

	test.zapLog.On(LogLevel, LogKeyRequest,
		zap.String(LogKeyURL, "/reservations/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b"),
		zap.String(LogKeyMethod, http.MethodPost),
		zap.Strings(LogKeyHeaders, []string{}),
		zap.String(LogKeyBody, ""),
	).Return(test.zapLog).Once()
	test.zapLog.On(LogLevel, LogKeyResponse,
		zap.Int(LogKeyStatus, http.StatusOK),
		zap.Strings(LogKeyHeaders, getHeaderStrings()),
		zap.String(LogKeyBody, `{"a":"A Value","b":"B Value"}`),
	).Return(test.zapLog).Once()

	test.router.ServeHTTP(test.recorder, test.request)

	test.Equal(http.StatusOK, test.recorder.Code)
}

func (test *LogTestSuite) TestEmptyResponseBodyLog() {
	test.request, _ = http.NewRequestWithContext(
		context.TODO(),
		http.MethodPost,
		"/empty/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b?a=1",
		bytes.NewBufferString(getRequestBody()),
	)
	test.router.POST("/empty/:guid", func(c *gin.Context) {
		c.Header("X-Custom-Response-Header", "Custom value")
		c.JSON(http.StatusOK, struct {
			A string `json:"a,omitempty"`
			B string `json:"b,omitempty"`
		}{})
	})

	test.log.On("Log", test.request.Context()).Return(test.zapLog).Once()
	test.zapLog.On(LogLevel, LogKeyRequest,
		zap.String(LogKeyURL, "/empty/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b?a=1"),
		zap.String(LogKeyMethod, http.MethodPost),
		zap.Strings(LogKeyHeaders, []string{}),
		zap.String(LogKeyBody, getRequestBody()),
	).Return(test.zapLog).Once()
	test.zapLog.On(LogLevel, LogKeyResponse,
		zap.Int(LogKeyStatus, http.StatusOK),
		zap.Strings(LogKeyHeaders, getHeaderStrings()),
		zap.String(LogKeyBody, `{}`),
	).Return(test.zapLog).Once()

	test.router.ServeHTTP(test.recorder, test.request)

	test.Equal(http.StatusOK, test.recorder.Code)
}

func (test *LogTestSuite) TestErrorResponseBodyLog() {
	test.request, _ = http.NewRequestWithContext(
		context.TODO(),
		http.MethodPost,
		"/error/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b",
		bytes.NewBufferString(getRequestBody()),
	)

	test.router.POST("/error/:guid", func(c *gin.Context) {
		c.Header("X-Custom-Response-Header", "Custom value")
		c.JSON(http.StatusBadRequest, struct {
			A string `json:"a,omitempty"`
			B string `json:"b,omitempty"`
		}{})
	})

	test.log.On("Log", test.request.Context()).Return(test.zapLog).Once()
	test.zapLog.On(LogLevel, LogKeyRequest,
		zap.String(LogKeyURL, "/error/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b"),
		zap.String(LogKeyMethod, http.MethodPost),
		zap.Strings(LogKeyHeaders, []string{}),
		zap.String(LogKeyBody, getRequestBody()),
	).Return(test.zapLog).Once()

	test.zapLog.On(LogLevel, LogKeyResponse,
		zap.Int(LogKeyStatus, http.StatusBadRequest),
		zap.Strings(LogKeyHeaders, getHeaderStrings()),
		zap.String(LogKeyBody, `{}`),
	).Return(test.zapLog).Once()

	test.router.ServeHTTP(test.recorder, test.request)

	test.Equal(http.StatusBadRequest, test.recorder.Code)
}
