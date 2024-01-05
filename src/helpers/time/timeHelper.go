//go:generate generate-interfaces.sh

package helpers

import (
	"time"
)

type TimeHelper time.Time

func NewTimeHelper() InterfaceTimeHelper {
	return &TimeHelper{}
}

// Now returns the current time.
func (timeHelper *TimeHelper) Now() time.Time {
	return time.Now()
}

// NowIn returns the current time at the given location string.
func (timeHelper *TimeHelper) NowIn(locationName, fallbackLocationName string) time.Time {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		loc, err = time.LoadLocation(fallbackLocationName)
	}

	if err != nil {
		return time.Now()
	}

	return time.Now().In(loc)
}

// NewDate is a shortcut that returns a new UTC date.
func (timeHelper *TimeHelper) NewDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
