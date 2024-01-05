package bundb

import (
	"context"
	"database/sql"
)

// InterfaceConn represents the bun.Conn struct
type InterfaceConn interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	NewValues(model interface{}) InterfaceValuesQuery
	NewSelect() InterfaceSelectQuery
	NewInsert() InterfaceInsertQuery
	NewUpdate() InterfaceUpdateQuery
	NewDelete() InterfaceDeleteQuery
	NewCreateTable() InterfaceCreateTableQuery
	NewDropTable() InterfaceDropTableQuery
	NewCreateIndex() InterfaceCreateIndexQuery
	NewDropIndex() InterfaceDropIndexQuery
	NewTruncateTable() InterfaceTruncateTableQuery
	NewAddColumn() InterfaceAddColumnQuery
	NewDropColumn() InterfaceDropColumnQuery
}
