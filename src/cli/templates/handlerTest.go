package templates

const HandlerTest = `package {{ .PrivateNameSingular }}Handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/safciplak/capila/src/convert"

	{{ .PrivateNameSingular }}Models "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/models"
	{{ .PrivateNameSingular }}Services "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/services"
	"github.com/safciplak/{{ .ApplicationName }}/src/models"
)

// Test Suite which encapsulate the tests for the handler!
type TestSuite struct {
	suite.Suite
	ctx      context.Context
	router   *gin.Engine
	recorder *httptest.ResponseRecorder

	handler Interface{{ .PublicNameSingular }}Handler
	service *{{ .PrivateNameSingular }}Services.MockInterface{{ .PublicNameSingular }}Service

	{{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)

	test.ctx = context.TODO()
	test.router = gin.New()
	test.recorder = httptest.NewRecorder()
	test.service = new({{ .PrivateNameSingular }}Services.MockInterface{{ .PublicNameSingular }}Service)
	test.handler = New{{ .PublicNameSingular }}Handler(test.service)

	test.{{ .PrivateNameSingular }} = &models.{{ .PublicNameSingular }}{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		},
	}

	// Register the routes just like in the routes.go
	test.router.GET("/{{ .PrivateNamePlural }}", test.handler.List())
	test.router.POST("/{{ .PrivateNamePlural }}", test.handler.Create())
	test.router.GET("/{{ .PrivateNamePlural }}/:guid", test.handler.Read())
	test.router.PUT("/{{ .PrivateNamePlural }}/:guid", test.handler.Update())
	test.router.DELETE("/{{ .PrivateNamePlural }}/:guid", test.handler.Delete())
}

// TearDownTest asserts whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	test.service.AssertExpectations(test.T())
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// ge{{ .PublicNameSingular }}Struct builds a test example of a {{ .PrivateNameSingular }} object
func get{{ .PublicNameSingular }}Struct() models.{{ .PublicNameSingular }} {
	return models.{{ .PublicNameSingular }}{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43dd",
		},
	}
}

// get{{ .PublicNameSingular }}RequestBody returns a {{ .PrivateNameSingular }} JSON request
func get{{ .PublicNameSingular }}RequestBody() string {
	return "{\"guid\": \"b06225b2-0eea-4e1f-b514-9cb8f7a43ddf\"}"
}

// get{{ .PublicNameSingular }}InvalidRequestBody returns a {{ .PrivateNameSingular }} JSON request with invalid validation
func get{{ .PublicNameSingular }}InvalidRequestBody() string {
	return "{\"guid\": \"b06225b2-0eea-4e1f-b514-\"}"
}

// get{{ .PublicNameSingular }}CreateRequestStruct builds a test example of a {{ .PrivateNameSingular }} create object
func get{{ .PublicNameSingular }}CreateRequestStruct() *{{ .PrivateNameSingular }}Models.CreateRequest {
	return &{{ .PrivateNameSingular }}Models.CreateRequest{
		GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		Language: convert.NewString("EN"),
	}
}

// get{{ .PublicNameSingular }}UpdateRequestStruct builds a test example of a {{ .PrivateNameSingular }} update object
func get{{ .PublicNameSingular }}UpdateRequestStruct() *{{ .PrivateNameSingular }}Models.UpdateRequest {
	return &{{ .PrivateNameSingular }}Models.UpdateRequest{
		GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		Language: convert.NewString("EN"),
	}
}

// TestList tests the happy flow for the List function
func (test *TestSuite) TestList() {
	expectedResult := make([]models.{{ .PublicNameSingular }}, 0)
	{{ .PrivateNameSingular }}1 := get{{ .PublicNameSingular }}Struct()
	{{ .PrivateNameSingular }}2 := get{{ .PublicNameSingular }}Struct()
	expectedResult = append(expectedResult, {{ .PrivateNameSingular }}1, {{ .PrivateNameSingular }}2)

	searchParams := {{ .PrivateNameSingular }}Models.ListRequest{
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/{{ .PrivateNamePlural }}", nil)
	test.Nil(err)

	test.service.On("List", test.ctx, &searchParams).Return(expectedResult, nil).Once()
	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestReadValidationError tests the validation handling of the Read function
func (test *TestSuite) TestListValidationError() {
	searchParams := {{ .PrivateNameSingular }}Models.ListRequest{
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/{{ .PrivateNamePlural }}", nil)
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusBadRequest, test.recorder.Code)

	// @TODO: add tests: (e.g. test.Equal("{\"data\":{\"Name\":\"min\"}}", test.recorder.Body.String()))
}

// TestListServiceError tests the service error handling of the List function
func (test *TestSuite) TestListServiceError() {
	searchParams := {{ .PrivateNameSingular }}Models.ListRequest{
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/{{ .PrivateNamePlural }}", nil)
	test.Nil(err)

	expectedError := errors.New("unknown error")
	test.service.On("List", test.ctx, &searchParams).Return(nil, expectedError).Once()
	test.router.ServeHTTP(test.recorder, request)

	// The correct statusCode has been given
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestRead tests the happy flow for the Read function
func (test *TestSuite) TestRead() {
	searchParams := {{ .PrivateNameSingular }}Models.BaseRequest{
		GUID:     test.{{ .PrivateNameSingular }}.GUID,
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, nil)
	test.Nil(err)

	test.service.
		On("Read", test.ctx, &searchParams).
		Return(test.{{ .PrivateNameSingular }}, nil).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestReadValidationError tests the validation handling of the Read function
func (test *TestSuite) TestReadValidationError() {
	searchParams := {{ .PrivateNameSingular }}Models.BaseRequest{
		GUID: "WRONG-GUID",
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/{{ .PrivateNamePlural }}/"+searchParams.GUID, nil)
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"GUID\":\"uuid\"}}", test.recorder.Body.String())
}

// TestReadServiceError tests the service error handling of the Read function
func (test *TestSuite) TestReadServiceError() {
	searchParams := {{ .PrivateNameSingular }}Models.BaseRequest{
		GUID:     test.{{ .PrivateNameSingular }}.GUID,
		Language: convert.NewString("EN"),
	}

	request, err := http.NewRequestWithContext(test.ctx, "GET", "/{{ .PrivateNamePlural }}/"+searchParams.GUID, nil)
	test.Nil(err)

	expectedError := errors.New("unknown error")

	test.service.
		On("Read", test.ctx, &searchParams).
		Return(nil, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestCreate tests the happy flow for the Create function
func (test *TestSuite) TestCreate() {
	requestBody := get{{ .PublicNameSingular }}RequestBody()
	{{ .PrivateNameSingular }} := get{{ .PublicNameSingular }}CreateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "POST", "/{{ .PrivateNamePlural }}", bytes.NewBufferString(requestBody))
	test.Nil(err)

	test.service.
		On("Create", test.ctx, {{ .PrivateNameSingular }}).
		Return(test.{{ .PrivateNameSingular }}, nil).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestCreateValidationError tests the validation handling of the Create function
func (test *TestSuite) TestCreateValidationError() {
	requestBody := get{{ .PublicNameSingular }}InvalidRequestBody()

	request, err := http.NewRequestWithContext(test.ctx, "POST", "/{{ .PrivateNamePlural }}", bytes.NewBufferString(requestBody))
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	// @TODO add tests here (e.g. test.Equal("{\"data\":{\"Name\":\"min\"}}", test.recorder.Body.String()) )
}

// TestCreateServiceError tests the service error handling of the Create function
func (test *TestSuite) TestCreateServiceError() {
	requestBody := get{{ .PublicNameSingular }}RequestBody()
	{{ .PrivateNameSingular }} := get{{ .PublicNameSingular }}CreateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "POST", "/{{ .PrivateNamePlural }}", bytes.NewBufferString(requestBody))
	test.Nil(err)

	expectedError := errors.New("unknown error")
	test.service.
		On("Create", test.ctx, {{ .PrivateNameSingular }}).
		Return(nil, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestUpdate tests the happy flow for the Update function
func (test *TestSuite) TestUpdate() {
	requestBody := get{{ .PublicNameSingular }}RequestBody()
	{{ .PrivateNameSingular }} := get{{ .PublicNameSingular }}UpdateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, bytes.NewBufferString(requestBody))
	test.Nil(err)

	test.service.
		On("Update", test.ctx, {{ .PrivateNameSingular }}).
		Return(test.{{ .PrivateNameSingular }}, nil).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestUpdateValidationError tests the validation handling of the Update function
func (test *TestSuite) TestUpdateValidationError() {
	requestBody := get{{ .PublicNameSingular }}InvalidRequestBody()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, bytes.NewBufferString(requestBody))
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	// @TODO add tests here (e.g. test.Equal("{\"data\":{\"Name\":\"min\"}}", test.recorder.Body.String()) )
}

// TestUpdateValidationErrorWrongGUID tests the validation handling of the Update function
func (test *TestSuite) TestUpdateValidationErrorWrongGUID() {
	test.{{ .PrivateNameSingular }}.GUID = "WRONG-GUID"
	requestBody := get{{ .PublicNameSingular }}RequestBody()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, bytes.NewBufferString(requestBody))
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"GUID\":\"uuid\"}}", test.recorder.Body.String())
}

// TestUpdateServiceError tests the service error handling of the Update function
func (test *TestSuite) TestUpdateServiceError() {
	requestBody := get{{ .PublicNameSingular }}RequestBody()
	{{ .PrivateNameSingular }} := get{{ .PublicNameSingular }}UpdateRequestStruct()

	request, err := http.NewRequestWithContext(test.ctx, "PUT", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, bytes.NewBufferString(requestBody))
	test.Nil(err)

	expectedError := errors.New("unknown error")
	test.service.
		On("Update", test.ctx, {{ .PrivateNameSingular }}).
		Return(nil, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}

// TestDelete tests the happy flow for the Delete function
func (test *TestSuite) TestDelete() {
	test.{{ .PrivateNameSingular }}.IsDeleted = true

	request, err := http.NewRequestWithContext(test.ctx, "DELETE", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, nil)
	test.Nil(err)

	test.service.
		On("Delete", test.ctx, &{{ .PrivateNameSingular }}Models.BaseRequest{
			GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
			Language: convert.NewString("EN"),
		}).
		Return(test.{{ .PrivateNameSingular }}, nil).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusOK, test.recorder.Code)
}

// TestDeleteValidationError tests the validation handling of the Delete function
func (test *TestSuite) TestDeleteValidationError() {
	test.{{ .PrivateNameSingular }}.GUID = "WRONG-GUID"
	test.{{ .PrivateNameSingular }}.IsDeleted = true

	request, err := http.NewRequestWithContext(test.ctx, "DELETE", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, nil)
	test.Nil(err)

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusBadRequest, test.recorder.Code)
	test.Equal("{\"data\":{\"GUID\":\"uuid\"}}", test.recorder.Body.String())
}

// TestDeleteServiceError tests the service error handling of the Delete function
func (test *TestSuite) TestDeleteServiceError() {
	test.{{ .PrivateNameSingular }}.IsDeleted = true

	request, err := http.NewRequestWithContext(test.ctx, "DELETE", "/{{ .PrivateNamePlural }}/"+test.{{ .PrivateNameSingular }}.GUID, nil)
	test.Nil(err)

	expectedError := errors.New("unknown error")

	test.service.
		On("Delete", test.ctx, &{{ .PrivateNameSingular }}Models.BaseRequest{
			GUID:     "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
			Language: convert.NewString("EN"),
		}).
		Return(test.{{ .PrivateNameSingular }}, expectedError).Once()

	test.router.ServeHTTP(test.recorder, request)
	test.Equal(http.StatusInternalServerError, test.recorder.Code)
	test.Equal("{\"data\":\"unknown error\"}", test.recorder.Body.String())
}`
