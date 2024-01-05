package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestSuite contains all the client test specific properties
type TestSuite struct {
	suite.Suite
	ctx context.Context
}

// SetupSuite initializes the environment in which the client will run
func (test *TestSuite) SetupSuite() {
	test.ctx = context.Background()
}

// TestTestSuite runs the testsuite
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// TestNewContext tests the context creator
func (test *TestSuite) TestNewContext() {
	ctx := context.Background()

	newCtx := NewContext()

	test.Equal(ctx, newCtx)
}

// TestGetTwoLetterLanguageCodee tests getting and setting the language
func (test *TestSuite) TestGetTwoLetterLanguageCodee() {
	contextOne := context.Background()
	contextTwo := SetLanguage(context.Background(), "fr")

	test.Equal(GetTwoLetterLanguageCode(contextOne), "en")
	test.Equal(GetTwoLetterLanguageCode(contextTwo), "fr")
}

// TestGetTwoLetterLanguageCodee tests getting and setting the language
func (test *TestSuite) TestGetFourLetterLanguageCode() {
	contextOne := context.Background()
	contextTwo := SetLanguage(context.Background(), "fr-FR")

	test.Equal(GetFourLetterLanguageCode(contextOne), "en-US")
	test.Equal(GetFourLetterLanguageCode(contextTwo), "fr-FR")
}
