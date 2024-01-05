package database

import (
	"context"
	"io"

	"github.com/go-pg/pg/v10/orm"
)

type adapterPgQuery struct {
	Q *orm.Query
}

func (q adapterPgQuery) TableModel() orm.TableModel {
	return q.Q.TableModel()
}

func (q adapterPgQuery) Count() (int, error) {
	return q.Q.Count()
}

func (q adapterPgQuery) First() error {
	return q.Q.First()
}

func (q adapterPgQuery) Last() error {
	return q.Q.Last()
}

func (q adapterPgQuery) Select(values ...interface{}) error {
	return q.Q.Select(values...)
}

func (q adapterPgQuery) SelectAndCount(values ...interface{}) (count int, firstErr error) {
	return q.Q.SelectAndCount(values...)
}

func (q adapterPgQuery) SelectAndCountEstimate(threshold int, values ...interface{}) (count int, firstErr error) {
	return q.Q.SelectAndCountEstimate(threshold, values...)
}

func (q adapterPgQuery) ForEach(fn interface{}) error {
	return q.Q.ForEach(fn)
}

func (q adapterPgQuery) Insert(values ...interface{}) (orm.Result, error) {
	return q.Q.Insert(values...)
}

func (q adapterPgQuery) SelectOrInsert(values ...interface{}) (inserted bool, _ error) {
	return q.Q.SelectOrInsert(values...)
}

func (q adapterPgQuery) Update(scan ...interface{}) (orm.Result, error) {
	return q.Q.Update(scan...)
}

func (q adapterPgQuery) UpdateNotZero(scan ...interface{}) (orm.Result, error) {
	return q.Q.UpdateNotZero(scan...)
}

func (q adapterPgQuery) Delete(values ...interface{}) (orm.Result, error) {
	return q.Q.Delete(values...)
}

func (q adapterPgQuery) ForceDelete(values ...interface{}) (orm.Result, error) {
	return q.Q.ForceDelete(values...)
}

func (q adapterPgQuery) CreateTable(opt *orm.CreateTableOptions) error {
	return q.Q.CreateTable(opt)
}

func (q adapterPgQuery) DropTable(opt *orm.DropTableOptions) error {
	return q.Q.DropTable(opt)
}

func (q adapterPgQuery) CreateComposite(opt *orm.CreateCompositeOptions) error {
	return q.Q.CreateComposite(opt)
}

func (q adapterPgQuery) DropComposite(opt *orm.DropCompositeOptions) error {
	return q.Q.DropComposite(opt)
}

func (q adapterPgQuery) Exec(query interface{}, params ...interface{}) (orm.Result, error) {
	return q.Q.Exec(query, params...)
}

func (q adapterPgQuery) ExecOne(query interface{}, params ...interface{}) (orm.Result, error) {
	return q.Q.ExecOne(query, params...)
}

func (q adapterPgQuery) Query(model, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.Q.Query(model, query, params...)
}

func (q adapterPgQuery) QueryOne(model, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.Q.QueryOne(model, query, params...)
}

func (q adapterPgQuery) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.Q.CopyFrom(r, query, params...)
}

func (q adapterPgQuery) CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error) {
	return q.Q.CopyTo(w, query, params...)
}

func (q adapterPgQuery) AppendQuery(fmter orm.QueryFormatter, b []byte) ([]byte, error) {
	return q.Q.AppendQuery(fmter, b)
}

func (q adapterPgQuery) Exists() (bool, error) {
	return q.Q.Exists()
}

func (q adapterPgQuery) CountEstimate(threshold int) (int, error) {
	return q.Q.CountEstimate(threshold)
}

func (q adapterPgQuery) New() InterfaceORMQuery {
	return adapterPgQuery{q.Q.New()}
}

func (q adapterPgQuery) Clone() InterfaceORMQuery {
	return adapterPgQuery{q.Q.Clone()}
}

func (q adapterPgQuery) Context(c context.Context) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Context(c)}
}

func (q adapterPgQuery) DB(db orm.DB) InterfaceORMQuery {
	return adapterPgQuery{q.Q.DB(db)}
}

func (q adapterPgQuery) Model(model ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Model(model...)}
}

func (q adapterPgQuery) Deleted() InterfaceORMQuery {
	return adapterPgQuery{q.Q.Deleted()}
}

func (q adapterPgQuery) AllWithDeleted() InterfaceORMQuery {
	return adapterPgQuery{q.Q.AllWithDeleted()}
}

func (q adapterPgQuery) With(name string, subq InterfaceORMQuery) InterfaceORMQuery {
	sub := subq.(adapterPgQuery)
	return adapterPgQuery{q.Q.With(name, sub.Q)}
}

func (q adapterPgQuery) WithInsert(name string, subq InterfaceORMQuery) InterfaceORMQuery {
	sub := subq.(adapterPgQuery)
	return adapterPgQuery{q.Q.WithInsert(name, sub.Q)}
}

func (q adapterPgQuery) WithUpdate(name string, subq InterfaceORMQuery) InterfaceORMQuery {
	sub := subq.(adapterPgQuery)
	return adapterPgQuery{q.Q.WithUpdate(name, sub.Q)}
}

func (q adapterPgQuery) WithDelete(name string, subq InterfaceORMQuery) InterfaceORMQuery {
	sub := subq.(adapterPgQuery)
	return adapterPgQuery{q.Q.WithDelete(name, sub.Q)}
}

