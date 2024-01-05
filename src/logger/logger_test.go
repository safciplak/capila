package logger

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zapcore"
)

// TestSuite contains all the client test specific properties
type TestSuite struct {
	suite.Suite
	ctx context.Context
}

// SetupSuite initializes the environment in which the client will run
func (test *TestSuite) SetupSuite() {
	test.ctx = context.TODO()
	_ = os.Setenv("LOG_LEVEL", "")
}

// TearDownSuite removes all side effects after the suite has been completed
func (*TestSuite) TearDownSuite() {
	_ = os.Unsetenv("LOG_LEVEL")
}

// TestTestSuite runs the testsuite
func TestTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TestSuite))
}

func (test *TestSuite) TestNewLogger() {
	logger := NewLogger()

	test.NotNil(logger)
}

func (test *TestSuite) TestGetLevelDebug() {
	level := getLevel("debug")

	test.Equal(zapcore.DebugLevel, level)
}

func (test *TestSuite) TestGetLevelInfo() {
	level := getLevel("info")

	test.Equal(zapcore.InfoLevel, level)
}

func (test *TestSuite) TestGetLevelWarn() {
	level := getLevel("warn")

	test.Equal(zapcore.WarnLevel, level)
}

func (test *TestSuite) TestGetDefaultLogLevel() {
	level := getLevel("")

	test.Equal(DefaultLogLevel, level)
}

func (test *TestSuite) TestGetDefaultLogLevelWhenLogLevelIsWrong() {
	level := getLevel("custom")

	test.Equal(DefaultLogLevel, level)
}

func (test *TestSuite) TestGetZapLogger() {
	logger := NewLogger()

	test.NotNil(logger.GetZapLogger())
}

func (test *TestSuite) TestLog() {
	logger := NewLogger()

	test.NotNil(logger.Log(test.ctx))
}

func (test *TestSuite) TestGetLogger() {
	t := GetLogger()
	t2 := GetLogger()

	test.True(reflect.DeepEqual(t, t2))
}

func (test *TestSuite) TestGetLoggerNewInstance() {
	t := NewLogger()
	t2 := GetLogger()

	test.False(reflect.DeepEqual(t, t2))
}

func (test *TestSuite) TestNewLoggerInstances() {
	t := NewLogger()
	t2 := NewLogger()

	test.False(reflect.DeepEqual(t, t2))
}
