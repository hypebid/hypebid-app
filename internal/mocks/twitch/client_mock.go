// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	twitch "github.com/hypebid/hypebid-app/internal/twitch"
	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// GetAccessToken provides a mock function with given fields:
func (_m *MockClient) GetAccessToken() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccessToken'
type MockClient_GetAccessToken_Call struct {
	*mock.Call
}

// GetAccessToken is a helper method to define mock.On call
func (_e *MockClient_Expecter) GetAccessToken() *MockClient_GetAccessToken_Call {
	return &MockClient_GetAccessToken_Call{Call: _e.mock.On("GetAccessToken")}
}

func (_c *MockClient_GetAccessToken_Call) Run(run func()) *MockClient_GetAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_GetAccessToken_Call) Return(_a0 string, _a1 error) *MockClient_GetAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetAccessToken_Call) RunAndReturn(run func() (string, error)) *MockClient_GetAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetFollowerCount provides a mock function with given fields: accessToken, broadcasterID
func (_m *MockClient) GetFollowerCount(accessToken string, broadcasterID string) (int, error) {
	ret := _m.Called(accessToken, broadcasterID)

	if len(ret) == 0 {
		panic("no return value specified for GetFollowerCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int, error)); ok {
		return rf(accessToken, broadcasterID)
	}
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(accessToken, broadcasterID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accessToken, broadcasterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetFollowerCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFollowerCount'
type MockClient_GetFollowerCount_Call struct {
	*mock.Call
}

// GetFollowerCount is a helper method to define mock.On call
//   - accessToken string
//   - broadcasterID string
func (_e *MockClient_Expecter) GetFollowerCount(accessToken interface{}, broadcasterID interface{}) *MockClient_GetFollowerCount_Call {
	return &MockClient_GetFollowerCount_Call{Call: _e.mock.On("GetFollowerCount", accessToken, broadcasterID)}
}

func (_c *MockClient_GetFollowerCount_Call) Run(run func(accessToken string, broadcasterID string)) *MockClient_GetFollowerCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockClient_GetFollowerCount_Call) Return(_a0 int, _a1 error) *MockClient_GetFollowerCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetFollowerCount_Call) RunAndReturn(run func(string, string) (int, error)) *MockClient_GetFollowerCount_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByLogin provides a mock function with given fields: accessToken, login
func (_m *MockClient) GetUserByLogin(accessToken string, login string) (string, error) {
	ret := _m.Called(accessToken, login)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByLogin")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(accessToken, login)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(accessToken, login)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accessToken, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetUserByLogin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByLogin'
type MockClient_GetUserByLogin_Call struct {
	*mock.Call
}

// GetUserByLogin is a helper method to define mock.On call
//   - accessToken string
//   - login string
func (_e *MockClient_Expecter) GetUserByLogin(accessToken interface{}, login interface{}) *MockClient_GetUserByLogin_Call {
	return &MockClient_GetUserByLogin_Call{Call: _e.mock.On("GetUserByLogin", accessToken, login)}
}

func (_c *MockClient_GetUserByLogin_Call) Run(run func(accessToken string, login string)) *MockClient_GetUserByLogin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockClient_GetUserByLogin_Call) Return(_a0 string, _a1 error) *MockClient_GetUserByLogin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetUserByLogin_Call) RunAndReturn(run func(string, string) (string, error)) *MockClient_GetUserByLogin_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsersByLogin provides a mock function with given fields: accessToken, logins
func (_m *MockClient) GetUsersByLogin(accessToken string, logins []string) ([]twitch.TwitchUser, error) {
	ret := _m.Called(accessToken, logins)

	if len(ret) == 0 {
		panic("no return value specified for GetUsersByLogin")
	}

	var r0 []twitch.TwitchUser
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) ([]twitch.TwitchUser, error)); ok {
		return rf(accessToken, logins)
	}
	if rf, ok := ret.Get(0).(func(string, []string) []twitch.TwitchUser); ok {
		r0 = rf(accessToken, logins)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]twitch.TwitchUser)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(accessToken, logins)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_GetUsersByLogin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsersByLogin'
type MockClient_GetUsersByLogin_Call struct {
	*mock.Call
}

// GetUsersByLogin is a helper method to define mock.On call
//   - accessToken string
//   - logins []string
func (_e *MockClient_Expecter) GetUsersByLogin(accessToken interface{}, logins interface{}) *MockClient_GetUsersByLogin_Call {
	return &MockClient_GetUsersByLogin_Call{Call: _e.mock.On("GetUsersByLogin", accessToken, logins)}
}

func (_c *MockClient_GetUsersByLogin_Call) Run(run func(accessToken string, logins []string)) *MockClient_GetUsersByLogin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *MockClient_GetUsersByLogin_Call) Return(_a0 []twitch.TwitchUser, _a1 error) *MockClient_GetUsersByLogin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_GetUsersByLogin_Call) RunAndReturn(run func(string, []string) ([]twitch.TwitchUser, error)) *MockClient_GetUsersByLogin_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
