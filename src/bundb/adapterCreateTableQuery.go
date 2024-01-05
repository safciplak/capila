package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterCreateTableQuery struct {
	Q *bun.CreateTableQuery
}

func (a adapterCreateTableQuery) Conn(db bun.IConn) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.Conn(db)}
}

func (a adapterCreateTableQuery) Model(model interface{}) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.Model(model)}
}

func (a adapterCreateTableQuery) Table(tables ...string) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.Table(tables...)}
}

func (a adapterCreateTableQuery) TableExpr(query string, args ...interface{}) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterCreateTableQuery) ModelTableExpr(query string, args ...interface{}) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterCreateTableQuery) ColumnExpr(query string, args ...interface{}) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.ColumnExpr(query, args...)}
}

func (a adapterCreateTableQuery) Temp() InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.Temp()}
}

func (a adapterCreateTableQuery) IfNotExists() InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.IfNotExists()}
}

func (a adapterCreateTableQuery) Varchar(n int) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.Varchar(n)}
}

func (a adapterCreateTableQuery) ForeignKey(query string, args ...interface{}) InterfaceCreateTableQuery {
	return adapterCreateTableQuery{a.Q.ForeignKey(query, args...)}
}

func (a adapterCreateTableQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterCreateTableQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterCreateTableQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
