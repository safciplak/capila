package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// InterfaceSelectQuery represents the bun.SelectQuery struct
type InterfaceSelectQuery interface {
	Conn(db bun.IConn) InterfaceSelectQuery
	Model(model interface{}) InterfaceSelectQuery
	Apply(fn func(InterfaceSelectQuery) InterfaceSelectQuery) InterfaceSelectQuery
	With(name string, query schema.QueryAppender) InterfaceSelectQuery
	Distinct() InterfaceSelectQuery
	DistinctOn(query string, args ...interface{}) InterfaceSelectQuery
	Table(tables ...string) InterfaceSelectQuery
	TableExpr(query string, args ...interface{}) InterfaceSelectQuery
	ModelTableExpr(query string, args ...interface{}) InterfaceSelectQuery
	Column(columns ...string) InterfaceSelectQuery
	ColumnExpr(query string, args ...interface{}) InterfaceSelectQuery
	ExcludeColumn(columns ...string) InterfaceSelectQuery
	WherePK(cols ...string) InterfaceSelectQuery
	Where(query string, args ...interface{}) InterfaceSelectQuery
	WhereOr(query string, args ...interface{}) InterfaceSelectQuery
	WhereGroup(sep string, fn func(InterfaceSelectQuery) InterfaceSelectQuery) InterfaceSelectQuery
	WhereDeleted() InterfaceSelectQuery
	WhereAllWithDeleted() InterfaceSelectQuery
	Group(columns ...string) InterfaceSelectQuery
	GroupExpr(group string, args ...interface{}) InterfaceSelectQuery
	Having(having string, args ...interface{}) InterfaceSelectQuery
	Order(orders ...string) InterfaceSelectQuery
	OrderExpr(query string, args ...interface{}) InterfaceSelectQuery
	Limit(n int) InterfaceSelectQuery
	Offset(n int) InterfaceSelectQuery
	For(s string, args ...interface{}) InterfaceSelectQuery
	Union(other InterfaceSelectQuery) InterfaceSelectQuery
	UnionAll(other InterfaceSelectQuery) InterfaceSelectQuery
	Intersect(other InterfaceSelectQuery) InterfaceSelectQuery
	IntersectAll(other InterfaceSelectQuery) InterfaceSelectQuery
	Except(other InterfaceSelectQuery) InterfaceSelectQuery
	ExceptAll(other InterfaceSelectQuery) InterfaceSelectQuery
	Join(join string, args ...interface{}) InterfaceSelectQuery
	JoinOn(cond string, args ...interface{}) InterfaceSelectQuery
	JoinOnOr(cond string, args ...interface{}) InterfaceSelectQuery
	Relation(name string, apply ...func(InterfaceSelectQuery) InterfaceSelectQuery) InterfaceSelectQuery
	Operation() string
	AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error)
	Rows(ctx context.Context) (*sql.Rows, error)
	Exec(ctx context.Context) (res sql.Result, err error)
	Scan(ctx context.Context, dest ...interface{}) error
	Count(ctx context.Context) (int, error)
	ScanAndCount(ctx context.Context, dest ...interface{}) (int, error)
	Exists(ctx context.Context) (bool, error)
}
