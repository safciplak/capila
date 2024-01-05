package helpers

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	timeHelper InterfaceTimeHelper
}

var (
	// Keep in mind, timezone locations are case sensitive on Linux
	locationAMS          = "Europe/Amsterdam"     //nolint:gochecknoglobals // just a test file
	locationNoneExistent = "some-non/existent-tz" //nolint:gochecknoglobals // just a test file
	testTZ               = "America/New_York"     //nolint:gochecknoglobals // just a test file
)

// TestTestSuite runs the testsuite
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	err := os.Setenv("TZ", testTZ)
	test.Nil(err, "Error setting the TZ env")

	test.timeHelper = NewTimeHelper()
}

func (test *TestSuite) TestTimeHelper() {
	test.NotNil(test.timeHelper)
	test.NotNil(test.timeHelper.Now())

	test.NotNil(test.timeHelper)
	test.NotNil(test.timeHelper.NewDate(2000, time.January, 1))
}

func (test *TestSuite) TestNewDate() {
	date := test.timeHelper.NewDate(2020, 5, 1)

	test.Equal(2020, date.Year())
	test.Equal(time.Month(5), date.Month())
	test.Equal(1, date.Day())
}

func (test *TestSuite) TestNowIn() {
	timeInAmsterdam := test.timeHelper.NowIn(locationAMS, locationNoneExistent)
	println(timeInAmsterdam.Location().String())
	test.Equal(locationAMS, timeInAmsterdam.Location().String())
}

func (test *TestSuite) TestNowInFallback() {
	timeInAmsterdam := test.timeHelper.NowIn(locationNoneExistent, locationAMS)
	test.Equal(locationAMS, timeInAmsterdam.Location().String())
}

func (test *TestSuite) TestNowInUnknownLocation() {
	// Both location and fallbackLocation are invalid. Use system tz
	timeInNY := test.timeHelper.NowIn(locationNoneExistent, locationNoneExistent)
	test.Equal(testTZ, timeInNY.Location().String())
}
