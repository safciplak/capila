package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateToken(t *testing.T) {
	result := GenerateToken(1, "jwtSigningKey", 1)

	assert.Greater(t, len(result), 0)
}

func Test_GenerateToken_Zero(t *testing.T) {
	result := GenerateToken(0, "jwtSigningKey", 1)

	assert.Greater(t, len(result), 0)
}

func Test_GenerateToken_Wrong(t *testing.T) {
	var err error

	result := GenerateToken(1, "jwtSigningKey", 1)

	if result == "1" {
		err = fmt.Errorf("token was not generated")
	}

	assert.Nil(t, err)
}
