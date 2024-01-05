package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceDropColumnQuery represents the bun.DropColumnQuery struct
type InterfaceDropColumnQuery interface {
	Conn(db bun.IConn) InterfaceDropColumnQuery
	Model(model interface{}) InterfaceDropColumnQuery
	Table(tables ...string) InterfaceDropColumnQuery
	TableExpr(query string, args ...interface{}) InterfaceDropColumnQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceDropColumnQuery
	Column(columns ...string) InterfaceDropColumnQuery
	ColumnExpr(query string, args ...interface{}) InterfaceDropColumnQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
