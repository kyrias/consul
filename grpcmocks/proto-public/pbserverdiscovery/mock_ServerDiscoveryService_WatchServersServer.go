// Code generated by mockery v2.41.0. DO NOT EDIT.

package mockpbserverdiscovery

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	pbserverdiscovery "github.com/hashicorp/consul/proto-public/pbserverdiscovery"
)

// ServerDiscoveryService_WatchServersServer is an autogenerated mock type for the ServerDiscoveryService_WatchServersServer type
type ServerDiscoveryService_WatchServersServer struct {
	mock.Mock
}

type ServerDiscoveryService_WatchServersServer_Expecter struct {
	mock *mock.Mock
}

func (_m *ServerDiscoveryService_WatchServersServer) EXPECT() *ServerDiscoveryService_WatchServersServer_Expecter {
	return &ServerDiscoveryService_WatchServersServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *ServerDiscoveryService_WatchServersServer) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// ServerDiscoveryService_WatchServersServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type ServerDiscoveryService_WatchServersServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *ServerDiscoveryService_WatchServersServer_Expecter) Context() *ServerDiscoveryService_WatchServersServer_Context_Call {
	return &ServerDiscoveryService_WatchServersServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *ServerDiscoveryService_WatchServersServer_Context_Call) Run(run func()) *ServerDiscoveryService_WatchServersServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_Context_Call) Return(_a0 context.Context) *ServerDiscoveryService_WatchServersServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_Context_Call) RunAndReturn(run func() context.Context) *ServerDiscoveryService_WatchServersServer_Context_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *ServerDiscoveryService_WatchServersServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerDiscoveryService_WatchServersServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type ServerDiscoveryService_WatchServersServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ServerDiscoveryService_WatchServersServer_Expecter) RecvMsg(m interface{}) *ServerDiscoveryService_WatchServersServer_RecvMsg_Call {
	return &ServerDiscoveryService_WatchServersServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *ServerDiscoveryService_WatchServersServer_RecvMsg_Call) Run(run func(m interface{})) *ServerDiscoveryService_WatchServersServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_RecvMsg_Call) Return(_a0 error) *ServerDiscoveryService_WatchServersServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *ServerDiscoveryService_WatchServersServer_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *ServerDiscoveryService_WatchServersServer) Send(_a0 *pbserverdiscovery.WatchServersResponse) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*pbserverdiscovery.WatchServersResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerDiscoveryService_WatchServersServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type ServerDiscoveryService_WatchServersServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *pbserverdiscovery.WatchServersResponse
func (_e *ServerDiscoveryService_WatchServersServer_Expecter) Send(_a0 interface{}) *ServerDiscoveryService_WatchServersServer_Send_Call {
	return &ServerDiscoveryService_WatchServersServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *ServerDiscoveryService_WatchServersServer_Send_Call) Run(run func(_a0 *pbserverdiscovery.WatchServersResponse)) *ServerDiscoveryService_WatchServersServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*pbserverdiscovery.WatchServersResponse))
	})
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_Send_Call) Return(_a0 error) *ServerDiscoveryService_WatchServersServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_Send_Call) RunAndReturn(run func(*pbserverdiscovery.WatchServersResponse) error) *ServerDiscoveryService_WatchServersServer_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *ServerDiscoveryService_WatchServersServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SendHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerDiscoveryService_WatchServersServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type ServerDiscoveryService_WatchServersServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *ServerDiscoveryService_WatchServersServer_Expecter) SendHeader(_a0 interface{}) *ServerDiscoveryService_WatchServersServer_SendHeader_Call {
	return &ServerDiscoveryService_WatchServersServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *ServerDiscoveryService_WatchServersServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *ServerDiscoveryService_WatchServersServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SendHeader_Call) Return(_a0 error) *ServerDiscoveryService_WatchServersServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SendHeader_Call) RunAndReturn(run func(metadata.MD) error) *ServerDiscoveryService_WatchServersServer_SendHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *ServerDiscoveryService_WatchServersServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerDiscoveryService_WatchServersServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type ServerDiscoveryService_WatchServersServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ServerDiscoveryService_WatchServersServer_Expecter) SendMsg(m interface{}) *ServerDiscoveryService_WatchServersServer_SendMsg_Call {
	return &ServerDiscoveryService_WatchServersServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *ServerDiscoveryService_WatchServersServer_SendMsg_Call) Run(run func(m interface{})) *ServerDiscoveryService_WatchServersServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SendMsg_Call) Return(_a0 error) *ServerDiscoveryService_WatchServersServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SendMsg_Call) RunAndReturn(run func(interface{}) error) *ServerDiscoveryService_WatchServersServer_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *ServerDiscoveryService_WatchServersServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerDiscoveryService_WatchServersServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type ServerDiscoveryService_WatchServersServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *ServerDiscoveryService_WatchServersServer_Expecter) SetHeader(_a0 interface{}) *ServerDiscoveryService_WatchServersServer_SetHeader_Call {
	return &ServerDiscoveryService_WatchServersServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *ServerDiscoveryService_WatchServersServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *ServerDiscoveryService_WatchServersServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SetHeader_Call) Return(_a0 error) *ServerDiscoveryService_WatchServersServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SetHeader_Call) RunAndReturn(run func(metadata.MD) error) *ServerDiscoveryService_WatchServersServer_SetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *ServerDiscoveryService_WatchServersServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// ServerDiscoveryService_WatchServersServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type ServerDiscoveryService_WatchServersServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *ServerDiscoveryService_WatchServersServer_Expecter) SetTrailer(_a0 interface{}) *ServerDiscoveryService_WatchServersServer_SetTrailer_Call {
	return &ServerDiscoveryService_WatchServersServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *ServerDiscoveryService_WatchServersServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *ServerDiscoveryService_WatchServersServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SetTrailer_Call) Return() *ServerDiscoveryService_WatchServersServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServerDiscoveryService_WatchServersServer_SetTrailer_Call) RunAndReturn(run func(metadata.MD)) *ServerDiscoveryService_WatchServersServer_SetTrailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewServerDiscoveryService_WatchServersServer creates a new instance of ServerDiscoveryService_WatchServersServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServerDiscoveryService_WatchServersServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServerDiscoveryService_WatchServersServer {
	mock := &ServerDiscoveryService_WatchServersServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
