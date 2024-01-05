package database

import (
	"github.com/go-pg/pg/v10/orm"
)

// HandleQueryResult just returns the error for now.
func HandleQueryResult(_ orm.Result, err error) error {
	return err
}
