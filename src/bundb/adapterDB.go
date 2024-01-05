package bundb

import (
	"context"
	"database/sql"
	"reflect"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterDB struct {
	*bun.DB
}

func (a adapterDB) String() string {
	return a.DB.String()
}

func (a adapterDB) DBStats() bun.DBStats {
	return a.DB.DBStats()
}

func (a adapterDB) NewValues(model interface{}) InterfaceValuesQuery {
	return adapterValuesQuery{a.DB.NewValues(model)}
}

func (a adapterDB) NewSelect() InterfaceSelectQuery {
	return adapterSelectQuery{a.DB.NewSelect()}
}

func (a adapterDB) NewInsert() InterfaceInsertQuery {
	return adapterInsertQuery{a.DB.NewInsert()}
}

func (a adapterDB) NewUpdate() InterfaceUpdateQuery {
	return adapterUpdateQuery{a.DB.NewUpdate()}
}

func (a adapterDB) NewDelete() InterfaceDeleteQuery {
	return adapterDeleteQuery{a.DB.NewDelete()}
}

func (a adapterDB) NewCreateTable() InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.DB.NewCreateTable()}
}

func (a adapterDB) NewDropTable() InterfaceDropTableQuery {
	return adapterDropTableQuery{a.DB.NewDropTable()}
}

func (a adapterDB) NewCreateIndex() InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.DB.NewCreateIndex()}
}

func (a adapterDB) NewDropIndex() InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.DB.NewDropIndex()}
}

func (a adapterDB) NewTruncateTable() InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.DB.NewTruncateTable()}
}

func (a adapterDB) NewAddColumn() InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.DB.NewAddColumn()}
}

func (a adapterDB) NewDropColumn() InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.DB.NewDropColumn()}
}

func (a adapterDB) ResetModel(ctx context.Context, models ...interface{}) error {
	return a.DB.ResetModel(ctx, models...)
}

func (a adapterDB) Dialect() schema.Dialect {
	return a.DB.Dialect()
}

func (a adapterDB) ScanRows(ctx context.Context, rows *sql.Rows, dest ...interface{}) error {
	return a.DB.ScanRows(ctx, rows, dest...)
}

func (a adapterDB) ScanRow(ctx context.Context, rows *sql.Rows, dest ...interface{}) error {
	return a.DB.ScanRow(ctx, rows, dest...)
}

func (a adapterDB) AddQueryHook(hook bun.QueryHook) {
	a.DB.AddQueryHook(hook)
}

func (a adapterDB) Table(typ reflect.Type) *schema.Table {
	return a.DB.Table(typ)
}

func (a adapterDB) RegisterModel(models ...interface{}) {
	a.DB.RegisterModel(models...)
}

func (a adapterDB) WithNamedArg(name string, value interface{}) InterfaceDB {
	return adapterDB{a.DB.WithNamedArg(name, value)}
}

func (a adapterDB) Formatter() schema.Formatter {
	return a.DB.Formatter()
}

func (a adapterDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return a.DB.Exec(query, args...)
}

func (a adapterDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return a.DB.ExecContext(ctx, query, args...)
}

func (a adapterDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return a.DB.Query(query, args...)
}

func (a adapterDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return a.DB.QueryContext(ctx, query, args...)
}

func (a adapterDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return a.DB.QueryRow(query, args...)
}

func (a adapterDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return a.DB.QueryRowContext(ctx, query, args...)
}

func (a adapterDB) Prepare(query string) (bun.Stmt, error) {
	return a.DB.Prepare(query)
}

func (a adapterDB) PrepareContext(ctx context.Context, query string) (bun.Stmt, error) {
	return a.DB.PrepareContext(ctx, query)
}

func (a adapterDB) RunInTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx InterfaceTx) error) error {
	adapterFunc := func(ctx context.Context, tx bun.Tx) error {
		err := fn(ctx, adapterTx{tx})
		return err
	}

	return a.DB.RunInTx(ctx, opts, adapterFunc)
}

func (a adapterDB) Begin() (InterfaceTx, error) {
	tx, err := a.DB.Begin()
	return adapterTx{tx}, err
}

func (a adapterDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (InterfaceTx, error) {
	tx, err := a.DB.BeginTx(ctx, opts)
	return adapterTx{tx}, err
}

func (a adapterDB) NewRaw(query string, args ...interface{}) InterfaceRawQuery {
	return adapterRawQuery{a.DB.NewRaw(query, args...)}
}
