package apm_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.elastic.co/apm/v2"

	capilaAPM "github.com/safciplak/capila/src/apm"
)

// TestTraceError checks to see if the traced error is the same as the given error.
func TestTraceError(t *testing.T) {
	// Given we have a new error
	err := errors.New("this is an error produced during testing")

	// It should log the error with context
	tracedErr := capilaAPM.TraceError(context.Background(), err)

	// The error should be the same as before
	assert.Equal(t, err, tracedErr)
}

// TestTraceErrorWithContext tests error logging with a context but without transaction.
func TestTraceErrorWithExistingContext(t *testing.T) {
	var (
		ctx = context.Background()
	)

	// Given we have a new error
	err := errors.New("this is an error produced during testing")

	// It should log the error
	tracedErr := capilaAPM.TraceError(ctx, err)

	// The error should be the same as before
	assert.Equal(t, err, tracedErr)
}

// TestTraceErrorWithoutError error tests error logging with a context but without an error.
func TestTraceErrorWithoutError(t *testing.T) {
	var (
		apmTransaction = new(apm.Transaction)
		ctx            = apm.ContextWithTransaction(context.Background(), apmTransaction)
	)

	// ContextWithTransaction
	capilaAPM.Start(ctx, "test", "testing")

	// It should log the error
	tracedErr := capilaAPM.TraceError(ctx, nil)

	// The error should be the same as before
	assert.Nil(t, tracedErr)
}

// TestEnd tests the ending of a span
func TestEnd(t *testing.T) {
	span := capilaAPM.Start(context.Background(), "test", "testing")

	capilaAPM.End(span)

	assert.True(t, span.Dropped())
}
