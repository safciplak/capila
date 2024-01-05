package templates

const ServiceTest = `package {{ .PrivateNameSingular }}Services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	{{ .PrivateNameSingular }}Models "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/models"
	{{ .PrivateNameSingular }}Repositories "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/repositories"
	"github.com/safciplak/{{ .ApplicationName }}/src/models"
)

// Test Suite which encapsulate the tests for the test.service
type TestSuite struct {
	suite.Suite
	ctx        context.Context
	repository *{{ .PrivateNameSingular }}Repositories.MockInterface{{ .PublicNameSingular }}Repository
	service    Interface{{ .PublicNameSingular }}Service
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	// Mocks used in the test
	test.repository = new({{ .PrivateNameSingular }}Repositories.MockInterface{{ .PublicNameSingular }}Repository)

	// Often used test objects
	test.ctx = context.TODO()

	// Object to be tested
	test.service = New{{ .PublicNameSingular }}Service(
		test.repository,
	)
}

// TearDownTest tests whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	test.repository.AssertExpectations(test.T())
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// get{{ .PublicNameSingular }}Struct builds a test example of a {{ .PrivateNameSingular }} object
func get{{ .PublicNameSingular }}Struct() models.{{ .PublicNameSingular }} {
	return models.{{ .PublicNameSingular }}{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		},
	}
}

// get{{ .PublicNameSingular }}CreateRequest builds a test example of a {{ .PrivateNameSingular }} create request
func get{{ .PublicNameSingular }}CreateRequest() {{ .PrivateNameSingular }}Models.CreateRequest {
	return {{ .PrivateNameSingular }}Models.CreateRequest{
		GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
	}
}

// get{{ .PublicNameSingular }}UpdateRequest builds a test example of a {{ .PrivateNameSingular }} update request
func get{{ .PublicNameSingular }}UpdateRequest() {{ .PrivateNameSingular }}Models.UpdateRequest {
	return {{ .PrivateNameSingular }}Models.UpdateRequest{
		GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43dde",
	}
}

// get{{ .PublicNameSingular }}BaseRequest builds a test example of a {{ .PrivateNameSingular }} base request
func get{{ .PublicNameSingular }}BaseRequest() {{ .PrivateNameSingular }}Models.BaseRequest {
	return {{ .PrivateNameSingular }}Models.BaseRequest{
		GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43dde",
	}
}

// TestList tests the happy flow for the List function
func (test *TestSuite) TestList() {
	var (
		expectedResult []models.{{ .PublicNameSingular }}
		{{ .PrivateNamePlural }}        []models.{{ .PublicNameSingular }}
		err            error
	)

	searchParams := &{{ .PrivateNameSingular }}Models.ListRequest{}

	{{ .PrivateNameSingular }}1 := get{{ .PublicNameSingular }}Struct()
	{{ .PrivateNameSingular }}2 := get{{ .PublicNameSingular }}Struct()
	expectedResult = append(expectedResult, {{ .PrivateNameSingular }}1, {{ .PrivateNameSingular }}2)

	test.repository.
		On("List", test.ctx, searchParams.ToQueryParams()).
		Return(expectedResult, nil).Once()

	{{ .PrivateNamePlural }}, err = test.service.List(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(2, len({{ .PrivateNamePlural }}))
}

// TestRead tests the happy flow for the Read function
func (test *TestSuite) TestRead() {
	expectedResult := get{{ .PublicNameSingular }}Struct()
	searchParams := &{{ .PrivateNameSingular }}Models.BaseRequest{
		GUID: expectedResult.GUID,
	}

	test.repository.
		On("Read", test.ctx, searchParams.ToQueryParams()).
		Return(&expectedResult, nil).Once()

	{{ .PrivateNameSingular }}, err := test.service.Read(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(expectedResult.GUID, {{ .PrivateNameSingular }}.GUID)
}

// TestCreate tests the happy flow for the Create function
func (test *TestSuite) TestCreate() {
	{{ .PrivateNameSingular }}Request := get{{ .PublicNameSingular }}CreateRequest()

	test.repository.
		On("Create", test.ctx, {{ .PrivateNameSingular }}Request.ToDBModel()).
		Return(nil).Once()

	{{ .PrivateNameSingular }}, err := test.service.Create(test.ctx, &{{ .PrivateNameSingular }}Request)

	test.Nil(err)
	test.Equal("b06225b2-0eea-4e1f-b514-9cb8f7a43ddf", {{ .PrivateNameSingular }}.GUID)

	// @TODO add tests to validate
}

// TestUpdate tests the happy flow for the Update function
func (test *TestSuite) TestUpdate() {
	{{ .PrivateNameSingular }}Request := get{{ .PublicNameSingular }}UpdateRequest()

	test.repository.
		On("Update", test.ctx, {{ .PrivateNameSingular }}Request.ToDBModel()).
		Return(nil).Once()

	{{ .PrivateNameSingular }}, err := test.service.Update(test.ctx, &{{ .PrivateNameSingular }}Request)

	test.Nil(err)
	test.Equal("b06225b2-0eea-4e1f-b514-9cb8f7a43dde", {{ .PrivateNameSingular }}.GUID)

	// @TODO: add tests to validate
}

// TestDelete tests the happy flow for the Delete function
func (test *TestSuite) TestDelete() {
	{{ .PrivateNameSingular }}Request := get{{ .PublicNameSingular }}BaseRequest()

	test.repository.
		On("Delete", test.ctx, {{ .PrivateNameSingular }}Request.ToDBModel(true)).
		Return(nil).Once()

	{{ .PrivateNameSingular }}, err := test.service.Delete(test.ctx, &{{ .PrivateNameSingular }}Request)

	test.Nil(err)
	test.Equal("b06225b2-0eea-4e1f-b514-9cb8f7a43dde", {{ .PrivateNameSingular }}.GUID)
}
`
