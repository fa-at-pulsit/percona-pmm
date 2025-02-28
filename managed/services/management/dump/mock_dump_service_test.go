// Code generated by mockery v2.39.1. DO NOT EDIT.

package dump

import (
	mock "github.com/stretchr/testify/mock"

	servicesdump "github.com/percona/pmm/managed/services/dump"
)

// mockDumpService is an autogenerated mock type for the dumpService type
type mockDumpService struct {
	mock.Mock
}

// DeleteDump provides a mock function with given fields: dumpID
func (_m *mockDumpService) DeleteDump(dumpID string) error {
	ret := _m.Called(dumpID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDump")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dumpID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFilePathsForDumps provides a mock function with given fields: dumpIDs
func (_m *mockDumpService) GetFilePathsForDumps(dumpIDs []string) (map[string]string, error) {
	ret := _m.Called(dumpIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetFilePathsForDumps")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) (map[string]string, error)); ok {
		return rf(dumpIDs)
	}
	if rf, ok := ret.Get(0).(func([]string) map[string]string); ok {
		r0 = rf(dumpIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(dumpIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartDump provides a mock function with given fields: params
func (_m *mockDumpService) StartDump(params *servicesdump.Params) (string, error) {
	ret := _m.Called(params)

	if len(ret) == 0 {
		panic("no return value specified for StartDump")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*servicesdump.Params) (string, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(*servicesdump.Params) string); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*servicesdump.Params) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockDumpService creates a new instance of mockDumpService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockDumpService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockDumpService {
	mock := &mockDumpService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
