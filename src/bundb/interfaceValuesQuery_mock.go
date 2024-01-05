// Code generated by mockery. DO NOT EDIT.

package bundb

import (
	mock "github.com/stretchr/testify/mock"
	bun "github.com/uptrace/bun"

	schema "github.com/uptrace/bun/schema"
)

// MockInterfaceValuesQuery is an autogenerated mock type for the InterfaceValuesQuery type
type MockInterfaceValuesQuery struct {
	mock.Mock
}

// AppendColumns provides a mock function with given fields: fmter, b
func (_m *MockInterfaceValuesQuery) AppendColumns(fmter schema.Formatter, b []byte) ([]byte, error) {
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

// AppendNamedArg provides a mock function with given fields: fmter, b, name
func (_m *MockInterfaceValuesQuery) AppendNamedArg(fmter schema.Formatter, b []byte, name string) ([]byte, bool) {
	ret := _m.Called(fmter, b, name)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(schema.Formatter, []byte, string) []byte); ok {
		r0 = rf(fmter, b, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(schema.Formatter, []byte, string) bool); ok {
		r1 = rf(fmter, b, name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// AppendQuery provides a mock function with given fields: fmter, b
func (_m *MockInterfaceValuesQuery) AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error) {
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

// Conn provides a mock function with given fields: db
func (_m *MockInterfaceValuesQuery) Conn(db bun.IConn) InterfaceValuesQuery {
	ret := _m.Called(db)

	var r0 InterfaceValuesQuery
	if rf, ok := ret.Get(0).(func(bun.IConn) InterfaceValuesQuery); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceValuesQuery)
		}
	}

	return r0
}

// Operation provides a mock function with given fields:
func (_m *MockInterfaceValuesQuery) Operation() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Value provides a mock function with given fields: column, expr, args
func (_m *MockInterfaceValuesQuery) Value(column string, expr string, args ...interface{}) InterfaceValuesQuery {
	var _ca []interface{}
	_ca = append(_ca, column, expr)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 InterfaceValuesQuery
	if rf, ok := ret.Get(0).(func(string, string, ...interface{}) InterfaceValuesQuery); ok {
		r0 = rf(column, expr, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceValuesQuery)
		}
	}

	return r0
}

// WithOrder provides a mock function with given fields:
func (_m *MockInterfaceValuesQuery) WithOrder() InterfaceValuesQuery {
	ret := _m.Called()

	var r0 InterfaceValuesQuery
	if rf, ok := ret.Get(0).(func() InterfaceValuesQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceValuesQuery)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockInterfaceValuesQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfaceValuesQuery creates a new instance of MockInterfaceValuesQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfaceValuesQuery(t mockConstructorTestingTNewMockInterfaceValuesQuery) *MockInterfaceValuesQuery {
	mock := &MockInterfaceValuesQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}