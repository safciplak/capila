package bundb

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type adapterSelectQuery struct {
	Q *bun.SelectQuery
}

func adapterSelectQueryFunc(fn func(q InterfaceSelectQuery) InterfaceSelectQuery) func(q *bun.SelectQuery) *bun.SelectQuery {
	return func(q *bun.SelectQuery) *bun.SelectQuery {
		result := fn(adapterSelectQuery{q})
		return result.(adapterSelectQuery).Q
	}
}

func (a adapterSelectQuery) Conn(db bun.IConn) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Conn(db)}
}

func (a adapterSelectQuery) Model(model interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Model(model)}
}

func (a adapterSelectQuery) Apply(fn func(q InterfaceSelectQuery) InterfaceSelectQuery) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Apply(adapterSelectQueryFunc(fn))}
}

func (a adapterSelectQuery) With(name string, query schema.QueryAppender) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.With(name, query)}
}

func (a adapterSelectQuery) Distinct() InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Distinct()}
}

func (a adapterSelectQuery) DistinctOn(query string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.DistinctOn(query, args...)}
}

func (a adapterSelectQuery) Table(tables ...string) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Table(tables...)}
}

func (a adapterSelectQuery) TableExpr(query string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.TableExpr(query, args...)}
}

func (a adapterSelectQuery) ModelTableExpr(query string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.ModelTableExpr(query, args...)}
}

func (a adapterSelectQuery) Column(columns ...string) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Column(columns...)}
}

func (a adapterSelectQuery) ColumnExpr(query string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.ColumnExpr(query, args...)}
}

func (a adapterSelectQuery) ExcludeColumn(columns ...string) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.ExcludeColumn(columns...)}
}

func (a adapterSelectQuery) WherePK(cols ...string) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.WherePK(cols...)}
}

func (a adapterSelectQuery) Where(query string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Where(query, args...)}
}

func (a adapterSelectQuery) WhereOr(query string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.WhereOr(query, args...)}
}

func (a adapterSelectQuery) WhereGroup(sep string, fn func(InterfaceSelectQuery) InterfaceSelectQuery) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.WhereGroup(sep, adapterSelectQueryFunc(fn))}
}

func (a adapterSelectQuery) WhereDeleted() InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.WhereDeleted()}
}

func (a adapterSelectQuery) WhereAllWithDeleted() InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.WhereAllWithDeleted()}
}

func (a adapterSelectQuery) Group(columns ...string) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Group(columns...)}
}

func (a adapterSelectQuery) GroupExpr(group string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.GroupExpr(group, args...)}
}

func (a adapterSelectQuery) Having(having string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Having(having, args...)}
}

func (a adapterSelectQuery) Order(orders ...string) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Order(orders...)}
}

func (a adapterSelectQuery) OrderExpr(query string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.OrderExpr(query, args...)}
}

func (a adapterSelectQuery) Limit(n int) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Limit(n)}
}

func (a adapterSelectQuery) Offset(n int) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Offset(n)}
}

func (a adapterSelectQuery) For(s string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.For(s, args...)}
}

func (a adapterSelectQuery) Union(other InterfaceSelectQuery) InterfaceSelectQuery {
	sub := other.(adapterSelectQuery)
	return adapterSelectQuery{a.Q.Union(sub.Q)}
}

func (a adapterSelectQuery) UnionAll(other InterfaceSelectQuery) InterfaceSelectQuery {
	sub := other.(adapterSelectQuery)
	return adapterSelectQuery{a.Q.UnionAll(sub.Q)}
}

func (a adapterSelectQuery) Intersect(other InterfaceSelectQuery) InterfaceSelectQuery {
	sub := other.(adapterSelectQuery)
	return adapterSelectQuery{a.Q.Intersect(sub.Q)}
}

func (a adapterSelectQuery) IntersectAll(other InterfaceSelectQuery) InterfaceSelectQuery {
	sub := other.(adapterSelectQuery)
	return adapterSelectQuery{a.Q.IntersectAll(sub.Q)}
}

func (a adapterSelectQuery) Except(other InterfaceSelectQuery) InterfaceSelectQuery {
	sub := other.(adapterSelectQuery)
	return adapterSelectQuery{a.Q.Except(sub.Q)}
}

func (a adapterSelectQuery) ExceptAll(other InterfaceSelectQuery) InterfaceSelectQuery {
	sub := other.(adapterSelectQuery)
	return adapterSelectQuery{a.Q.ExceptAll(sub.Q)}
}

func (a adapterSelectQuery) Join(join string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.Join(join, args...)}
}

func (a adapterSelectQuery) JoinOn(cond string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.JoinOn(cond, args...)}
}

func (a adapterSelectQuery) JoinOnOr(cond string, args ...interface{}) InterfaceSelectQuery {
	return adapterSelectQuery{a.Q.JoinOnOr(cond, args...)}
}

func (a adapterSelectQuery) Relation(name string, apply ...func(InterfaceSelectQuery) InterfaceSelectQuery) InterfaceSelectQuery {
	adapterFuncCollection := make([]func(query *bun.SelectQuery) *bun.SelectQuery, 0)

	for _, value := range apply {
		adapterFuncCollection = append(adapterFuncCollection, adapterSelectQueryFunc(value))
	}

	return adapterSelectQuery{a.Q.Relation(name, adapterFuncCollection...)}
}

func (a adapterSelectQuery) Rows(ctx context.Context) (*sql.Rows, error) {
	return a.Q.Rows(ctx)
}

func (a adapterSelectQuery) Count(ctx context.Context) (int, error) {
	return a.Q.Count(ctx)
}

func (a adapterSelectQuery) ScanAndCount(ctx context.Context, dest ...interface{}) (int, error) {
	return a.Q.ScanAndCount(ctx, dest...)
}

func (a adapterSelectQuery) Exists(ctx context.Context) (bool, error) {
	return a.Q.Exists(ctx)
}

func (a adapterSelectQuery) Operation() string {
	return a.Q.Operation()
}

func (a adapterSelectQuery) AppendQuery(fmter schema.Formatter, b []byte) (_ []byte, err error) {
	return a.Q.AppendQuery(fmter, b)
}

func (a adapterSelectQuery) Scan(ctx context.Context, dest ...interface{}) error {
	return a.Q.Scan(ctx, dest...)
}

func (a adapterSelectQuery) Exec(ctx context.Context) (sql.Result, error) {
	return a.Q.Exec(ctx)
}
