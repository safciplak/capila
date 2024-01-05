package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceTruncateTableQuery represents the bun.TruncateTableQuery struct
type InterfaceTruncateTableQuery interface {
	Conn(db bun.IConn) InterfaceTruncateTableQuery
	Model(model interface{}) InterfaceTruncateTableQuery
	Table(tables ...string) InterfaceTruncateTableQuery
	TableExpr(query string, args ...interface{}) InterfaceTruncateTableQuery
	ContinueIdentity() InterfaceTruncateTableQuery
	Restrict() InterfaceTruncateTableQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