func (q adapterPgQuery) WrapWith(name string) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WrapWith(name)}
}

func (q adapterPgQuery) Table(tables ...string) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Table(tables...)}
}

func (q adapterPgQuery) TableExpr(expr string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.TableExpr(expr, params...)}
}

func (q adapterPgQuery) Distinct() InterfaceORMQuery {
	return adapterPgQuery{q.Q.Distinct()}
}

func (q adapterPgQuery) DistinctOn(expr string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.DistinctOn(expr, params...)}
}

func (q adapterPgQuery) Column(columns ...string) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Column(columns...)}
}

func (q adapterPgQuery) ColumnExpr(expr string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.ColumnExpr(expr, params...)}
}

func (q adapterPgQuery) ExcludeColumn(columns ...string) InterfaceORMQuery {
	return adapterPgQuery{q.Q.ExcludeColumn(columns...)}
}

func (q adapterPgQuery) Relation(name string, apply ...func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery {
	var adapterFuncCollection = make([]func(query *orm.Query) (*orm.Query, error), 0)

	for _, value := range apply {
		adapterFuncCollection = append(adapterFuncCollection, adapterFunc(value))
	}

	return adapterPgQuery{q.Q.Relation(name, adapterFuncCollection...)}
}

func (q adapterPgQuery) Set(set string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Set(set, params...)}
}

func (q adapterPgQuery) Value(column, value string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Value(column, value, params...)}
}

func (q adapterPgQuery) Where(condition string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Where(condition, params...)}
}

func (q adapterPgQuery) WhereOr(condition string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WhereOr(condition, params...)}
}

func (q adapterPgQuery) WhereGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WhereGroup(adapterFunc(fn))}
}

func (q adapterPgQuery) WhereNotGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WhereNotGroup(adapterFunc(fn))}
}

func (q adapterPgQuery) WhereOrGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WhereOrGroup(adapterFunc(fn))}
}

func (q adapterPgQuery) WhereOrNotGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WhereOrNotGroup(adapterFunc(fn))}
}

func (q adapterPgQuery) WhereIn(where string, slice interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WhereIn(where, slice)}
}

func (q adapterPgQuery) WhereInMulti(where string, values ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.WhereInMulti(where, values...)}
}

func (q adapterPgQuery) WherePK() InterfaceORMQuery {
	return adapterPgQuery{q.Q.WherePK()}
}

func (q adapterPgQuery) Join(join string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Join(join, params...)}
}

func (q adapterPgQuery) JoinOn(condition string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.JoinOn(condition, params...)}
}

func (q adapterPgQuery) JoinOnOr(condition string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.JoinOnOr(condition, params...)}
}

func (q adapterPgQuery) Group(columns ...string) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Group(columns...)}
}

func (q adapterPgQuery) GroupExpr(group string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.GroupExpr(group, params...)}
}

func (q adapterPgQuery) Having(having string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Having(having, params...)}
}

func (q adapterPgQuery) Union(other InterfaceORMQuery) InterfaceORMQuery {
	sub := other.(adapterPgQuery)
	return adapterPgQuery{q.Q.Union(sub.Q)}
}

func (q adapterPgQuery) UnionAll(other InterfaceORMQuery) InterfaceORMQuery {
	sub := other.(adapterPgQuery)
	return adapterPgQuery{q.Q.UnionAll(sub.Q)}
}

func (q adapterPgQuery) Intersect(other InterfaceORMQuery) InterfaceORMQuery {
	sub := other.(adapterPgQuery)
	return adapterPgQuery{q.Q.Intersect(sub.Q)}
}

func (q adapterPgQuery) IntersectAll(other InterfaceORMQuery) InterfaceORMQuery {
	sub := other.(adapterPgQuery)
	return adapterPgQuery{q.Q.IntersectAll(sub.Q)}
}

func (q adapterPgQuery) Except(other InterfaceORMQuery) InterfaceORMQuery {
	sub := other.(adapterPgQuery)
	return adapterPgQuery{q.Q.Except(sub.Q)}
}

func (q adapterPgQuery) ExceptAll(other InterfaceORMQuery) InterfaceORMQuery {
	sub := other.(adapterPgQuery)
	return adapterPgQuery{q.Q.ExceptAll(sub.Q)}
}

func (q adapterPgQuery) Order(orders ...string) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Order(orders...)}
}

func (q adapterPgQuery) OrderExpr(order string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.OrderExpr(order, params...)}
}

func (q adapterPgQuery) Limit(n int) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Limit(n)}
}

func (q adapterPgQuery) Offset(n int) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Offset(n)}
}

func (q adapterPgQuery) OnConflict(s string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.OnConflict(s, params...)}
}

func (q adapterPgQuery) Returning(s string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Returning(s, params...)}
}

func (q adapterPgQuery) For(s string, params ...interface{}) InterfaceORMQuery {
	return adapterPgQuery{q.Q.For(s, params...)}
}

func (q adapterPgQuery) Apply(fn func(q InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery {
	return adapterPgQuery{q.Q.Apply(adapterFunc(fn))}
}

func adapterFunc(fn func(q InterfaceORMQuery) (InterfaceORMQuery, error)) func(tx *orm.Query) (*orm.Query, error) {
	return func(tx *orm.Query) (*orm.Query, error) {
		result, err := fn(adapterPgQuery{tx})
		return result.(adapterPgQuery).Q, err
	}
}
