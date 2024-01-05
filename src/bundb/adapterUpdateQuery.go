package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterUpdateQuery struct {
	Q *bun.UpdateQuery
}

func adapterUpdateQueryFunc(fn func(q InterfaceUpdateQuery) InterfaceUpdateQuery) func(q *bun.UpdateQuery) *bun.UpdateQuery {
	return func(q *bun.UpdateQuery) *bun.UpdateQuery {
		result := fn(adapterUpdateQuery{q})

		return result.(adapterUpdateQuery).Q
	}
}

func (a adapterUpdateQuery) Conn(db bun.IConn) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Conn(db)}
}

func (a adapterUpdateQuery) Model(model interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Model(model)}
}

func (a adapterUpdateQuery) Apply(fn func(q InterfaceUpdateQuery) InterfaceUpdateQuery) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Apply(adapterUpdateQueryFunc(fn))}
}

func (a adapterUpdateQuery) With(name string, query schema.QueryAppender) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.With(name, query)}
}

func (a adapterUpdateQuery) Table(tables ...string) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Table(tables...)}
}

func (a adapterUpdateQuery) TableExpr(query string, args ...interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterUpdateQuery) ModelTableExpr(query string, args ...interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterUpdateQuery) Column(columns ...string) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Column(columns...)}
}

func (a adapterUpdateQuery) ExcludeColumn(columns ...string) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.ExcludeColumn(columns...)}
}

func (a adapterUpdateQuery) Set(query string, args ...interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Set(query, args...)}
}

func (a adapterUpdateQuery) Value(column, expr string, args ...interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Value(column, expr, args...)}
}

func (a adapterUpdateQuery) OmitZero() InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.OmitZero()}
}

func (a adapterUpdateQuery) WherePK(cols ...string) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.WherePK(cols...)}
}

func (a adapterUpdateQuery) Where(query string, args ...interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Where(query, args...)}
}

func (a adapterUpdateQuery) WhereOr(query string, args ...interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.WhereOr(query, args...)}
}

func (a adapterUpdateQuery) WhereGroup(sep string, fn func(InterfaceUpdateQuery) InterfaceUpdateQuery) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.WhereGroup(sep, adapterUpdateQueryFunc(fn))}
}

func (a adapterUpdateQuery) WhereDeleted() InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.WhereDeleted()}
}

func (a adapterUpdateQuery) WhereAllWithDeleted() InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.WhereAllWithDeleted()}
}

func (a adapterUpdateQuery) Returning(query string, args ...interface{}) InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Returning(query, args...)}
}

func (a adapterUpdateQuery) Bulk() InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Q.Bulk()}
}

func (a adapterUpdateQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterUpdateQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterUpdateQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}

func (a adapterUpdateQuery) FQN(name string) bun.Ident {
	return a.Q.FQN(name)
}
