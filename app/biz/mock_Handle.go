// Code generated by mockery v2.42.2. DO NOT EDIT.

package biz

import mock "github.com/stretchr/testify/mock"

// MockHandle is an autogenerated mock type for the Handle type
type MockHandle struct {
	mock.Mock
}

// NewMockHandle creates a new instance of MockHandle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandle(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandle {
	mock := &MockHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
