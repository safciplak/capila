package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceDeleteQuery represents the bun.DeleteQuery struct
type InterfaceDeleteQuery interface {
	Conn(db bun.IConn) InterfaceDeleteQuery
	Model(model interface{}) InterfaceDeleteQuery
	Apply(fn func(InterfaceDeleteQuery) InterfaceDeleteQuery) InterfaceDeleteQuery
	With(name string, query schema.QueryAppender) InterfaceDeleteQuery
	Table(tables ...string) InterfaceDeleteQuery
	TableExpr(query string, args ...interface{}) InterfaceDeleteQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceDeleteQuery
	WherePK(cols ...string) InterfaceDeleteQuery
	Where(query string, args ...interface{}) InterfaceDeleteQuery
	WhereOr(query string, args ...interface{}) InterfaceDeleteQuery
	WhereGroup(sep string, fn func(InterfaceDeleteQuery) InterfaceDeleteQuery) InterfaceDeleteQuery
	WhereDeleted() InterfaceDeleteQuery
	WhereAllWithDeleted() InterfaceDeleteQuery
	ForceDelete() InterfaceDeleteQuery
	Returning(query string, args ...interface{}) InterfaceDeleteQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
