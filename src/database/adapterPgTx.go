package database

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// adapterPgTx contains the pg transaction
type adapterPgTx struct {
	*pg.Tx
}

// Begin wraps the Begin
func (a adapterPgTx) Begin() (InterfacePgTx, error) {
	var tx, err = a.Tx.Begin()
	return adapterPgTx{tx}, err
}

// RunInTransaction wraps the RunInTransaction
func (a adapterPgTx) RunInTransaction(ctx context.Context, fn func(InterfacePgTx) error) error {
	adapterFunc := func(tx *pg.Tx) error {
		err := fn(adapterPgTx{tx})
		return err
	}

	return a.Tx.RunInTransaction(ctx, adapterFunc)
}

// Stmt wraps the Stmt
func (a adapterPgTx) Stmt(stmt InterfacePgStmt) InterfacePgStmt {
	sub := stmt.(adapterPgStmt)
	return adapterPgStmt{a.Tx.Stmt(sub.Stmt)}
}

// Prepare wraps the Prepare
func (a adapterPgTx) Prepare(q string) (InterfacePgStmt, error) {
	stmt, err := a.Tx.Prepare(q)
	return adapterPgStmt{stmt}, err
}

// Model wraps the Model
func (a adapterPgTx) Model(model ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{a.Tx.Model(model...)}
}

// ModelContext wraps the ModelContext
func (a adapterPgTx) ModelContext(c context.Context, model ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{a.Tx.ModelContext(c, model...)}
}
