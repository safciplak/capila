package errors //nolint:dupl // Every error should be tested in the same manner

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorBadRequestWrap(t *testing.T) {
	var (
		expectedErr = errors.New("Child error")
		result      = NewErrorBadRequest().Wrap(expectedErr)
	)

	// Validate proper error interface compliance
	assert.ErrorIs(t, result, expectedErr)
	assert.ErrorIs(t, result, NewErrorBadRequest())
	assert.Equal(t, expectedErr, result.Unwrap())

	// Validate that the return of Wrap is actually a specific error and not the base error
	assert.IsType(t, result, ErrorBadRequest{})

	// Validate default values
	assert.Equal(t, ErrorCodeBadRequest, result.GetCode())
	assert.Equal(t, 400, result.GetStatusCode())
	assert.Equal(t, "", result.GetDetail())
	assert.Equal(t, "Bad Request", result.Error())
}
