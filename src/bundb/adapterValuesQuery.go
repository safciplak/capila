package bundb

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterValuesQuery struct {
	Q *bun.ValuesQuery
}

func (a adapterValuesQuery) Conn(db bun.IConn) InterfaceValuesQuery {
	return adapterValuesQuery{a.Q.Conn(db)}
}

func (a adapterValuesQuery) Value(column, expr string, args ...interface{}) InterfaceValuesQuery {
	return adapterValuesQuery{a.Q.Value(column, expr, args...)}
}

func (a adapterValuesQuery) WithOrder() InterfaceValuesQuery {
	return adapterValuesQuery{a.Q.WithOrder()}
}

func (a adapterValuesQuery) AppendNamedArg(fmter schema.Formatter, b []byte, name string) ([]byte, bool) {
	return a.Q.AppendNamedArg(fmter, b, name)
}

func (a adapterValuesQuery) AppendColumns(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendColumns(fmter, b)
}

func (a adapterValuesQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterValuesQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}
