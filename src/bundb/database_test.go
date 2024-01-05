package bundb

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	environmentHelper "github.com/safciplak/capila/src/helpers/environment"
	"github.com/safciplak/capila/src/logger"
)

// TestSuite encapsulates all the tests
type TestSuite struct {
	suite.Suite
	ctx               context.Context
	environmentHelper *environmentHelper.MockInterfaceEnvironmentHelper
	logger            logger.InterfaceLogger
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	// Mocks used in the test
	test.environmentHelper = environmentHelper.NewMockInterfaceEnvironmentHelper(test.T())

	test.logger = logger.NewLogger()

	// Often used test objects
	test.ctx = context.TODO()

	_ = os.Setenv("DB_DATABASE", "mydatabase")
	_ = os.Setenv("DB_SCHEMA", "myschema")
	_ = os.Setenv("DB_READ_USER", "reader")
}

// TearDownTest tests whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	os.Clearenv()
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// TestTableNameInflector tests the tableNameInflector func
func (test *TestSuite) TestTableNameInflector() {
	tableName := tableNameInflector("testTable")

	test.Equal("myschema.testTable", tableName)
}

// TestGenerateConnections tests the GenerateConnections func
func (test *TestSuite) TestGenerateConnections() {
	_ = os.Setenv("APPLICATION_NAME", "development")
	_, err := GenerateConnections(test.ctx)

	// Since we don't actually want to connect to a DB in our tests and error should be given
	test.NotNil(err)
}

// TestGenerateConnections tests the GenerateConnections func
func (test *TestSuite) TestGenerateConnectionsEnvError() {
	_, err := GenerateConnections(test.ctx)

	test.Equal("empty environmentVariable provided: APPLICATION_NAME", err.Error())
}

// TestNewDatabase tests the NewDatabase func
func (test *TestSuite) TestNewDatabase() {
	_ = os.Setenv("APPLICATION_NAME", "development")
	db := NewDatabase(test.ctx, test.logger)

	test.Nil(db.Read)
	test.Nil(db.Write)
}
