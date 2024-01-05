package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	capilaContext "github.com/safciplak/capila/src/http/context"
)

// LanguageTestSuite contains all the helper test specific properties
type LanguageTestSuite struct {
	suite.Suite

	router   *gin.Engine
	recorder *httptest.ResponseRecorder

	reservationRequest *http.Request
}

// TestLanguageTestSuite runs the testsuite
func TestLanguageTestSuite(t *testing.T) {
	suite.Run(t, new(LanguageTestSuite))
}

// SetupTest initializes the environment in which the service will run
func (test *LanguageTestSuite) SetupTest() {
	test.router = gin.New()
	test.router.Use(LanguageMiddleware())

	test.recorder = httptest.NewRecorder()

	test.reservationRequest, _ = http.NewRequestWithContext(
		context.TODO(),
		http.MethodGet,
		"/reservations/c0da7fd4-484d-42a9-8204-ca5bc9e6ec5b",
		http.NoBody,
	)
}

// TestAcceptLanguageHeader tests the middleware when a accept-language is given.
func (test *LanguageTestSuite) TestAcceptLanguageHeader() {
	test.router.GET("/reservations/:guid", func(c *gin.Context) {
		test.Equal("nl", capilaContext.GetTwoLetterLanguageCode(c.Request.Context()))
	})

	test.reservationRequest.Header.Add("Accept-Language", "nl")
	test.router.ServeHTTP(test.recorder, test.reservationRequest)

	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestFallbackForLanguageHeader tests the middleware when no accept-language is given.
func (test *LanguageTestSuite) TestFallbackForLanguageHeader() {
	test.router.GET("/reservations/:guid", func(c *gin.Context) {
		test.Equal("en", capilaContext.GetTwoLetterLanguageCode(c.Request.Context()))
	})

	test.router.ServeHTTP(test.recorder, test.reservationRequest)

	test.Equal(http.StatusOK, test.recorder.Code)
}
