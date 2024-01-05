package apm

import (
	"context"

	"go.elastic.co/apm/v2"
)

// TraceError sends the error to APM.
// This needs to be called before the error is returned.
func TraceError(ctx context.Context, err error) error {
	apm.CaptureError(ctx, err).Send()

	return err
}

// Start starts a span without parent and returns the new span.
// Returns nil if the span could not be started.
func Start(ctx context.Context, name, spanType string) *apm.Span {
	span, _ := apm.StartSpan(ctx, name, spanType)

	return span
}

// End ends the given span.
// Will check if the given span is not a nil object.
func End(span *apm.Span) {
	if span != nil {
		span.End()
	}
}
