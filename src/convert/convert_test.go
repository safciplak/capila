package convert

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

// TestTestSuite runs the testsuite
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (test *TestSuite) TestNew() {
	var (
		testString  = "test"
		testUint    = uint(1337)
		testBool    = true
		testFloat32 = float32(13.37)
		testFloat64 = 13.37
	)

	test.Equal(&testString, NewString(testString))
	test.Equal(&testUint, NewUint(testUint))
	test.Equal(&testBool, NewBool(testBool))
	test.Equal(&testFloat32, NewFloat32(testFloat32))
	test.Equal(&testFloat64, NewFloat64(testFloat64))
}

func (test *TestSuite) TestNewBool() {
	var (
		testBoolTrue  = true
		testBoolFalse = false
	)

	test.Equal(&testBoolTrue, NewBool(true))
	test.Equal(&testBoolFalse, NewBool(false))
}

func (test *TestSuite) TestNewFloat32() {
	var (
		testFloat32 float32 = 123.10
	)

	test.Equal(&testFloat32, NewFloat32(123.10))
}

func (test *TestSuite) TestNewInteger() {
	var (
		testInteger = 1234
	)

	test.Equal(&testInteger, NewInteger(1234))
}

func (test *TestSuite) TestNewInteger32() {
	var (
		testInteger = int32(1234)
	)

	test.Equal(&testInteger, NewInteger32(1234))
}

func (test *TestSuite) TestNewInteger64() {
	var (
		testInteger = int64(1234)
	)

	test.Equal(&testInteger, NewInteger64(1234))
}

func (test *TestSuite) TestStringToBool() {
	test.Equal(false, StringToBool("test"))
	test.Equal(true, StringToBool("True"))
	test.Equal(true, StringToBool("true"))
	test.Equal(true, StringToBool("1"))
	test.Equal(false, StringToBool("2"))
}

func (test *TestSuite) TestPointerToString() {
	var (
		testString  = "test"
		testPointer = NewString("test")
	)

	test.Equal(testString, PointerToString(testPointer))
}

func (test *TestSuite) TestIntPointerToStringPointer() {
	var (
		testString     = "88"
		testIntPointer = 88
	)

	test.Equal(&testString, IntPointerToStringPointer(&testIntPointer))
}
