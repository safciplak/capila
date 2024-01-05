package database

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type InterfacePgStmt interface {
	Exec(params ...interface{}) (pg.Result, error)
	ExecContext(c context.Context, params ...interface{}) (pg.Result, error)
	ExecOne(params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, params ...interface{}) (pg.Result, error)
	Query(model interface{}, params ...interface{}) (pg.Result, error)
	QueryContext(c context.Context, model interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model interface{}, params ...interface{}) (pg.Result, error)
	Close() error
}
