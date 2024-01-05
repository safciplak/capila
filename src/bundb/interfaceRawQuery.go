package bundb

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceRawQuery is the interface for the raw query option
type InterfaceRawQuery interface {
	Conn(db bun.IConn) InterfaceRawQuery
	Scan(ctx context.Context, dest ...interface{}) error
	AppendQuery(formatter schema.Formatter, b []byte) ([]byte, error)
	Operation() string
}
