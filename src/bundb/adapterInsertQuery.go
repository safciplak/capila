package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterInsertQuery struct {
	Q *bun.InsertQuery
}

func adapterInsertQueryFunc(fn func(q InterfaceInsertQuery) InterfaceInsertQuery) func(q *bun.InsertQuery) *bun.InsertQuery {
	return func(q *bun.InsertQuery) *bun.InsertQuery {
		result := fn(adapterInsertQuery{q})

		return result.(adapterInsertQuery).Q
	}
}

func (a adapterInsertQuery) Conn(db bun.IConn) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Conn(db)}
}

func (a adapterInsertQuery) Model(model interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Model(model)}
}

func (a adapterInsertQuery) Apply(fn func(q InterfaceInsertQuery) InterfaceInsertQuery) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Apply(adapterInsertQueryFunc(fn))}
}

func (a adapterInsertQuery) With(name string, query schema.QueryAppender) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.With(name, query)}
}

func (a adapterInsertQuery) Table(tables ...string) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Table(tables...)}
}

func (a adapterInsertQuery) TableExpr(query string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterInsertQuery) ModelTableExpr(query string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterInsertQuery) Column(columns ...string) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Column(columns...)}
}

func (a adapterInsertQuery) ColumnExpr(query string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.ColumnExpr(query, args...)}
}

func (a adapterInsertQuery) ExcludeColumn(columns ...string) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.ExcludeColumn(columns...)}
}

func (a adapterInsertQuery) Value(column, expr string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Value(column, expr, args...)}
}

func (a adapterInsertQuery) Where(query string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Where(query, args...)}
}

func (a adapterInsertQuery) WhereOr(query string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.WhereOr(query, args...)}
}

func (a adapterInsertQuery) Returning(query string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Returning(query, args...)}
}

func (a adapterInsertQuery) Ignore() InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Ignore()}
}

func (a adapterInsertQuery) Replace() InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Replace()}
}

func (a adapterInsertQuery) OnConflict(s string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.On(s, args...)}
}

func (a adapterInsertQuery) Set(query string, args ...interface{}) InterfaceInsertQuery {
	return adapterInsertQuery{a.Q.Set(query, args...)}
}

func (a adapterInsertQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterInsertQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterInsertQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
