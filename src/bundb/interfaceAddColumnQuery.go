package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceAddColumnQuery represents the bun.AddColumnQuery struct
type InterfaceAddColumnQuery interface {
	Conn(db bun.IConn) InterfaceAddColumnQuery
	Model(model interface{}) InterfaceAddColumnQuery
	Table(tables ...string) InterfaceAddColumnQuery
	TableExpr(query string, args ...interface{}) InterfaceAddColumnQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceAddColumnQuery
	ColumnExpr(query string, args ...interface{}) InterfaceAddColumnQuery
	IfNotExists() InterfaceAddColumnQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Exec(ctx context.Context, dest ...interface{}) (sql.Result, error)
}
