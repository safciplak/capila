// Code generated by mockery. DO NOT EDIT.

package handlers

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// mockRequest is an autogenerated mock type for the Request type
type mockRequest struct {
	mock.Mock
}

// Validate provides a mock function with given fields: ctx
func (_m *mockRequest) Validate(ctx *gin.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gin.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewMockRequest interface {
	mock.TestingT
	Cleanup(func())
}

// newMockRequest creates a new instance of mockRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockRequest(t mockConstructorTestingTnewMockRequest) *mockRequest {
	mock := &mockRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
