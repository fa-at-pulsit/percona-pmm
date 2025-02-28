// Code generated by mockery v2.39.1. DO NOT EDIT.

package qan

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	grpc "google.golang.org/grpc"

	qanv1beta1 "github.com/percona/pmm/api/qanpb"
)

// mockQanCollectorClient is an autogenerated mock type for the qanCollectorClient type
type mockQanCollectorClient struct {
	mock.Mock
}

// Collect provides a mock function with given fields: ctx, in, opts
func (_m *mockQanCollectorClient) Collect(ctx context.Context, in *qanv1beta1.CollectRequest, opts ...grpc.CallOption) (*qanv1beta1.CollectResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Collect")
	}

	var r0 *qanv1beta1.CollectResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *qanv1beta1.CollectRequest, ...grpc.CallOption) (*qanv1beta1.CollectResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *qanv1beta1.CollectRequest, ...grpc.CallOption) *qanv1beta1.CollectResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*qanv1beta1.CollectResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *qanv1beta1.CollectRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockQanCollectorClient creates a new instance of mockQanCollectorClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockQanCollectorClient(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockQanCollectorClient {
	mock := &mockQanCollectorClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
