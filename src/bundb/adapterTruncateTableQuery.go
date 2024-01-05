package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterTruncateTableQuery struct {
	Q *bun.TruncateTableQuery
}

func (a adapterTruncateTableQuery) Conn(db bun.IConn) InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Q.Conn(db)}
}

func (a adapterTruncateTableQuery) Model(model interface{}) InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Q.Model(model)}
}
func (a adapterTruncateTableQuery) Table(tables ...string) InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Q.Table(tables...)}
}

func (a adapterTruncateTableQuery) TableExpr(query string, args ...interface{}) InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterTruncateTableQuery) ContinueIdentity() InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Q.ContinueIdentity()}
}

func (a adapterTruncateTableQuery) Restrict() InterfaceTruncateTableQuery {
	return adapterTruncateTableQuery{a.Q.Restrict()}
}

func (a adapterTruncateTableQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterTruncateTableQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterTruncateTableQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	return a.Q.Exec(ctx, dest...)
}
