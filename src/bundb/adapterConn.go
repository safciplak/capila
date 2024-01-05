package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type adapterConn struct {
	bun.Conn
}

func (a adapterConn) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return a.Conn.ExecContext(ctx, query, args...)
}

func (a adapterConn) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return a.Conn.QueryContext(ctx, query, args...)
}

func (a adapterConn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return a.Conn.QueryRowContext(ctx, query, args...)
}

func (a adapterConn) NewValues(model interface{}) InterfaceValuesQuery {
	return adapterValuesQuery{a.Conn.NewValues(model)}
}

func (a adapterConn) NewSelect() InterfaceSelectQuery {
	return adapterSelectQuery{a.Conn.NewSelect()}
}

func (a adapterConn) NewInsert() InterfaceInsertQuery {
	return adapterInsertQuery{a.Conn.NewInsert()}
}

func (a adapterConn) NewUpdate() InterfaceUpdateQuery {
	return adapterUpdateQuery{a.Conn.NewUpdate()}
}

func (a adapterConn) NewDelete() InterfaceDeleteQuery {
	return adapterDeleteQuery{a.Conn.NewDelete()}
}

func (a adapterConn) NewCreateTable() InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Conn.NewCreateTable()}
}

func (a adapterConn) NewDropTable() InterfaceDropTableQuery {
	return adapterDropTableQuery{a.Conn.NewDropTable()}
}

func (a adapterConn) NewCreateIndex() InterfaceCreateIndexQuery {
	return adapterCreateIndexQuery{a.Conn.NewCreateIndex()}
}

func (a adapterConn) NewDropIndex() InterfaceDropIndexQuery {
	return adapterDropIndexQuery{a.Conn.NewDropIndex()}
}

func (a adapterConn) NewTruncateTable() InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Conn.NewTruncateTable()}
}

func (a adapterConn) NewAddColumn() InterfaceAddColumnQuery {
	return adapterAddColumnQuery{a.Conn.NewAddColumn()}
}

func (a adapterConn) NewDropColumn() InterfaceDropColumnQuery {
	return adapterDropColumnQuery{a.Conn.NewDropColumn()}
}
