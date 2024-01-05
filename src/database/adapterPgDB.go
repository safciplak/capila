package database

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

// adapterPgDB contains the Database
type adapterPgDB struct {
	*pg.DB
}

// WithContext wraps the WithContext
func (a adapterPgDB) WithContext(ctx context.Context) InterfacePGDB {
	return adapterPgDB{a.DB.WithContext(ctx)}
}

// WithTimeout wraps the WithTimeout
func (a adapterPgDB) WithTimeout(d time.Duration) InterfacePGDB {
	return adapterPgDB{a.DB.WithTimeout(d)}
}

// WithParam wraps the WithParam
func (a adapterPgDB) WithParam(param string, value interface{}) InterfacePGDB {
	return adapterPgDB{a.DB.WithParam(param, value)}
}

// Model wraps the Model
func (a adapterPgDB) Model(model ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{a.DB.Model(model...)}
}

// ModelContext wraps the ModelContext
func (a adapterPgDB) ModelContext(c context.Context, model ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{a.DB.ModelContext(c, model...)}
}

// Prepare wraps the Prepare
func (a adapterPgDB) Prepare(q string) (InterfacePgStmt, error) {
	var statement, err = a.DB.Prepare(q)
	return adapterPgStmt{statement}, err
}

// Begin wraps the Begin
func (a adapterPgDB) Begin() (InterfacePgTx, error) {
	var tx, err = a.DB.Begin()
	return adapterPgTx{tx}, err
}

// BeginContext wraps the BeginContext
func (a adapterPgDB) BeginContext(ctx context.Context) (InterfacePgTx, error) {
	var tx, err = a.DB.BeginContext(ctx)
	return adapterPgTx{tx}, err
}

// RunInTransaction wraps the RunInTransaction
func (a adapterPgDB) RunInTransaction(ctx context.Context, fn func(InterfacePgTx) error) error {
	adapterFunc := func(tx *pg.Tx) error {
		err := fn(adapterPgTx{tx})
		return err
	}

	return a.DB.RunInTransaction(ctx, adapterFunc)
}
