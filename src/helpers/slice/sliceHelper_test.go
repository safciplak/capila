package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineOnlyMatchingValues(t *testing.T) {
	var (
		test       = assert.New(t)
		sliceOne   = []string{"one", "two", "three"}
		sliceTwo   = []string{"one", "three"}
		sliceThree = []string{"two", "three"}
	)

	test.Equal(sliceTwo, CombineOnlyMatchingValues(sliceOne, sliceTwo))
	test.Equal([]string{"three"}, CombineOnlyMatchingValues(sliceTwo, sliceThree))
}
