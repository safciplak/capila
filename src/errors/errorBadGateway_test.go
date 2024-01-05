package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorBadGatewayWrap(t *testing.T) {
	var (
		expectedErr = errors.New("Child error")
		result      = NewErrorBadGateway().Wrap(expectedErr)
	)

	// Validate proper error interface compliance
	assert.ErrorIs(t, result, expectedErr)
	assert.ErrorIs(t, result, NewErrorBadGateway())
	assert.Equal(t, expectedErr, result.Unwrap())

	// Validate that the return of Wrap is actually a specific error and not the base error
	assert.IsType(t, ErrorBadGateway{}, result)

	// Validate default values
	assert.Equal(t, ErrorCodeBadGateway, result.GetCode())
	assert.Equal(t, 502, result.GetStatusCode())
	assert.Equal(t, "", result.GetDetail())
	assert.Equal(t, "Bad Gateway", result.Error())
}
