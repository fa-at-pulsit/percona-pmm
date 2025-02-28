// Code generated by mockery v2.39.1. DO NOT EDIT.

package server

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockAlertmanagerService is an autogenerated mock type for the alertmanagerService type
type mockAlertmanagerService struct {
	mock.Mock
}

// IsReady provides a mock function with given fields: ctx
func (_m *mockAlertmanagerService) IsReady(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsReady")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestConfigurationUpdate provides a mock function with given fields:
func (_m *mockAlertmanagerService) RequestConfigurationUpdate() {
	_m.Called()
}

// newMockAlertmanagerService creates a new instance of mockAlertmanagerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAlertmanagerService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockAlertmanagerService {
	mock := &mockAlertmanagerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
