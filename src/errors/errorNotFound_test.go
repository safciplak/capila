package errors //nolint:dupl // Every error should be tested in the same manner

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotFoundWrap(t *testing.T) {
	var (
		expectedErr = errors.New("Child error")
		result      = NewErrorNotFound().Wrap(expectedErr)
	)

	// Validate proper error interface compliance
	assert.ErrorIs(t, result, expectedErr)
	assert.ErrorIs(t, result, NewErrorNotFound())
	assert.Equal(t, expectedErr, result.Unwrap())

	// Validate that the return of Wrap is actually a specific error and not the base error
	assert.IsType(t, result, ErrorNotFound{})

	// Validate default values
	assert.Equal(t, ErrorCodeNotFound, result.GetCode())
	assert.Equal(t, 404, result.GetStatusCode())
	assert.Equal(t, "", result.GetDetail())
	assert.Equal(t, "Not Found", result.Error())
}
