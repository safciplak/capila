package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
	"go.elastic.co/apm/module/apmsql/v2"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/stacktrace"
)

//nolint:gochecknoinits // this init is required...
func init() {
	stacktrace.RegisterLibraryPackage("github.com/go-pg/pg")
}

const elasticApmSpanKey = "go-apm-agent:span"

// queryHook is an implementation of pg.queryHook that reports queries as spans to Elastic APM.
type queryHook struct{}

// BeforeQuery initiates the span for the database query
func (qh queryHook) BeforeQuery(ctx context.Context, evt *pg.QueryEvent) (context.Context, error) {
	var (
		database string
		dbUser   string
	)

	if db, ok := evt.DB.(*pg.DB); ok {
		opts := db.Options()
		dbUser = opts.User
		database = opts.Database
	}

	query, err := evt.UnformattedQuery()
	sql := string(query)

	if err != nil {
		// Expose the error making it a bit easier to debug
		sql = fmt.Sprintf("[go-pg] error: %s", err.Error())
	}

	span, _ := apm.StartSpan(ctx, apmsql.QuerySignature(sql), "db.postgresql.query")
	span.Context.SetDatabase(apm.DatabaseSpanContext{
		Statement: sql,

		// Static
		Type:     "sql",
		User:     dbUser,
		Instance: database,
	})

	if evt.Stash == nil {
		evt.Stash = make(map[interface{}]interface{})
	}

	evt.Stash[elasticApmSpanKey] = span

	return ctx, err
}

// AfterQuery ends the initiated span from BeforeQuery
func (qh queryHook) AfterQuery(_ context.Context, evt *pg.QueryEvent) error {
	span, ok := evt.Stash[elasticApmSpanKey]
	if !ok {
		return errors.New("unable to retrieve APM span after a query")
	}

	if s, spanOk := span.(*apm.Span); spanOk {
		s.End()
	}

	return nil
}
