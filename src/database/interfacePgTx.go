package database

import (
	"context"
	"io"

	v10 "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type InterfacePgTx interface {
	Context() context.Context
	Begin() (InterfacePgTx, error)
	RunInTransaction(ctx context.Context, fn func(InterfacePgTx) error) error
	Stmt(stmt InterfacePgStmt) InterfacePgStmt
	Prepare(q string) (InterfacePgStmt, error)
	Exec(query interface{}, params ...interface{}) (v10.Result, error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (v10.Result, error)
	ExecOne(query interface{}, params ...interface{}) (v10.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (v10.Result, error)
	Query(model interface{}, query interface{}, params ...interface{}) (v10.Result, error)
	QueryContext(
		c context.Context,
		model interface{},
		query interface{},
		params ...interface{},
	) (v10.Result, error)
	QueryOne(model interface{}, query interface{}, params ...interface{}) (v10.Result, error)
	Model(model ...interface{}) InterfaceORMQuery
	ModelContext(c context.Context, model ...interface{}) InterfaceORMQuery
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res v10.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res v10.Result, err error)
	Formatter() orm.QueryFormatter
	Commit() error
	CommitContext(ctx context.Context) error
	Rollback() error
	RollbackContext(ctx context.Context) error
	Close() error
	CloseContext(ctx context.Context) error
}
