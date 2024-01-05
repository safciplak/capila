package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterDropColumnQuery struct {
	Q *bun.DropColumnQuery
}

func (a adapterDropColumnQuery) Conn(db bun.IConn) InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Q.Conn(db)}
}

func (a adapterDropColumnQuery) Model(model interface{}) InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Q.Model(model)}
}

func (a adapterDropColumnQuery) Table(tables ...string) InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Q.Table(tables...)}
}

func (a adapterDropColumnQuery) TableExpr(query string, args ...interface{}) InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterDropColumnQuery) ModelTableExpr(query string, args ...interface{}) InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterDropColumnQuery) Column(columns ...string) InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Q.Column(columns...)}
}

func (a adapterDropColumnQuery) ColumnExpr(query string, args ...interface{}) InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Q.ColumnExpr(query, args...)}
}

func (a adapterDropColumnQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterDropColumnQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterDropColumnQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
