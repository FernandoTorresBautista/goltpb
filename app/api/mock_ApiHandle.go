// Code generated by mockery v2.42.2. DO NOT EDIT.

package api

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockApiHandle is an autogenerated mock type for the ApiHandle type
type MockApiHandle struct {
	mock.Mock
}

// Run provides a mock function with given fields: ctx
func (_m *MockApiHandle) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TurnOff provides a mock function with given fields:
func (_m *MockApiHandle) TurnOff() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TurnOff")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockApiHandle creates a new instance of MockApiHandle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockApiHandle(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockApiHandle {
	mock := &MockApiHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
