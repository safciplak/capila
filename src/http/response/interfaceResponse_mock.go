// Code generated by mockery. DO NOT EDIT.

package response

import (
	context "context"

	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	presenters "github.com/safciplak/capila/src/http/presenters"
)

// MockInterfaceResponse is an autogenerated mock type for the InterfaceResponse type
type MockInterfaceResponse struct {
	mock.Mock
}

// CheckForErrors provides a mock function with given fields: ctx, validationErrors, err
func (_m *MockInterfaceResponse) CheckForErrors(ctx context.Context, validationErrors map[string]string, err error) {
	_m.Called(ctx, validationErrors, err)
}

// HandleError provides a mock function with given fields: ctx, err
func (_m *MockInterfaceResponse) HandleError(ctx context.Context, err error) *Response {
	ret := _m.Called(ctx, err)

	var r0 *Response
	if rf, ok := ret.Get(0).(func(context.Context, error) *Response); ok {
		r0 = rf(ctx, err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	return r0
}

// HandleValidationError provides a mock function with given fields: ctx, err
func (_m *MockInterfaceResponse) HandleValidationError(ctx context.Context, err error) *Response {
	ret := _m.Called(ctx, err)

	var r0 *Response
	if rf, ok := ret.Get(0).(func(context.Context, error) *Response); ok {
		r0 = rf(ctx, err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Response)
		}
	}

	return r0
}

// HasErrors provides a mock function with given fields:
func (_m *MockInterfaceResponse) HasErrors() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasValidationErrors provides a mock function with given fields:
func (_m *MockInterfaceResponse) HasValidationErrors() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Output provides a mock function with given fields: presenter
func (_m *MockInterfaceResponse) Output(presenter presenters.InterfacePresenter) {
	_m.Called(presenter)
}

// ReturnJSON provides a mock function with given fields: ctx
func (_m *MockInterfaceResponse) ReturnJSON(ctx *gin.Context) {
	_m.Called(ctx)
}

// SetError provides a mock function with given fields: ctx, err
func (_m *MockInterfaceResponse) SetError(ctx context.Context, err error) {
	_m.Called(ctx, err)
}

type mockConstructorTestingTNewMockInterfaceResponse interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInterfaceResponse creates a new instance of MockInterfaceResponse. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInterfaceResponse(t mockConstructorTestingTNewMockInterfaceResponse) *MockInterfaceResponse {
	mock := &MockInterfaceResponse{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
