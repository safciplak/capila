package database

import (
	"context"
	"io"
	"time"

	v10 "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// InterfacePGDB contains the Postgress connection interface
type InterfacePGDB interface {
	Context() context.Context
	WithContext(ctx context.Context) InterfacePGDB
	WithTimeout(d time.Duration) InterfacePGDB
	WithParam(param string, value interface{}) InterfacePGDB
	AddQueryHook(hook v10.QueryHook)
	PoolStats() *v10.PoolStats
	Param(param string) interface{}
	Close() error
	Exec(query interface{}, params ...interface{}) (res v10.Result, err error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (v10.Result, error)
	ExecOne(query interface{}, params ...interface{}) (v10.Result, error)
	ExecOneContext(ctx context.Context, query interface{}, params ...interface{}) (v10.Result, error)
	Query(model, query interface{}, params ...interface{}) (res v10.Result, err error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (v10.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (v10.Result, error)
	QueryOneContext(
		ctx context.Context, model, query interface{}, params ...interface{},
	) (v10.Result, error)
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res v10.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res v10.Result, err error)
	Ping(ctx context.Context) error
	Model(model ...interface{}) InterfaceORMQuery
	ModelContext(c context.Context, model ...interface{}) InterfaceORMQuery
	Formatter() orm.QueryFormatter
	Prepare(q string) (InterfacePgStmt, error)
	Begin() (InterfacePgTx, error)
	BeginContext(ctx context.Context) (InterfacePgTx, error)
	RunInTransaction(ctx context.Context, fn func(InterfacePgTx) error) error
}
