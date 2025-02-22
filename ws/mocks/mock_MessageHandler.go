// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import (
	ws "github.com/lorenzodonini/ocpp-go/ws"
	mock "github.com/stretchr/testify/mock"
)

// MockMessageHandler is an autogenerated mock type for the MessageHandler type
type MockMessageHandler struct {
	mock.Mock
}

type MockMessageHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMessageHandler) EXPECT() *MockMessageHandler_Expecter {
	return &MockMessageHandler_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: c, data
func (_m *MockMessageHandler) Execute(c ws.Channel, data []byte) error {
	ret := _m.Called(c, data)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(ws.Channel, []byte) error); ok {
		r0 = rf(c, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockMessageHandler_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockMessageHandler_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - c ws.Channel
//   - data []byte
func (_e *MockMessageHandler_Expecter) Execute(c interface{}, data interface{}) *MockMessageHandler_Execute_Call {
	return &MockMessageHandler_Execute_Call{Call: _e.mock.On("Execute", c, data)}
}

func (_c *MockMessageHandler_Execute_Call) Run(run func(c ws.Channel, data []byte)) *MockMessageHandler_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ws.Channel), args[1].([]byte))
	})
	return _c
}

func (_c *MockMessageHandler_Execute_Call) Return(_a0 error) *MockMessageHandler_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMessageHandler_Execute_Call) RunAndReturn(run func(ws.Channel, []byte) error) *MockMessageHandler_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMessageHandler creates a new instance of MockMessageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessageHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessageHandler {
	mock := &MockMessageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
