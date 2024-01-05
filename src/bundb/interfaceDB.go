package bundb

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceDB represents the bun.DB struct
type InterfaceDB interface {
	String() string
	DBStats() bun.DBStats
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
	NewRaw(query string, args ...interface{}) InterfaceRawQuery
	ResetModel(ctx context.Context, models ...interface{}) error
	Dialect() schema.Dialect
	ScanRows(ctx context.Context, rows *sql.Rows, dest ...interface{}) error
	ScanRow(ctx context.Context, rows *sql.Rows, dest ...interface{}) error
	AddQueryHook(hook bun.QueryHook)
	Table(typ reflect.Type) *schema.Table
	RegisterModel(models ...interface{})
	WithNamedArg(name string, value interface{}) InterfaceDB
	Formatter() schema.Formatter
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Prepare(query string) (bun.Stmt, error)
	PrepareContext(ctx context.Context, query string) (bun.Stmt, error)
	RunInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx InterfaceTx) error) error
	Begin() (InterfaceTx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (InterfaceTx, error)
}
