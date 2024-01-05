package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceDropTableQuery represents the bun.DropTableQuery struct
type InterfaceDropTableQuery interface {
	Conn(db bun.IConn) InterfaceDropTableQuery
	Model(model interface{}) InterfaceDropTableQuery
	Table(tables ...string) InterfaceDropTableQuery
	TableExpr(query string, args ...interface{}) InterfaceDropTableQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceDropTableQuery
	IfExists() InterfaceDropTableQuery
	Restrict() InterfaceDropTableQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
