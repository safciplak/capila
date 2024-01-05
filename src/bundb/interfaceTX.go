package bundb

import (
	"context"
	"database/sql"
)

// InterfaceTx represents the bun.Tx struct
type InterfaceTx interface {
	Commit() error
	Rollback() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
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
