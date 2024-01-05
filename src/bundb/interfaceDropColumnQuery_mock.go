// Code generated by mockery. DO NOT EDIT.

package bundb

import (
	context "context"

	bun "github.com/uptrace/bun"

	mock "github.com/stretchr/testify/mock"

	schema "github.com/uptrace/bun/schema"

	sql "database/sql"
)

// MockInterfaceDropColumnQuery is an autogenerated mock type for the InterfaceDropColumnQuery type
type MockInterfaceDropColumnQuery struct {
	mock.Mock
}

// AppendQuery provides a mock function with given fields: fmter, b
func (_m *MockInterfaceDropColumnQuery) AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error) {
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

// Column provides a mock function with given fields: columns
func (_m *MockInterfaceDropColumnQuery) Column(columns ...string) InterfaceDropColumnQuery {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func(...string) InterfaceDropColumnQuery); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

// ColumnExpr provides a mock function with given fields: query, args
func (_m *MockInterfaceDropColumnQuery) ColumnExpr(query string, args ...interface{}) InterfaceDropColumnQuery {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func(string, ...interface{}) InterfaceDropColumnQuery); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

// Conn provides a mock function with given fields: db
func (_m *MockInterfaceDropColumnQuery) Conn(db bun.IConn) InterfaceDropColumnQuery {
	ret := _m.Called(db)

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func(bun.IConn) InterfaceDropColumnQuery); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

// Exec provides a mock function with given fields: ctx, dest
func (_m *MockInterfaceDropColumnQuery) Exec(ctx context.Context, dest ...interface{}) (sql.Result, error) {
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

// Model provides a mock function with given fields: model
func (_m *MockInterfaceDropColumnQuery) Model(model interface{}) InterfaceDropColumnQuery {
	ret := _m.Called(model)

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func(interface{}) InterfaceDropColumnQuery); ok {
		r0 = rf(model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

// ModelTableExpr provides a mock function with given fields: query, args
func (_m *MockInterfaceDropColumnQuery) ModelTableExpr(query string, args ...interface{}) InterfaceDropColumnQuery {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func(string, ...interface{}) InterfaceDropColumnQuery); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

// Operation provides a mock function with given fields:
func (_m *MockInterfaceDropColumnQuery) Operation() string {
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
func (_m *MockInterfaceDropColumnQuery) Table(tables ...string) InterfaceDropColumnQuery {
	_va := make([]interface{}, len(tables))
	for _i := range tables {
		_va[_i] = tables[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func(...string) InterfaceDropColumnQuery); ok {
		r0 = rf(tables...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

// TableExpr provides a mock function with given fields: query, args
func (_m *MockInterfaceDropColumnQuery) TableExpr(query string, args ...interface{}) InterfaceDropColumnQuery {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func(string, ...interface{}) InterfaceDropColumnQuery); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockInterfaceDropColumnQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfaceDropColumnQuery creates a new instance of MockInterfaceDropColumnQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfaceDropColumnQuery(t mockConstructorTestingTNewMockInterfaceDropColumnQuery) *MockInterfaceDropColumnQuery {
	mock := &MockInterfaceDropColumnQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
