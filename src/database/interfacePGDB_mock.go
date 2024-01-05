// Code generated by mockery. DO NOT EDIT.

package database

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	orm "github.com/go-pg/pg/v10/orm"

	pg "github.com/go-pg/pg/v10"

	time "time"
)

// MockInterfacePGDB is an autogenerated mock type for the InterfacePGDB type
type MockInterfacePGDB struct {
	mock.Mock
}

// AddQueryHook provides a mock function with given fields: hook
func (_m *MockInterfacePGDB) AddQueryHook(hook pg.QueryHook) {
	_m.Called(hook)
}

// Begin provides a mock function with given fields:
func (_m *MockInterfacePGDB) Begin() (InterfacePgTx, error) {
	ret := _m.Called()

	var r0 InterfacePgTx
	if rf, ok := ret.Get(0).(func() InterfacePgTx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfacePgTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginContext provides a mock function with given fields: ctx
func (_m *MockInterfacePGDB) BeginContext(ctx context.Context) (InterfacePgTx, error) {
	ret := _m.Called(ctx)

	var r0 InterfacePgTx
	if rf, ok := ret.Get(0).(func(context.Context) InterfacePgTx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfacePgTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *MockInterfacePGDB) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *MockInterfacePGDB) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// CopyFrom provides a mock function with given fields: r, query, params
func (_m *MockInterfacePGDB) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, r, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(io.Reader, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(r, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader, interface{}, ...interface{}) error); ok {
		r1 = rf(r, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyTo provides a mock function with given fields: w, query, params
func (_m *MockInterfacePGDB) CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, w, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(io.Writer, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(w, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Writer, interface{}, ...interface{}) error); ok {
		r1 = rf(w, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exec provides a mock function with given fields: query, params
func (_m *MockInterfacePGDB) Exec(query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecContext provides a mock function with given fields: c, query, params
func (_m *MockInterfacePGDB) ExecContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, c, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(c, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(c, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecOne provides a mock function with given fields: query, params
func (_m *MockInterfacePGDB) ExecOne(query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecOneContext provides a mock function with given fields: ctx, query, params
func (_m *MockInterfacePGDB) ExecOneContext(ctx context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(ctx, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(ctx, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Formatter provides a mock function with given fields:
func (_m *MockInterfacePGDB) Formatter() orm.QueryFormatter {
	ret := _m.Called()

	var r0 orm.QueryFormatter
	if rf, ok := ret.Get(0).(func() orm.QueryFormatter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QueryFormatter)
		}
	}

	return r0
}

// Model provides a mock function with given fields: model
func (_m *MockInterfacePGDB) Model(model ...interface{}) InterfaceORMQuery {
	var _ca []interface{}
	_ca = append(_ca, model...)
	ret := _m.Called(_ca...)

	var r0 InterfaceORMQuery
	if rf, ok := ret.Get(0).(func(...interface{}) InterfaceORMQuery); ok {
		r0 = rf(model...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceORMQuery)
		}
	}

	return r0
}

// ModelContext provides a mock function with given fields: c, model
func (_m *MockInterfacePGDB) ModelContext(c context.Context, model ...interface{}) InterfaceORMQuery {
	var _ca []interface{}
	_ca = append(_ca, c)
	_ca = append(_ca, model...)
	ret := _m.Called(_ca...)

	var r0 InterfaceORMQuery
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) InterfaceORMQuery); ok {
		r0 = rf(c, model...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceORMQuery)
		}
	}

	return r0
}

// Param provides a mock function with given fields: param
func (_m *MockInterfacePGDB) Param(param string) interface{} {
	ret := _m.Called(param)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Ping provides a mock function with given fields: ctx
func (_m *MockInterfacePGDB) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PoolStats provides a mock function with given fields:
func (_m *MockInterfacePGDB) PoolStats() *pg.PoolStats {
	ret := _m.Called()

	var r0 *pg.PoolStats
	if rf, ok := ret.Get(0).(func() *pg.PoolStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.PoolStats)
		}
	}

	return r0
}

// Prepare provides a mock function with given fields: q
func (_m *MockInterfacePGDB) Prepare(q string) (InterfacePgStmt, error) {
	ret := _m.Called(q)

	var r0 InterfacePgStmt
	if rf, ok := ret.Get(0).(func(string) InterfacePgStmt); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfacePgStmt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: model, query, params
func (_m *MockInterfacePGDB) Query(model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryContext provides a mock function with given fields: c, model, query, params
func (_m *MockInterfacePGDB) QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, c, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(c, model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(c, model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryOne provides a mock function with given fields: model, query, params
func (_m *MockInterfacePGDB) QueryOne(model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryOneContext provides a mock function with given fields: ctx, model, query, params
func (_m *MockInterfacePGDB) QueryOneContext(ctx context.Context, model interface{}, query interface{}, params ...interface{}) (orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Result
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...interface{}) orm.Result); ok {
		r0 = rf(ctx, model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(ctx, model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunInTransaction provides a mock function with given fields: ctx, fn
func (_m *MockInterfacePGDB) RunInTransaction(ctx context.Context, fn func(InterfacePgTx) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(InterfacePgTx) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *MockInterfacePGDB) WithContext(ctx context.Context) InterfacePGDB {
	ret := _m.Called(ctx)

	var r0 InterfacePGDB
	if rf, ok := ret.Get(0).(func(context.Context) InterfacePGDB); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfacePGDB)
		}
	}

	return r0
}

// WithParam provides a mock function with given fields: param, value
func (_m *MockInterfacePGDB) WithParam(param string, value interface{}) InterfacePGDB {
	ret := _m.Called(param, value)

	var r0 InterfacePGDB
	if rf, ok := ret.Get(0).(func(string, interface{}) InterfacePGDB); ok {
		r0 = rf(param, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfacePGDB)
		}
	}

	return r0
}

// WithTimeout provides a mock function with given fields: d
func (_m *MockInterfacePGDB) WithTimeout(d time.Duration) InterfacePGDB {
	ret := _m.Called(d)

	var r0 InterfacePGDB
	if rf, ok := ret.Get(0).(func(time.Duration) InterfacePGDB); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfacePGDB)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockInterfacePGDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfacePGDB creates a new instance of MockInterfacePGDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfacePGDB(t mockConstructorTestingTNewMockInterfacePGDB) *MockInterfacePGDB {
	mock := &MockInterfacePGDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}