package bundb

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHandleQueryResult tests if the error is returned correctly
func TestHandleQueryResult(t *testing.T) {
	expectedErr := errors.New("expected error")

	err := HandleQueryResult(nil, expectedErr)

	assert.Equal(t, expectedErr, err)
}
