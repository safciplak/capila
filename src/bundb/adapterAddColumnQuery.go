package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterAddColumnQuery struct {
	Q *bun.AddColumnQuery
}

func (a adapterAddColumnQuery) Conn(db bun.IConn) InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Q.Conn(db)}
}

func (a adapterAddColumnQuery) Model(model interface{}) InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Q.Model(model)}
}

func (a adapterAddColumnQuery) Table(tables ...string) InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Q.Table(tables...)}
}

func (a adapterAddColumnQuery) TableExpr(query string, args ...interface{}) InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterAddColumnQuery) ModelTableExpr(query string, args ...interface{}) InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterAddColumnQuery) ColumnExpr(query string, args ...interface{}) InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Q.ColumnExpr(query, args...)}
}

func (a adapterAddColumnQuery) IfNotExists() InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Q.IfNotExists()}
}

func (a adapterAddColumnQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterAddColumnQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterAddColumnQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
