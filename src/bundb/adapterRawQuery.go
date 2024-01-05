package bundb

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// AdapterRawQuery is the adapter
type adapterRawQuery struct {
	Q *bun.RawQuery
}

func (a adapterRawQuery) Conn(db bun.IConn) InterfaceRawQuery {
	return adapterRawQuery{a.Q.Conn(db)}
}

func (a adapterRawQuery) Scan(ctx context.Context, dest ...interface{}) error {
	return a.Q.Scan(ctx, dest...)
}

func (a adapterRawQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterRawQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}
