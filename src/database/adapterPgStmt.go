package database

import (
	"github.com/go-pg/pg/v10"
)

// adapterPgStmt contains a single pg statement
type adapterPgStmt struct {
	*pg.Stmt
}
