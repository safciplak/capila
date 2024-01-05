package bundb

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceValuesQuery represents the bun.ValuesQuery struct
type InterfaceValuesQuery interface {
	Conn(db bun.IConn) InterfaceValuesQuery
	Value(column string, expr string, args ...interface{}) InterfaceValuesQuery
	WithOrder() InterfaceValuesQuery
	AppendNamedArg(fmter schema.Formatter, b []byte, name string) ([]byte, bool)
	AppendColumns(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
}
