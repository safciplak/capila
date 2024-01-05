package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceCreateTableQuery represents the bun.CreateTableQuery struct
type InterfaceCreateTableQuery interface {
	Conn(db bun.IConn) InterfaceCreateTableQuery
	Model(model interface{}) InterfaceCreateTableQuery
	Table(tables ...string) InterfaceCreateTableQuery
	TableExpr(query string, args ...interface{}) InterfaceCreateTableQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceCreateTableQuery
	ColumnExpr(query string, args ...interface{}) InterfaceCreateTableQuery
	Temp() InterfaceCreateTableQuery
	IfNotExists() InterfaceCreateTableQuery
	Varchar(n int) InterfaceCreateTableQuery
	ForeignKey(query string, args ...interface{}) InterfaceCreateTableQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
