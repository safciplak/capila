package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceDropIndexQuery represents the bun.DropIndexQuery struct
type InterfaceDropIndexQuery interface {
	Conn(db bun.IConn) InterfaceDropIndexQuery
	Model(model interface{}) InterfaceDropIndexQuery
	Concurrently() InterfaceDropIndexQuery
	IfExists() InterfaceDropIndexQuery
	Restrict() InterfaceDropIndexQuery
	Index(query string, args ...interface{}) InterfaceDropIndexQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
