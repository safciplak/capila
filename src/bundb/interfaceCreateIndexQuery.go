package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceCreateIndexQuery represents the bun.CreateIndexQuery struct
type InterfaceCreateIndexQuery interface {
	Conn(db bun.IConn) InterfaceCreateIndexQuery
	Model(model interface{}) InterfaceCreateIndexQuery
	Unique() InterfaceCreateIndexQuery
	Concurrently() InterfaceCreateIndexQuery
	IfNotExists() InterfaceCreateIndexQuery
	Index(query string) InterfaceCreateIndexQuery
	IndexExpr(query string, args ...interface{}) InterfaceCreateIndexQuery
	Table(tables ...string) InterfaceCreateIndexQuery
	TableExpr(query string, args ...interface{}) InterfaceCreateIndexQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceCreateIndexQuery
	Using(query string, args ...interface{}) InterfaceCreateIndexQuery
	Column(columns ...string) InterfaceCreateIndexQuery
	ColumnExpr(query string, args ...interface{}) InterfaceCreateIndexQuery
	ExcludeColumn(columns ...string) InterfaceCreateIndexQuery
	Include(columns ...string) InterfaceCreateIndexQuery
	IncludeExpr(query string, args ...interface{}) InterfaceCreateIndexQuery
	Where(query string, args ...interface{}) InterfaceCreateIndexQuery
	WhereOr(query string, args ...interface{}) InterfaceCreateIndexQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
