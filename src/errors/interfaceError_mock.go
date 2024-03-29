// Code generated by mockery. DO NOT EDIT.

package errors

import mock "github.com/stretchr/testify/mock"

// MockInterfaceError is an autogenerated mock type for the InterfaceError type
type MockInterfaceError struct {
	mock.Mock
}

// Error provides a mock function with given fields:
func (_m *MockInterfaceError) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetCode provides a mock function with given fields:
func (_m *MockInterfaceError) GetCode() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDetail provides a mock function with given fields:
func (_m *MockInterfaceError) GetDetail() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStatusCode provides a mock function with given fields:
func (_m *MockInterfaceError) GetStatusCode() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Unwrap provides a mock function with given fields:
func (_m *MockInterfaceError) Unwrap() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Wrap provides a mock function with given fields: err
func (_m *MockInterfaceError) Wrap(err error) InterfaceError {
	ret := _m.Called(err)

	var r0 InterfaceError
	if rf, ok := ret.Get(0).(func(error) InterfaceError); ok {
		r0 = rf(err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(InterfaceError)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockInterfaceError interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfaceError creates a new instance of MockInterfaceError. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfaceError(t mockConstructorTestingTNewMockInterfaceError) *MockInterfaceError {
	mock := &MockInterfaceError{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
