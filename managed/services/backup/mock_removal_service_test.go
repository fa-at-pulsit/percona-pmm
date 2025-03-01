// Code generated by mockery v2.39.1. DO NOT EDIT.

package backup

import mock "github.com/stretchr/testify/mock"

// mockRemovalService is an autogenerated mock type for the removalService type
type mockRemovalService struct {
	mock.Mock
}

// DeleteArtifact provides a mock function with given fields: storage, artifactID, removeFiles
func (_m *mockRemovalService) DeleteArtifact(storage Storage, artifactID string, removeFiles bool) error {
	ret := _m.Called(storage, artifactID, removeFiles)

	if len(ret) == 0 {
		panic("no return value specified for DeleteArtifact")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Storage, string, bool) error); ok {
		r0 = rf(storage, artifactID, removeFiles)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TrimPITRArtifact provides a mock function with given fields: storage, artifactID, firstN
func (_m *mockRemovalService) TrimPITRArtifact(storage Storage, artifactID string, firstN int) error {
	ret := _m.Called(storage, artifactID, firstN)

	if len(ret) == 0 {
		panic("no return value specified for TrimPITRArtifact")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(Storage, string, int) error); ok {
		r0 = rf(storage, artifactID, firstN)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockRemovalService creates a new instance of mockRemovalService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockRemovalService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockRemovalService {
	mock := &mockRemovalService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
