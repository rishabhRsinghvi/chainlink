// Code generated by mockery v2.0.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// RPCClient is an autogenerated mock type for the RPCClient type
type RPCClient struct {
	mock.Mock
}

// CallContext provides a mock function with given fields: ctx, result, method, args
func (_m *RPCClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
