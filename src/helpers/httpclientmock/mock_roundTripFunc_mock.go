// Code generated by mockery. DO NOT EDIT.

package httpclientmock

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// mockRoundTripFunc is an autogenerated mock type for the roundTripFunc type
type mockRoundTripFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: req
func (_m *mockRoundTripFunc) Execute(req *http.Request) *http.Response {
	ret := _m.Called(req)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	return r0
}

type mockConstructorTestingTnewMockRoundTripFunc interface {
	mock.TestingT
	Cleanup(func())
}

// newMockRoundTripFunc creates a new instance of mockRoundTripFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockRoundTripFunc(t mockConstructorTestingTnewMockRoundTripFunc) *mockRoundTripFunc {
	mock := &mockRoundTripFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
