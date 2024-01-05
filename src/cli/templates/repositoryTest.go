package templates

const RepositoryTest = `package {{ .PrivateNameSingular }}Repositories

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/safciplak/capila/src/database"

	{{ .PrivateNameSingular }}Models "github.com/safciplak/{{ .ApplicationName }}/src/business/{{ .PrivateNamePlural }}/models"
	"github.com/safciplak/{{ .ApplicationName }}/src/models"
)

// Test Suite which encapsulate the tests for the repository
type TestSuite struct {
	suite.Suite

	ctx   context.Context
	{{ .PrivateNameSingular }} *models.{{ .PublicNameSingular }}
	now   time.Time

	readConn  *database.MockInterfacePGDB
	writeConn *database.MockInterfacePGDB
	query     *database.MockInterfaceORMQuery

	repository Interface{{ .PublicNameSingular }}Repository
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	// Mocks used in the test
	test.readConn = new(database.MockInterfacePGDB)
	test.writeConn = new(database.MockInterfacePGDB)
	test.query = new(database.MockInterfaceORMQuery)

	// Often used test objects
	test.{{ .PrivateNameSingular }} = get{{ .PublicNameSingular }}Struct()
	test.now = time.Now().UTC()
	test.ctx = context.TODO()

	// Object to be tested
	test.repository = New{{ .PublicNameSingular }}Repository(
		&database.Connection{
			Read:  test.readConn,
			Write: test.writeConn,
		},
	)
}

// TearDownTest asserts whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	test.readConn.AssertExpectations(test.T())
	test.writeConn.AssertExpectations(test.T())
	test.query.AssertExpectations(test.T())
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// ge{{ .PublicNameSingular }}Struct builds a test example of a {{ .PrivateNameSingular }} object
func get{{ .PublicNameSingular }}Struct() *models.{{ .PublicNameSingular }} {
	return &models.{{ .PublicNameSingular }}{
		BaseTableModel: models.BaseTableModel{
			GUID: "b06225b2-0eea-4e1f-b514-9cb8f7a43ddf",
		},
	}
}

// TestList tests the happy flow for the List function
func (test *TestSuite) TestList() {
	var (
		expectedResult = make([]models.{{ .PublicNameSingular }}, 0)
		{{ .PrivateNamePlural }}        = make([]models.{{ .PublicNameSingular }}, 0)
	)

	searchParams := &{{ .PrivateNameSingular }}Models.QueryParams{}

	{{ .PrivateNameSingular }}1 := get{{ .PublicNameSingular }}Struct()
	{{ .PrivateNameSingular }}2 := get{{ .PublicNameSingular }}Struct()
	expectedResult = append(expectedResult, *{{ .PrivateNameSingular }}1, *{{ .PrivateNameSingular }}2)

	test.readConn.On("ModelContext", test.ctx, &{{ .PrivateNamePlural }}).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*[]models.{{ .PublicNameSingular }})
			*arg = expectedResult
		}).
		Return(test.query).
		Once()
	test.query.On("Where", "{{ .PrivateNameSingular }}.isdeleted = false").Return(test.query).Once()
	test.query.On("Select").Return(nil).Once()

	data, err := test.repository.List(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(2, len(data))
}

// TestRead tests the happy flow for the Read function
func (test *TestSuite) TestRead() {
	searchParams := &{{ .PrivateNameSingular }}Models.QueryParams{
		GUID: test.{{ .PrivateNameSingular }}.GUID,
	}

	test.readConn.On("ModelContext", test.ctx, &models.{{ .PublicNameSingular }}{}).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.{{ .PublicNameSingular }})
			*arg = *test.{{ .PrivateNameSingular }}
		}).
		Return(test.query).
		Once()
	test.query.On("Where", "{{ .PrivateNameSingular }}.guid = ?", searchParams.GUID).Return(test.query).Once()
	test.query.On("Where", "{{ .PrivateNameSingular }}.isdeleted = false").Return(test.query).Once()
	test.query.On("First").Return(nil).Once()

	data, err := test.repository.Read(test.ctx, searchParams)

	test.Nil(err)
	test.Equal(test.{{ .PrivateNameSingular }}.GUID, data.GUID)
}

// TestCreate tests the happy flow for the Create function
func (test *TestSuite) TestCreate() {
	test.writeConn.On("ModelContext", test.ctx, test.{{ .PrivateNameSingular }}).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*models.{{ .PublicNameSingular }})
			*arg = *test.{{ .PrivateNameSingular }}
			arg.CreatedAt = test.now.String()
		}).Return(test.query).Once()
	test.query.On("Returning", "id, createdby, createdat, updatedby, updatedat, guid, isdeleted").Return(test.query).Once()
	test.query.On("Insert").Return(nil, nil).Once()

	err := test.repository.Create(test.ctx, test.{{ .PrivateNameSingular }})

	test.Nil(err)
	test.Equal(test.now.String(), test.{{ .PrivateNameSingular }}.CreatedAt)
}

// TestUpdate tests the happy flow for the Update function
func (test *TestSuite) TestUpdate() {
	test.writeConn.On("ModelContext", test.ctx, test.{{ .PrivateNameSingular }}).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*models.{{ .PublicNameSingular }})
		*arg = *test.{{ .PrivateNameSingular }}
		arg.UpdatedAt = test.now.String()
	}).Return(test.query).Once()
	test.query.On("Column", "").Return(test.query).Once()  // @TODO: select columns to update
	test.query.On("Returning", "id, createdby, createdat, updatedby, updatedat, guid, isdeleted").Return(test.query).Once()
	test.query.On("Where", "{{ .PrivateNameSingular }}.guid = ?", test.{{ .PrivateNameSingular }}.GUID).Return(test.query).Once()
	test.query.On("Where", "{{ .PrivateNameSingular }}.isdeleted = false").Return(test.query).Once()
	test.query.On("Update").Return(nil, nil).Once()

	err := test.repository.Update(test.ctx, test.{{ .PrivateNameSingular }})

	test.Nil(err)
	test.Equal(test.now.String(), test.{{ .PrivateNameSingular }}.UpdatedAt)
}

// TestDelete tests the happy flow for the Delete function
func (test *TestSuite) TestDelete() {
	test.writeConn.On("ModelContext", test.ctx, test.{{ .PrivateNameSingular }}).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*models.{{ .PublicNameSingular }})
		*arg = *test.{{ .PrivateNameSingular }}
		arg.IsDeleted = true
		arg.UpdatedAt = test.now.String()
	}).Return(test.query).Once()
	test.query.On("Column", "isdeleted").Return(test.query).Once()
	test.query.On("Returning", "id, createdby, createdat, updatedby, updatedat, guid, isdeleted").Return(test.query).Once()
	test.query.On("Where", "{{ .PrivateNameSingular }}.guid = ?", test.{{ .PrivateNameSingular }}.GUID).Return(test.query).Once()
	test.query.On("Where", "{{ .PrivateNameSingular }}.isdeleted = false").Return(test.query).Once()
	test.query.On("Update").Return(nil, nil).Once()

	err := test.repository.Delete(test.ctx, test.{{ .PrivateNameSingular }})

	test.Nil(err)
	test.Equal(test.now.String(), test.{{ .PrivateNameSingular }}.UpdatedAt)
}
`
