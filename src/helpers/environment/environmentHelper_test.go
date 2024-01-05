package helpers

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Test Suite which encapsulate the tests for the handler!
type TestSuite struct {
	suite.Suite
	ctx context.Context

	environmentHelper InterfaceEnvironmentHelper
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	test.ctx = context.TODO()
	test.environmentHelper = NewEnvironmentHelper()

	_ = os.Setenv("MY_TEST_STRING", "Japie Schakel")
	_ = os.Setenv("MY_TEST_STRING_TWO", "Lonnie de Vrij")
	_ = os.Setenv("MY_TEST_BOOLEAN", "true")
	_ = os.Setenv("MY_TEST_NON_BOOLEAN", "ghello")
	_ = os.Setenv("MY_TEST_INTEGER", "2")
	_ = os.Setenv("MY_TEST_NON_INTEGER", "999999999999999999999999999999999999999999999999999999999999999999999999999999999")
}

// TearDownTest asserts whether the mock has been handled correctly after each test
func (test *TestSuite) TearDownTest() {
	os.Clearenv()
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (test *TestSuite) TestEnvironmentHelper() {
	myTestString := test.environmentHelper.Get("I_DONT_EXIST_NOOOO")
	myTestStringTwo := test.environmentHelper.Get("I_DONT_EXIST")

	test.NotNil(test.environmentHelper.Error())
	test.Equal("", myTestString)
	test.Equal("", myTestStringTwo)
}

func (test *TestSuite) TestEnvironmentHelperWithCorrectValues() {
	myTestString := test.environmentHelper.Get("MY_TEST_STRING")
	myTestStringTwo := test.environmentHelper.Get("MY_TEST_STRING_TWO")

	test.Nil(test.environmentHelper.Error())
	test.Equal("Japie Schakel", myTestString)
	test.Equal("Lonnie de Vrij", myTestStringTwo)
}

func (test *TestSuite) TestGetEnvironmentString() {
	nonExistentVariable, firstErr := test.environmentHelper.GetString("MY_TEST_VARIABLE_THAT_DOES_NOT_EXIST")

	test.Equal("", nonExistentVariable)
	test.NotNil(firstErr)

	existingVariable, err := test.environmentHelper.GetString("MY_TEST_STRING")
	test.Equal("Japie Schakel", existingVariable)
	test.Nil(err)
}

func (test *TestSuite) TestGetEnvironmentInteger() {
	nonExistentVariable, firstErr := test.environmentHelper.GetInteger("MY_TEST_INTEGERZZZZ")

	test.Equal(0, nonExistentVariable)
	test.NotNil(firstErr)

	existingButNonIntegerVariable, secondErr := test.environmentHelper.GetInteger("MY_TEST_NON_INTEGER")
	test.Equal(0, existingButNonIntegerVariable)
	test.NotNil(secondErr)

	existingVariable, err := test.environmentHelper.GetInteger("MY_TEST_INTEGER")
	test.Equal(2, existingVariable)
	test.Nil(err)
}

func (test *TestSuite) Test_GetEnvironmentBoolean() {
	nonExistentVariable, firstErr := test.environmentHelper.GetBoolean("MY_TEST_BOOLEANZZZ")

	test.Equal(false, nonExistentVariable)
	test.NotNil(firstErr)

	existingButNonBooleanVariable, secondErr := test.environmentHelper.GetBoolean("MY_TEST_NON_BOOLEAN")
	test.Equal(false, existingButNonBooleanVariable)
	test.NotNil(secondErr)

	existingVariable, err := test.environmentHelper.GetBoolean("MY_TEST_BOOLEAN")
	test.Equal(true, existingVariable)
	test.Nil(err)
}
