package database

import (
	"context"
	"io"

	"github.com/go-pg/pg/v10/orm"
)

// InterfaceORMQuery contains the ORM Query interface
type InterfaceORMQuery interface {
	New() InterfaceORMQuery
	Clone() InterfaceORMQuery
	Context(c context.Context) InterfaceORMQuery
	DB(db orm.DB) InterfaceORMQuery
	Model(model ...interface{}) InterfaceORMQuery
	TableModel() orm.TableModel
	Deleted() InterfaceORMQuery
	AllWithDeleted() InterfaceORMQuery
	With(name string, subq InterfaceORMQuery) InterfaceORMQuery
	WithInsert(name string, subq InterfaceORMQuery) InterfaceORMQuery
	WithUpdate(name string, subq InterfaceORMQuery) InterfaceORMQuery
	WithDelete(name string, subq InterfaceORMQuery) InterfaceORMQuery
	WrapWith(name string) InterfaceORMQuery
	Table(tables ...string) InterfaceORMQuery
	TableExpr(expr string, params ...interface{}) InterfaceORMQuery
	Distinct() InterfaceORMQuery
	DistinctOn(expr string, params ...interface{}) InterfaceORMQuery
	Column(columns ...string) InterfaceORMQuery
	ColumnExpr(expr string, params ...interface{}) InterfaceORMQuery
	ExcludeColumn(columns ...string) InterfaceORMQuery
	Relation(name string, apply ...func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery
	Set(set string, params ...interface{}) InterfaceORMQuery
	Value(column string, value string, params ...interface{}) InterfaceORMQuery
	Where(condition string, params ...interface{}) InterfaceORMQuery
	WhereOr(condition string, params ...interface{}) InterfaceORMQuery
	WhereGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery
	WhereNotGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery
	WhereOrGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery
	WhereOrNotGroup(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery
	WhereIn(where string, slice interface{}) InterfaceORMQuery
	WhereInMulti(where string, values ...interface{}) InterfaceORMQuery
	WherePK() InterfaceORMQuery
	Join(join string, params ...interface{}) InterfaceORMQuery
	JoinOn(condition string, params ...interface{}) InterfaceORMQuery
	JoinOnOr(condition string, params ...interface{}) InterfaceORMQuery
	Group(columns ...string) InterfaceORMQuery
	GroupExpr(group string, params ...interface{}) InterfaceORMQuery
	Having(having string, params ...interface{}) InterfaceORMQuery
	Union(other InterfaceORMQuery) InterfaceORMQuery
	UnionAll(other InterfaceORMQuery) InterfaceORMQuery
	Intersect(other InterfaceORMQuery) InterfaceORMQuery
	IntersectAll(other InterfaceORMQuery) InterfaceORMQuery
	Except(other InterfaceORMQuery) InterfaceORMQuery
	ExceptAll(other InterfaceORMQuery) InterfaceORMQuery
	Order(orders ...string) InterfaceORMQuery
	OrderExpr(order string, params ...interface{}) InterfaceORMQuery
	Limit(n int) InterfaceORMQuery
	Offset(n int) InterfaceORMQuery
	OnConflict(s string, params ...interface{}) InterfaceORMQuery
	Returning(s string, params ...interface{}) InterfaceORMQuery
	For(s string, params ...interface{}) InterfaceORMQuery
	Apply(fn func(InterfaceORMQuery) (InterfaceORMQuery, error)) InterfaceORMQuery
	Count() (int, error)
	First() error
	Last() error
	Select(values ...interface{}) error
	SelectAndCount(values ...interface{}) (count int, firstErr error)
	SelectAndCountEstimate(threshold int, values ...interface{}) (count int, firstErr error)
	ForEach(fn interface{}) error
	Insert(values ...interface{}) (orm.Result, error)
	SelectOrInsert(values ...interface{}) (inserted bool, _ error)
	Update(scan ...interface{}) (orm.Result, error)
	UpdateNotZero(scan ...interface{}) (orm.Result, error)
	Delete(values ...interface{}) (orm.Result, error)
	ForceDelete(values ...interface{}) (orm.Result, error)
	CreateTable(opt *orm.CreateTableOptions) error
	DropTable(opt *orm.DropTableOptions) error
	CreateComposite(opt *orm.CreateCompositeOptions) error
	DropComposite(opt *orm.DropCompositeOptions) error
	Exec(query interface{}, params ...interface{}) (orm.Result, error)
	ExecOne(query interface{}, params ...interface{}) (orm.Result, error)
	Query(model, query interface{}, params ...interface{}) (orm.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (orm.Result, error)
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error)
	AppendQuery(fmter orm.QueryFormatter, b []byte) ([]byte, error)
	Exists() (bool, error)
	CountEstimate(threshold int) (int, error)
}
