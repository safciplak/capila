// Code generated by mockery. DO NOT EDIT.

package bundb

import (
	context "context"

	bun "github.com/uptrace/bun"

	mock "github.com/stretchr/testify/mock"

	schema "github.com/uptrace/bun/schema"
)

// MockInterfaceRawQuery is an autogenerated mock type for the InterfaceRawQuery type
type MockInterfaceRawQuery struct {
	mock.Mock
}

// AppendQuery provides a mock function with given fields: formatter, b
func (_m *MockInterfaceRawQuery) AppendQuery(formatter schema.Formatter, b []byte) ([]byte, error) {
	ret := _m.Called(formatter, b)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(schema.Formatter, []byte) []byte); ok {
		r0 = rf(formatter, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(schema.Formatter, []byte) error); ok {
		r1 = rf(formatter, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Conn provides a mock function with given fields: db
func (_m *MockInterfaceRawQuery) Conn(db bun.IConn) InterfaceRawQuery {
	ret := _m.Called(db)

	var r0 InterfaceRawQuery
	if rf, ok := ret.Get(0).(func(bun.IConn) InterfaceRawQuery); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceRawQuery)
		}
	}

	return r0
}

// Operation provides a mock function with given fields:
func (_m *MockInterfaceRawQuery) Operation() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Scan provides a mock function with given fields: ctx, dest
func (_m *MockInterfaceRawQuery) Scan(ctx context.Context, dest ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, dest...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) error); ok {
		r0 = rf(ctx, dest...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockInterfaceRawQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfaceRawQuery creates a new instance of MockInterfaceRawQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfaceRawQuery(t mockConstructorTestingTNewMockInterfaceRawQuery) *MockInterfaceRawQuery {
	mock := &MockInterfaceRawQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}