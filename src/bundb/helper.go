package bundb

import "database/sql"

// HandleQueryResult just returns the error for now.
func HandleQueryResult(_ sql.Result, err error) error {
	return err
}
