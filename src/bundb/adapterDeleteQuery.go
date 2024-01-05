package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterDeleteQuery struct {
	Q *bun.DeleteQuery
}

func adapterDeleteQueryFunc(fn func(q InterfaceDeleteQuery) InterfaceDeleteQuery) func(q *bun.DeleteQuery) *bun.DeleteQuery {
	return func(q *bun.DeleteQuery) *bun.DeleteQuery {
		result := fn(adapterDeleteQuery{q})
		return result.(adapterDeleteQuery).Q
	}
}

func (a adapterDeleteQuery) Conn(db bun.IConn) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.Conn(db)}
}

func (a adapterDeleteQuery) Model(model interface{}) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.Model(model)}
}

func (a adapterDeleteQuery) Apply(fn func(q InterfaceDeleteQuery) InterfaceDeleteQuery) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.Apply(adapterDeleteQueryFunc(fn))}
}

func (a adapterDeleteQuery) With(name string, query schema.QueryAppender) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.With(name, query)}
}

func (a adapterDeleteQuery) Table(tables ...string) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.Table(tables...)}
}

func (a adapterDeleteQuery) TableExpr(query string, args ...interface{}) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterDeleteQuery) ModelTableExpr(query string, args ...interface{}) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterDeleteQuery) WherePK(cols ...string) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.WherePK(cols...)}
}

func (a adapterDeleteQuery) Where(query string, args ...interface{}) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.Where(query, args...)}
}

func (a adapterDeleteQuery) WhereOr(query string, args ...interface{}) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.WhereOr(query, args...)}
}

func (a adapterDeleteQuery) WhereGroup(sep string, fn func(InterfaceDeleteQuery) InterfaceDeleteQuery) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.WhereGroup(sep, adapterDeleteQueryFunc(fn))}
}

func (a adapterDeleteQuery) WhereDeleted() InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.WhereDeleted()}
}

func (a adapterDeleteQuery) WhereAllWithDeleted() InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.WhereAllWithDeleted()}
}

func (a adapterDeleteQuery) ForceDelete() InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.ForceDelete()}
}

func (a adapterDeleteQuery) Returning(query string, args ...interface{}) InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Q.Returning(query, args...)}
}

func (a adapterDeleteQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterDeleteQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterDeleteQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
