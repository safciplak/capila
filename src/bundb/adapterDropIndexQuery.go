package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterDropIndexQuery struct {
	Q *bun.DropIndexQuery
}

func (a adapterDropIndexQuery) Conn(db bun.IConn) InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Q.Conn(db)}
}

func (a adapterDropIndexQuery) Model(model interface{}) InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Q.Model(model)}
}

func (a adapterDropIndexQuery) Concurrently() InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Q.Concurrently()}
}

func (a adapterDropIndexQuery) IfExists() InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Q.IfExists()}
}

func (a adapterDropIndexQuery) Restrict() InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Q.Restrict()}
}

func (a adapterDropIndexQuery) Index(query string, args ...interface{}) InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Q.Index(query, args...)}
}

func (a adapterDropIndexQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterDropIndexQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterDropIndexQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
