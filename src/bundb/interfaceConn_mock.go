// Code generated by mockery. DO NOT EDIT.

package bundb

import (
	context "context"
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"
)

// MockInterfaceConn is an autogenerated mock type for the InterfaceConn type
type MockInterfaceConn struct {
	mock.Mock
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *MockInterfaceConn) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAddColumn provides a mock function with given fields:
func (_m *MockInterfaceConn) NewAddColumn() InterfaceAddColumnQuery {
	ret := _m.Called()

	var r0 InterfaceAddColumnQuery
	if rf, ok := ret.Get(0).(func() InterfaceAddColumnQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceAddColumnQuery)
		}
	}

	return r0
}

// NewCreateIndex provides a mock function with given fields:
func (_m *MockInterfaceConn) NewCreateIndex() InterfaceCreateIndexQuery {
	ret := _m.Called()

	var r0 InterfaceCreateIndexQuery
	if rf, ok := ret.Get(0).(func() InterfaceCreateIndexQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceCreateIndexQuery)
		}
	}

	return r0
}

// NewCreateTable provides a mock function with given fields:
func (_m *MockInterfaceConn) NewCreateTable() InterfaceCreateTableQuery {
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

// NewDelete provides a mock function with given fields:
func (_m *MockInterfaceConn) NewDelete() InterfaceDeleteQuery {
	ret := _m.Called()

	var r0 InterfaceDeleteQuery
	if rf, ok := ret.Get(0).(func() InterfaceDeleteQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDeleteQuery)
		}
	}

	return r0
}

// NewDropColumn provides a mock function with given fields:
func (_m *MockInterfaceConn) NewDropColumn() InterfaceDropColumnQuery {
	ret := _m.Called()

	var r0 InterfaceDropColumnQuery
	if rf, ok := ret.Get(0).(func() InterfaceDropColumnQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropColumnQuery)
		}
	}

	return r0
}

// NewDropIndex provides a mock function with given fields:
func (_m *MockInterfaceConn) NewDropIndex() InterfaceDropIndexQuery {
	ret := _m.Called()

	var r0 InterfaceDropIndexQuery
	if rf, ok := ret.Get(0).(func() InterfaceDropIndexQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropIndexQuery)
		}
	}

	return r0
}

// NewDropTable provides a mock function with given fields:
func (_m *MockInterfaceConn) NewDropTable() InterfaceDropTableQuery {
	ret := _m.Called()

	var r0 InterfaceDropTableQuery
	if rf, ok := ret.Get(0).(func() InterfaceDropTableQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceDropTableQuery)
		}
	}

	return r0
}

// NewInsert provides a mock function with given fields:
func (_m *MockInterfaceConn) NewInsert() InterfaceInsertQuery {
	ret := _m.Called()

	var r0 InterfaceInsertQuery
	if rf, ok := ret.Get(0).(func() InterfaceInsertQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceInsertQuery)
		}
	}

	return r0
}

// NewSelect provides a mock function with given fields:
func (_m *MockInterfaceConn) NewSelect() InterfaceSelectQuery {
	ret := _m.Called()

	var r0 InterfaceSelectQuery
	if rf, ok := ret.Get(0).(func() InterfaceSelectQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceSelectQuery)
		}
	}

	return r0
}

// NewTruncateTable provides a mock function with given fields:
func (_m *MockInterfaceConn) NewTruncateTable() InterfaceTruncateTableQuery {
	ret := _m.Called()

	var r0 InterfaceTruncateTableQuery
	if rf, ok := ret.Get(0).(func() InterfaceTruncateTableQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceTruncateTableQuery)
		}
	}

	return r0
}

// NewUpdate provides a mock function with given fields:
func (_m *MockInterfaceConn) NewUpdate() InterfaceUpdateQuery {
	ret := _m.Called()

	var r0 InterfaceUpdateQuery
	if rf, ok := ret.Get(0).(func() InterfaceUpdateQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceUpdateQuery)
		}
	}

	return r0
}

// NewValues provides a mock function with given fields: model
func (_m *MockInterfaceConn) NewValues(model interface{}) InterfaceValuesQuery {
	ret := _m.Called(model)

	var r0 InterfaceValuesQuery
	if rf, ok := ret.Get(0).(func(interface{}) InterfaceValuesQuery); ok {
		r0 = rf(model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceValuesQuery)
		}
	}

	return r0
}

// QueryContext provides a mock function with given fields: ctx, query, args
func (_m *MockInterfaceConn) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Rows
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryRowContext provides a mock function with given fields: ctx, query, args
func (_m *MockInterfaceConn) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockInterfaceConn interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfaceConn creates a new instance of MockInterfaceConn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfaceConn(t mockConstructorTestingTNewMockInterfaceConn) *MockInterfaceConn {
	mock := &MockInterfaceConn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
