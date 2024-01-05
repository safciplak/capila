package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceInsertQuery represents the bun.InsertQuery struct
type InterfaceInsertQuery interface {
	Conn(db bun.IConn) InterfaceInsertQuery
	Model(model interface{}) InterfaceInsertQuery
	Apply(fn func(InterfaceInsertQuery) InterfaceInsertQuery) InterfaceInsertQuery
	With(name string, query schema.QueryAppender) InterfaceInsertQuery
	Table(tables ...string) InterfaceInsertQuery
	TableExpr(query string, args ...interface{}) InterfaceInsertQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceInsertQuery
	Column(columns ...string) InterfaceInsertQuery
	ColumnExpr(query string, args ...interface{}) InterfaceInsertQuery
	ExcludeColumn(columns ...string) InterfaceInsertQuery
	Value(column string, expr string, args ...interface{}) InterfaceInsertQuery
	Where(query string, args ...interface{}) InterfaceInsertQuery
	WhereOr(query string, args ...interface{}) InterfaceInsertQuery
	Returning(query string, args ...interface{}) InterfaceInsertQuery
	Ignore() InterfaceInsertQuery
	Replace() InterfaceInsertQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	OnConflict(s string, args ...interface{}) InterfaceInsertQuery
	Set(query string, args ...interface{}) InterfaceInsertQuery
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
