package file

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestFileSuite struct {
	suite.Suite

	baseDir         string
	sourceFile      string
	destinationFile string
}

func (test *TestFileSuite) SetupTest() {
	test.baseDir = GetRootDir()
	test.sourceFile = fmt.Sprintf("%s/README.md", test.baseDir)
	test.destinationFile = fmt.Sprintf("%s/README.copy", test.baseDir)
}

func (test *TestFileSuite) TearDownTest() {
	if Exists(test.destinationFile) {
		test.Nil(os.Remove(test.destinationFile))
	}
}

func TestFileTestSuite(t *testing.T) {
	suite.Run(t, new(TestFileSuite))
}

func (test *TestFileSuite) TestCopy() {
	test.Nil(Copy(test.sourceFile, test.destinationFile))
}

func (test *TestFileSuite) TestCopyWrongInputFile() {
	test.sourceFile = fmt.Sprintf("%s/IDONOTEXIST", test.baseDir)

	test.NotNil(Copy(test.sourceFile, test.destinationFile))
}

func (test *TestFileSuite) TestCopyWrongOutputFile() {
	test.destinationFile = fmt.Sprintf("%s/bla/IDONOTEXIST", test.baseDir)

	test.NotNil(Copy(test.sourceFile, test.destinationFile))
}

func (test *TestFileSuite) TestRootDir() {
	rootDir := GetRootDir()
	// Todo: figure out how to test findParentDir
	test.Equal(test.baseDir, rootDir)
}
