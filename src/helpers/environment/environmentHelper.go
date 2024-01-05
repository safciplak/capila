//go:generate generate-interfaces.sh

package helpers

import (
	"errors"
	"os"
	"strconv"
)

// EnvironmentHelper helps to get multiple environment variables.
type EnvironmentHelper struct {
	Err error
	Key string
}

// NewEnvironmentHelper instantiates the environment helper
func NewEnvironmentHelper() InterfaceEnvironmentHelper {
	return &EnvironmentHelper{}
}

// Get tries to get the value and fills the error field on fail.
func (environmentHelper *EnvironmentHelper) Get(environmentVariable string) string {
	value, err := environmentHelper.GetString(environmentVariable)

	// Allow multiple gets but always write away the (last) error and key for error handling
	if err != nil {
		environmentHelper.Key = environmentVariable
		environmentHelper.Err = err
	}

	return value
}

// Error exposes the error or nil.
func (environmentHelper *EnvironmentHelper) Error() error {
	return environmentHelper.Err
}

// GetString returns an environment value as a string ( if found ).
func (environmentHelper *EnvironmentHelper) GetString(key string) (string, error) {
	environmentVariable := os.Getenv(key)

	if environmentVariable == "" {
		return environmentVariable, errors.New("empty environmentVariable provided: " + key)
	}

	return environmentVariable, nil
}

// GetInteger returns an environment value as an integer ( if found ).
func (environmentHelper *EnvironmentHelper) GetInteger(key string) (int, error) {
	environmentVariable, err := environmentHelper.GetString(key)
	if err != nil {
		return 0, err
	}

	integer, convertError := strconv.Atoi(environmentVariable)
	if convertError != nil {
		return 0, convertError
	}

	return integer, nil
}

// GetBoolean returns an environment value as a boolean ( if found ).
func (environmentHelper *EnvironmentHelper) GetBoolean(key string) (bool, error) {
	environmentVariable, err := environmentHelper.GetString(key)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(environmentVariable)
}
