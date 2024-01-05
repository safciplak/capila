package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterCreateIndexQuery struct {
	Q *bun.CreateIndexQuery
}

func (a adapterCreateIndexQuery) Conn(db bun.IConn) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Conn(db)}
}

func (a adapterCreateIndexQuery) Model(model interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Model(model)}
}

func (a adapterCreateIndexQuery) Unique() InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Unique()}
}

func (a adapterCreateIndexQuery) Concurrently() InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Concurrently()}
}

func (a adapterCreateIndexQuery) IfNotExists() InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.IfNotExists()}
}

func (a adapterCreateIndexQuery) Index(query string) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Index(query)}
}

func (a adapterCreateIndexQuery) IndexExpr(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.IndexExpr(query, args...)}
}

func (a adapterCreateIndexQuery) Table(tables ...string) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Table(tables...)}
}

func (a adapterCreateIndexQuery) TableExpr(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterCreateIndexQuery) ModelTableExpr(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterCreateIndexQuery) Using(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Using(query, args...)}
}

func (a adapterCreateIndexQuery) Column(columns ...string) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Column(columns...)}
}

func (a adapterCreateIndexQuery) ColumnExpr(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.ColumnExpr(query, args...)}
}

func (a adapterCreateIndexQuery) ExcludeColumn(columns ...string) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.ExcludeColumn(columns...)}
}

func (a adapterCreateIndexQuery) Include(columns ...string) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Include(columns...)}
}

func (a adapterCreateIndexQuery) IncludeExpr(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.IncludeExpr(query, args...)}
}

func (a adapterCreateIndexQuery) Where(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.Where(query, args...)}
}

func (a adapterCreateIndexQuery) WhereOr(query string, args ...interface{}) InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Q.WhereOr(query, args...)}
}

func (a adapterCreateIndexQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterCreateIndexQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterCreateIndexQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
