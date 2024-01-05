// Code generated by mockery. DO NOT EDIT.

package bundb

import (
	context "context"

	bun "github.com/uptrace/bun"

	mock "github.com/stretchr/testify/mock"

	schema "github.com/uptrace/bun/schema"

	sql "database/sql"
)

// MockInterfaceCreateTableQuery is an autogenerated mock type for the InterfaceCreateTableQuery type
type MockInterfaceCreateTableQuery struct {
	mock.Mock
}

// AppendQuery provides a mock function with given fields: fmter, b
func (_m *MockInterfaceCreateTableQuery) AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error) {
	ret := _m.Called(fmter, b)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(schema.Formatter, []byte) []byte); ok {
		r0 = rf(fmter, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(schema.Formatter, []byte) error); ok {
		r1 = rf(fmter, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ColumnExpr provides a mock function with given fields: query, args
func (_m *MockInterfaceCreateTableQuery) ColumnExpr(query string, args ...interface{}) InterfaceCreateTableQuery {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(string, ...interface{}) InterfaceCreateTableQuery); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// Conn provides a mock function with given fields: db
func (_m *MockInterfaceCreateTableQuery) Conn(db bun.IConn) InterfaceCreateTableQuery {
	ret := _m.Called(db)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(bun.IConn) InterfaceCreateTableQuery); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// Exec provides a mock function with given fields: ctx, dest
func (_m *MockInterfaceCreateTableQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, dest...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, dest...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...interface{}) error); ok {
		r1 = rf(ctx, dest...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForeignKey provides a mock function with given fields: query, args
func (_m *MockInterfaceCreateTableQuery) ForeignKey(query string, args ...interface{}) InterfaceCreateTableQuery {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(string, ...interface{}) InterfaceCreateTableQuery); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// IfNotExists provides a mock function with given fields:
func (_m *MockInterfaceCreateTableQuery) IfNotExists() InterfaceCreateTableQuery {
	ret := _m.Called()

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func() InterfaceCreateTableQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// Model provides a mock function with given fields: model
func (_m *MockInterfaceCreateTableQuery) Model(model interface{}) InterfaceCreateTableQuery {
	ret := _m.Called(model)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(interface{}) InterfaceCreateTableQuery); ok {
		r0 = rf(model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// ModelTableExpr provides a mock function with given fields: query, args
func (_m *MockInterfaceCreateTableQuery) ModelTableExpr(query string, args ...interface{}) InterfaceCreateTableQuery {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(string, ...interface{}) InterfaceCreateTableQuery); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// Operation provides a mock function with given fields:
func (_m *MockInterfaceCreateTableQuery) Operation() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Table provides a mock function with given fields: tables
func (_m *MockInterfaceCreateTableQuery) Table(tables ...string) InterfaceCreateTableQuery {
	_va := make([]interface{}, len(tables))
	for _i := range tables {
		_va[_i] = tables[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(...string) InterfaceCreateTableQuery); ok {
		r0 = rf(tables...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// TableExpr provides a mock function with given fields: query, args
func (_m *MockInterfaceCreateTableQuery) TableExpr(query string, args ...interface{}) InterfaceCreateTableQuery {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(string, ...interface{}) InterfaceCreateTableQuery); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// Temp provides a mock function with given fields:
func (_m *MockInterfaceCreateTableQuery) Temp() InterfaceCreateTableQuery {
	ret := _m.Called()

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func() InterfaceCreateTableQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

// Varchar provides a mock function with given fields: n
func (_m *MockInterfaceCreateTableQuery) Varchar(n int) InterfaceCreateTableQuery {
	ret := _m.Called(n)

	var r0 InterfaceCreateTableQuery
	if rf, ok := ret.Get(0).(func(int) InterfaceCreateTableQuery); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateTableQuery)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockInterfaceCreateTableQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfaceCreateTableQuery creates a new instance of MockInterfaceCreateTableQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfaceCreateTableQuery(t mockConstructorTestingTNewMockInterfaceCreateTableQuery) *MockInterfaceCreateTableQuery {
	mock := &MockInterfaceCreateTableQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}