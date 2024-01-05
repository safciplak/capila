package mock

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/safciplak/capila/src/file"
)

// Test Suite which encapsulate the tests for the test.
type TestSuite struct {
	suite.Suite
	rootDir  string
	mockPath string
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	test.rootDir = file.GetRootDir()
	test.mockPath = test.getMockFile("hotel.json")
}

func (test *TestSuite) getMockFile(targetFile string) string {
	return filepath.Join(test.rootDir, "src", "helpers", "mock", "mocks", targetFile)
}

// TestClientTestSuite Runs the testsuite
func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// TestMockFullPath cannot test the exact value because this differs on every system.
func (test *TestSuite) TestMockFullPath() {
	test.Contains(FullPath(test.rootDir, "v1", "hotelService/hotels", "200.json"), "/mock/v1/hotelService/hotels/200.json")
}

func (test *TestSuite) TestStructFromMockJSON() {
	response := hotel{}
	err := StructFromMockJSON(test.mockPath, &response)
	test.Nil(err)
}

func (test *TestSuite) TestSliceStructFromMockJSON() {
	var response []hotel
	err := StructFromMockJSON(test.getMockFile("hotels.json"), &response)

	test.Nil(err)
	test.Len(response, 2)
}

func (test *TestSuite) TestStructFromMockJSON_NoPointer() {
	response := hotel{}
	err := StructFromMockJSON(test.mockPath, response)
	test.EqualError(err, "targetStruct should be a pointer")
}

func (test *TestSuite) TestStructFromMockJSONWrongPath() {
	response := hotel{}
	err := StructFromMockJSON(test.getMockFile("empty.json"), &response)

	test.NotNil(err)
	test.Contains(err.Error(), "no such file or directory")
}

func (test *TestSuite) TestStructFromMockJSONInvalidStruct() {
	invalidMarshalStruct := &struct {
		Code int
	}{}
	err := StructFromMockJSON(test.mockPath, invalidMarshalStruct)

	test.NotNil(err)
	test.Contains(err.Error(), "json: cannot unmarshal")
}
