package response

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"

	capilaErrors "github.com/safciplak/capila/src/errors"
	"github.com/safciplak/capila/src/http/presenters"
)

type TestSuite struct {
	suite.Suite
	ctx           context.Context
	router        *gin.Engine
	recorder      *httptest.ResponseRecorder
	mockPresenter *presenters.MockInterfacePresenter
	response      *Response
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// SetupTest sets up the test
func (test *TestSuite) SetupTest() {
	test.ctx = context.Background()
	test.router = gin.New()
	test.recorder = httptest.NewRecorder()
	test.mockPresenter = presenters.NewMockInterfacePresenter(test.T())
	test.response = Create()
}

// TestSuccess tests Success
func (test *TestSuite) TestSuccess() {
	test.response = Create()
}

func (test *TestSuite) TestErrorWrapping() {
	//nolint:govet // legibility precedes performance
	var tests = []struct {
		name        string
		inputError  error
		outputError error
		statusCode  int
	}{
		{
			"No error provided", nil, nil, 200,
		},
		{
			"Unhandled error slipped through",
			errors.New("Oh noes"), capilaErrors.NewErrorInternalServerError(), 500,
		},
		{
			"Postgres no results",
			pg.ErrNoRows, capilaErrors.NewErrorInternalServerError(), 500,
		},
		{
			"Validation error",
			validator.ValidationErrors{}, capilaErrors.NewErrorBadRequest(), 400,
		},
		{
			"Bad gateway error provided",
			capilaErrors.NewErrorBadGateway(), capilaErrors.NewErrorBadGateway(), 502,
		},
		{
			"Bad request error provided",
			capilaErrors.NewErrorBadRequest(), capilaErrors.NewErrorBadRequest(), 400,
		},
		{
			"Gateway timeout error provided",
			capilaErrors.NewErrorGatewayTimeout(), capilaErrors.NewErrorGatewayTimeout(), 504,
		},
		{
			"Internal server error provided",
			capilaErrors.NewErrorInternalServerError(), capilaErrors.NewErrorInternalServerError(), 500,
		},
		{
			"Not Found error provided",
			capilaErrors.NewErrorNotFound(), capilaErrors.NewErrorNotFound(), 404,
		},
		{
			"Service unavailable error provided",
			capilaErrors.NewErrorServiceUnavailable(), capilaErrors.NewErrorServiceUnavailable(), 503,
		},
	}

	for _, iterator := range tests {
		testCase := iterator

		test.Run(testCase.name, func() {
			test.response.CheckForErrors(test.ctx, nil, testCase.inputError)
			test.ErrorIs(testCase.outputError, test.response.Error)
			test.Equal(testCase.statusCode, test.response.StatusCode)
		})
	}
}

// TestHasErrors tests HasErrors
func (test *TestSuite) TestHasErrors() {
	test.response.Error = nil
	test.Equal(false, test.response.HasErrors())

	test.response.SetError(test.ctx, errors.New("unhandled error"))
	test.Equal(true, test.response.HasErrors())
}

func (test *TestSuite) TestHasValidationErrors() {
	test.response.Error = nil
	test.Equal(false, test.response.HasValidationErrors())

	test.response.SetError(test.ctx, errors.New("unhandled error"))
	test.Equal(false, test.response.HasValidationErrors())

	test.response.SetError(test.ctx, capilaErrors.NewErrorBadRequest().Wrap(errors.New("custom validation error")))
	test.Equal(true, test.response.HasValidationErrors())

	test.response.SetError(test.ctx, validator.ValidationErrors{})
	test.Equal(true, test.response.HasValidationErrors())
}

// TestHandleValidationErrorWithDifferentError
func (test *TestSuite) TestHandleValidationErrorWithDifferentError() {
	var nonValidationError = errors.New("no validation error but something else")

	test.response.HandleValidationError(test.ctx, nonValidationError)

	test.ErrorIs(test.response.Error, capilaErrors.NewErrorBadRequest())
	test.Equal(http.StatusBadRequest, test.response.StatusCode)
}

// TestHandleValidationErrorWithDifferentError
func (test *TestSuite) TestHandleErrorWithDifferentError() {
	var nonValidationError = errors.New("generic error")

	test.response.HandleError(test.ctx, nonValidationError)

	test.ErrorIs(test.response.Error, capilaErrors.NewErrorInternalServerError())
	test.Equal(http.StatusInternalServerError, test.response.StatusCode)
}

// TestResponseHasValidationErrors tests HasValidationErrors with a validation error.
func (test *TestSuite) TestResponseHasValidationErrors() {
	var body []byte

	test.router.GET("/:guid", func(ctx *gin.Context) {
		var testRequest = struct {
			GUID string `binding:"required,uuid" json:"guid"`
		}{}

		Create().HandleValidationError(ctx, ctx.Bind(&testRequest)).ReturnJSON(ctx)
	})

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/wrong-guid", nil)
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)

	body, err = ioutil.ReadAll(test.recorder.Body)
	test.Nil(err)

	// The correct statusCode has been given
	test.Equal(http.StatusBadRequest, test.recorder.Code)

	var result = `{"errors":[` +
		`{"code":"ERROR_BAD_REQUEST","title":"errors.ErrorBadRequest","status":400},` +
		`{"code":"ERROR_INPUT_VALIDATION","title":"validator.fieldError","detail":"Key: 'GUID' Error:Field validation for 'GUID' failed on the 'required' tag"}],` +
		`"_links":{"self":{"href":"/wrong-guid"}},"meta":{"query":{}}}`

	test.Equal(string(body), result)
}

// TestOutputError tests the error output with a presenter
func (test *TestSuite) TestOutputError() {
	var expectedError = capilaErrors.NewErrorInternalServerError()

	test.response.HandleError(test.ctx, expectedError)
	test.mockPresenter.On("Error", expectedError.GetStatusCode(), expectedError).Once().Return()

	test.response.Output(test.mockPresenter)
}

// TestOutputSuccess tests the success output with a presenter
func (test *TestSuite) TestOutputSuccess() {
	test.response.Data = "Test created"
	test.response.StatusCode = 202

	test.mockPresenter.On("Success", test.response.StatusCode, test.response.Data).Once().Return()

	test.response.Output(test.mockPresenter)
}
