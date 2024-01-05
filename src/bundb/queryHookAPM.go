package bundb

import (
	"context"

	"github.com/uptrace/bun"

	"go.elastic.co/apm/module/apmsql/v2"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/stacktrace"
)

//nolint:gochecknoinits // this init is required...
func init() {
	stacktrace.RegisterLibraryPackage("github.com/uptrace/bun")
}

const elasticApmSpanKey = "go-apm-agent:span"

// QueryHook is an implementation of bun.queryHook that reports queries as spans to Elastic APM.
type QueryHook struct {
	opt *Options
}

type Options struct {
	User     string
	Database string
}

func NewQueryHook(opts *Options) *QueryHook {
	return &QueryHook{opt: opts}
}

// BeforeQuery initiates the span for the database query
func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	span, _ := apm.StartSpan(ctx, apmsql.QuerySignature(event.Query), "db.postgresql.query")
	span.Context.SetDatabase(apm.DatabaseSpanContext{
		Instance:  h.GetOptions().Database,
		Statement: event.Query,
		Type:      "sql",
		User:      h.GetOptions().User,
	})

	if event.Stash == nil {
		event.Stash = make(map[interface{}]interface{})
	}

	event.Stash[elasticApmSpanKey] = span

	return ctx
}

// AfterQuery ends the initiated span from BeforeQuery
func (h *QueryHook) AfterQuery(_ context.Context, event *bun.QueryEvent) {
	span, _ := event.Stash[elasticApmSpanKey]

	if s, spanOk := span.(*apm.Span); spanOk {
		s.End()
	}
}

// GetOptions returns option object
func (h *QueryHook) GetOptions() *Options {
	if h.opt != nil {
		return h.opt
	}
	return &Options{}
}
