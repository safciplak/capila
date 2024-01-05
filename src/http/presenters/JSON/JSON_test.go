package JSON

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"

	capilaErrors "github.com/safciplak/capila/src/errors"
)

type TestSuite struct {
	suite.Suite

	ctx      context.Context
	router   *gin.Engine
	recorder *httptest.ResponseRecorder

	presenter Presenter
}

// TestTestSuite runs the actual tests
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// SetupTest sets up the test
func (test *TestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)

	test.ctx = context.TODO()
	test.router = gin.Default()
	test.recorder = httptest.NewRecorder()

	test.presenter = Presenter{}
}

// TearDownTest tears down the test
func (test *TestSuite) TearDownTest() {
}

// Test_MockSuccessResponse tests a success response
func (test *TestSuite) Test_MockSuccessResponse() {
	var (
		result Response
		path   = "/ping?arrivalDate=01-01-2019&departureDate=10-01-2019"
	)

	test.router.GET("/ping", func(ctx *gin.Context) {
		var presenter = Present(ctx)

		presenter.Success(http.StatusOK, "ok!")
	})

	req, _ := http.NewRequestWithContext(test.ctx, http.MethodGet, path, nil)
	test.router.ServeHTTP(test.recorder, req)

	if err := json.NewDecoder(test.recorder.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	test.Equal(path, result.Links["self"].Href)
	test.Equal("01-01-2019", result.Meta.Query.Get("arrivalDate"))
	test.Equal("10-01-2019", result.Meta.Query.Get("departureDate"))
	test.Equal(200, test.recorder.Code)
}

// Test_MockErrorResponse tests the error response
func (test *TestSuite) Test_MockErrorResponse() {
	var (
		result      Response
		customError = "You've got an error error!"
		path        = "/ping?arrivalDate=01-01-2019&departureDate=10-01-2019"
		errorCode   = http.StatusInternalServerError
	)

	test.router.GET("/ping", func(ctx *gin.Context) {
		var presenter = Present(ctx)

		presenter.Error(errorCode, errors.New(customError))
	})

	req, _ := http.NewRequestWithContext(test.ctx, http.MethodGet, path, nil)

	test.router.ServeHTTP(test.recorder, req)

	if err := json.NewDecoder(test.recorder.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	test.Equal(result.Errors[0], ErrorObject{
		ID:     "",
		Status: 0,
		Code:   "ERROR_UNKNOWN",
		Title:  "errors.errorString",
		Detail: "You've got an error error!",
	})
	test.Equal(errorCode, test.recorder.Code)
}

// Test_MockBadRequestResponse tests a bad request response
func (test *TestSuite) Test_MockBadRequestResponse() {
	var (
		result    Response
		path      = "/ping?arrivalDate=01-01-2019&departureDate=10-01-2019"
		errorCode = http.StatusBadRequest
		validate  = validator.New()
	)

	type dummyStruct struct {
		FirstName     string `validate:"required"`
		LastName      string `validate:"required"`
		Street        string `validate:"min=3"`
		HouseAddition string `validate:"max=3"`
	}

	dummy := dummyStruct{
		HouseAddition: "kelder",
	}
	validationErrors := validate.Struct(dummy)

	test.router.GET("/ping", func(ctx *gin.Context) {
		var presenter = Present(ctx)

		presenter.Error(errorCode, validationErrors)
	})

	req, _ := http.NewRequestWithContext(test.ctx, http.MethodGet, path, nil)

	test.router.ServeHTTP(test.recorder, req)

	if err := json.NewDecoder(test.recorder.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	test.Equal(errorCode, test.recorder.Code)
}

// TestAddLinkObject tests the working of adding an item to the link element
func (test *TestSuite) TestAddLinkObject() {
	var (
		result Response
		path   = "/ping?arrivalDate=01-01-2019&departureDate=10-01-2019"
	)

	test.router.GET("/ping", func(ctx *gin.Context) {
		var presenter = Present(ctx)

		presenter.AddLinkObject("first-item", "/ping?item=first")
		presenter.AddLinkObject("second-item", "/ping?item=second")

		presenter.Success(200, "")
	})

	req, _ := http.NewRequestWithContext(test.ctx, http.MethodGet, path, nil)

	test.router.ServeHTTP(test.recorder, req)

	if err := json.NewDecoder(test.recorder.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	test.Equal(path, result.Links["self"].Href)
	test.Equal("/ping?item=first", result.Links["first-item"].Href)
	test.Equal("/ping?item=second", result.Links["second-item"].Href)

	test.Equal(200, test.recorder.Code)
}

// TestWithoutLinks tests the working of not adding an item to the link element
func (test *TestSuite) TestWithoutLinks() {
	var (
		result Response
		path   = "/ping?arrivalDate=01-01-2019&departureDate=10-01-2019"
	)

	test.router.GET("/ping", func(ctx *gin.Context) {
		var presenter = Present(ctx)

		presenter.Success(200, "")
	})

	req, _ := http.NewRequestWithContext(test.ctx, http.MethodGet, path, nil)

	test.router.ServeHTTP(test.recorder, req)

	if err := json.NewDecoder(test.recorder.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	test.Equal(path, result.Links["self"].Href)
	test.Equal(1, len(result.Links))
	test.Equal(200, test.recorder.Code)
}

func (test *TestSuite) TestConvertErrorToErrorObjectsNumError() {
	var result = test.presenter.convertErrorToErrorObjects(capilaErrors.NewErrorBadRequest().Wrap(&strconv.NumError{
		Func: "Test",
		Num:  "Test",
		Err:  errors.New("invalid input"),
	}))

	test.Equal(result[0].Title, "errors.ErrorBadRequest")
	test.Equal(result[0].Code, "ERROR_BAD_REQUEST")
	test.Equal(result[0].Detail, "")
	test.Equal(result[0].ID, "")
	test.Equal(result[0].Status, 400)

	test.Equal(result[1].Title, "strconv.NumError")
	test.Equal(result[1].Code, "ERROR_INPUT_VALIDATION")
	test.Equal(result[1].Detail, "strconv.Test: parsing \"Test\": invalid input")
	test.Equal(result[1].ID, "")
	test.Equal(result[1].Status, 0)
}

func (test *TestSuite) TestConvertErrorToErrorObjectsValidationError() {
	var result = test.presenter.convertErrorToErrorObjects(capilaErrors.NewErrorBadRequest().Wrap(validator.ValidationErrors{
		MockFieldError{},
		MockFieldError{},
	}))

	test.Equal(result[0].Title, "errors.ErrorBadRequest")
	test.Equal(result[0].Code, "ERROR_BAD_REQUEST")
	test.Equal(result[0].Detail, "")
	test.Equal(result[0].ID, "")
	test.Equal(result[0].Status, 400)

	test.Equal(result[1].Title, "JSON.MockFieldError")
	test.Equal(result[1].Code, "ERROR_INPUT_VALIDATION")
	test.Equal(result[1].Detail, "Tested")
	test.Equal(result[1].ID, "")
	test.Equal(result[1].Status, 0)
}

// Due to validator v10 not exposing their concrete types we have to recreate them...
type MockFieldError struct {
}

func (MockFieldError) Tag() string {
	panic("implement me")
}

func (MockFieldError) ActualTag() string {
	panic("implement me")
}

func (MockFieldError) Namespace() string {
	panic("implement me")
}

func (MockFieldError) StructNamespace() string {
	panic("implement me")
}

func (MockFieldError) Field() string {
	panic("implement me")
}

func (MockFieldError) StructField() string {
	panic("implement me")
}

func (MockFieldError) Value() interface{} {
	panic("implement me")
}

func (MockFieldError) Param() string {
	panic("implement me")
}

func (MockFieldError) Kind() reflect.Kind {
	panic("implement me")
}

func (MockFieldError) Type() reflect.Type {
	panic("implement me")
}

func (MockFieldError) Translate(_ ut.Translator) string {
	return "Tested"
}

func (MockFieldError) Error() string {
	panic("implement me")
}
