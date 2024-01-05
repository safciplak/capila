package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type adapterTx struct {
	bun.Tx
}

func (a adapterTx) Commit() error {
	return a.Tx.Commit()
}

func (a adapterTx) Rollback() error {
	return a.Tx.Rollback()
}

func (a adapterTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return a.Tx.Exec(query, args...)
}

func (a adapterTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return a.Tx.Query(query, args...)
}

func (a adapterTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return a.Tx.QueryRow(query, args...)
}

func (a adapterTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return a.Tx.ExecContext(ctx, query, args...)
}

func (a adapterTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return a.Tx.QueryContext(ctx, query, args...)
}

func (a adapterTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return a.Tx.QueryRowContext(ctx, query, args...)
}

func (a adapterTx) NewValues(model interface{}) InterfaceValuesQuery {
	return adapterValuesQuery{a.Tx.NewValues(model)}
}

func (a adapterTx) NewSelect() InterfaceSelectQuery {
	return adapterSelectQuery{a.Tx.NewSelect()}
}

func (a adapterTx) NewInsert() InterfaceInsertQuery {
	return adapterInsertQuery{a.Tx.NewInsert()}
}

func (a adapterTx) NewUpdate() InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Tx.NewUpdate()}
}

func (a adapterTx) NewDelete() InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Tx.NewDelete()}
}

func (a adapterTx) NewCreateTable() InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Tx.NewCreateTable()}
}

func (a adapterTx) NewDropTable() InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Tx.NewDropTable()}
}

func (a adapterTx) NewCreateIndex() InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Tx.NewCreateIndex()}
}

func (a adapterTx) NewDropIndex() InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Tx.NewDropIndex()}
}

func (a adapterTx) NewTruncateTable() InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Tx.NewTruncateTable()}
}

func (a adapterTx) NewAddColumn() InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Tx.NewAddColumn()}
}

func (a adapterTx) NewDropColumn() InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Tx.NewDropColumn()}
}
