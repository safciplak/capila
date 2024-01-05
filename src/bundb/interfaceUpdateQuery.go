package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceUpdateQuery represents the bun.UpdateQuery struct
type InterfaceUpdateQuery interface {
	Conn(db bun.IConn) InterfaceUpdateQuery
	Model(model interface{}) InterfaceUpdateQuery
	Apply(fn func(InterfaceUpdateQuery) InterfaceUpdateQuery) InterfaceUpdateQuery
	With(name string, query schema.QueryAppender) InterfaceUpdateQuery
	Table(tables ...string) InterfaceUpdateQuery
	TableExpr(query string, args ...interface{}) InterfaceUpdateQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceUpdateQuery
	Column(columns ...string) InterfaceUpdateQuery
	ExcludeColumn(columns ...string) InterfaceUpdateQuery
	Set(query string, args ...interface{}) InterfaceUpdateQuery
	Value(column string, expr string, args ...interface{}) InterfaceUpdateQuery
	OmitZero() InterfaceUpdateQuery
	WherePK(cols ...string) InterfaceUpdateQuery
	Where(query string, args ...interface{}) InterfaceUpdateQuery
	WhereOr(query string, args ...interface{}) InterfaceUpdateQuery
	WhereGroup(sep string, fn func(InterfaceUpdateQuery) InterfaceUpdateQuery) InterfaceUpdateQuery
	WhereDeleted() InterfaceUpdateQuery
	WhereAllWithDeleted() InterfaceUpdateQuery
	Returning(query string, args ...interface{}) InterfaceUpdateQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Bulk() InterfaceUpdateQuery
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
	FQN(name string) bun.Ident
}
