package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterDropTableQuery struct {
	Q *bun.DropTableQuery
}

func (a adapterDropTableQuery) Conn(db bun.IConn) InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Q.Conn(db)}
}

func (a adapterDropTableQuery) Model(model interface{}) InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Q.Model(model)}
}

func (a adapterDropTableQuery) Table(tables ...string) InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Q.Table(tables...)}
}

func (a adapterDropTableQuery) TableExpr(query string, args ...interface{}) InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterDropTableQuery) ModelTableExpr(query string, args ...interface{}) InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterDropTableQuery) IfExists() InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Q.IfExists()}
}

func (a adapterDropTableQuery) Restrict() InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Q.Restrict()}
}

func (a adapterDropTableQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterDropTableQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterDropTableQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
