// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/hypebid/hypebid-app/pkg/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockShareHolderService is an autogenerated mock type for the ShareHolderService type
type MockShareHolderService struct {
	mock.Mock
}

type MockShareHolderService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockShareHolderService) EXPECT() *MockShareHolderService_Expecter {
	return &MockShareHolderService_Expecter{mock: &_m.Mock}
}

// CreateShareHolderForChannel provides a mock function with given fields: userID, marketChannelID
func (_m *MockShareHolderService) CreateShareHolderForChannel(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
	ret := _m.Called(userID, marketChannelID)

	if len(ret) == 0 {
		panic("no return value specified for CreateShareHolderForChannel")
	}

	var r0 *models.ShareHolder
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) (*models.ShareHolder, error)); ok {
		return rf(userID, marketChannelID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) *models.ShareHolder); ok {
		r0 = rf(userID, marketChannelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShareHolder)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(userID, marketChannelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShareHolderService_CreateShareHolderForChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateShareHolderForChannel'
type MockShareHolderService_CreateShareHolderForChannel_Call struct {
	*mock.Call
}

// CreateShareHolderForChannel is a helper method to define mock.On call
//   - userID uuid.UUID
//   - marketChannelID uuid.UUID
func (_e *MockShareHolderService_Expecter) CreateShareHolderForChannel(userID interface{}, marketChannelID interface{}) *MockShareHolderService_CreateShareHolderForChannel_Call {
	return &MockShareHolderService_CreateShareHolderForChannel_Call{Call: _e.mock.On("CreateShareHolderForChannel", userID, marketChannelID)}
}

func (_c *MockShareHolderService_CreateShareHolderForChannel_Call) Run(run func(userID uuid.UUID, marketChannelID uuid.UUID)) *MockShareHolderService_CreateShareHolderForChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockShareHolderService_CreateShareHolderForChannel_Call) Return(_a0 *models.ShareHolder, _a1 error) *MockShareHolderService_CreateShareHolderForChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShareHolderService_CreateShareHolderForChannel_Call) RunAndReturn(run func(uuid.UUID, uuid.UUID) (*models.ShareHolder, error)) *MockShareHolderService_CreateShareHolderForChannel_Call {
	_c.Call.Return(run)
	return _c
}

// GetShareHolderByUserIDAndMarketChannelID provides a mock function with given fields: userID, marketChannelID
func (_m *MockShareHolderService) GetShareHolderByUserIDAndMarketChannelID(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
	ret := _m.Called(userID, marketChannelID)

	if len(ret) == 0 {
		panic("no return value specified for GetShareHolderByUserIDAndMarketChannelID")
	}

	var r0 *models.ShareHolder
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) (*models.ShareHolder, error)); ok {
		return rf(userID, marketChannelID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) *models.ShareHolder); ok {
		r0 = rf(userID, marketChannelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShareHolder)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(userID, marketChannelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShareHolderByUserIDAndMarketChannelID'
type MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call struct {
	*mock.Call
}

// GetShareHolderByUserIDAndMarketChannelID is a helper method to define mock.On call
//   - userID uuid.UUID
//   - marketChannelID uuid.UUID
func (_e *MockShareHolderService_Expecter) GetShareHolderByUserIDAndMarketChannelID(userID interface{}, marketChannelID interface{}) *MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call {
	return &MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call{Call: _e.mock.On("GetShareHolderByUserIDAndMarketChannelID", userID, marketChannelID)}
}

func (_c *MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call) Run(run func(userID uuid.UUID, marketChannelID uuid.UUID)) *MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call) Return(_a0 *models.ShareHolder, _a1 error) *MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call) RunAndReturn(run func(uuid.UUID, uuid.UUID) (*models.ShareHolder, error)) *MockShareHolderService_GetShareHolderByUserIDAndMarketChannelID_Call {
	_c.Call.Return(run)
	return _c
}

// InitializeShareHolder provides a mock function with given fields: userID, marketChannelID
func (_m *MockShareHolderService) InitializeShareHolder(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
	ret := _m.Called(userID, marketChannelID)

	if len(ret) == 0 {
		panic("no return value specified for InitializeShareHolder")
	}

	var r0 *models.ShareHolder
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) (*models.ShareHolder, error)); ok {
		return rf(userID, marketChannelID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) *models.ShareHolder); ok {
		r0 = rf(userID, marketChannelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShareHolder)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(userID, marketChannelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShareHolderService_InitializeShareHolder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitializeShareHolder'
type MockShareHolderService_InitializeShareHolder_Call struct {
	*mock.Call
}

// InitializeShareHolder is a helper method to define mock.On call
//   - userID uuid.UUID
//   - marketChannelID uuid.UUID
func (_e *MockShareHolderService_Expecter) InitializeShareHolder(userID interface{}, marketChannelID interface{}) *MockShareHolderService_InitializeShareHolder_Call {
	return &MockShareHolderService_InitializeShareHolder_Call{Call: _e.mock.On("InitializeShareHolder", userID, marketChannelID)}
}

func (_c *MockShareHolderService_InitializeShareHolder_Call) Run(run func(userID uuid.UUID, marketChannelID uuid.UUID)) *MockShareHolderService_InitializeShareHolder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockShareHolderService_InitializeShareHolder_Call) Return(_a0 *models.ShareHolder, _a1 error) *MockShareHolderService_InitializeShareHolder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShareHolderService_InitializeShareHolder_Call) RunAndReturn(run func(uuid.UUID, uuid.UUID) (*models.ShareHolder, error)) *MockShareHolderService_InitializeShareHolder_Call {
	_c.Call.Return(run)
	return _c
}

// TransferShares provides a mock function with given fields: marketChannelID, fromUserID, toUserID, shareCount
func (_m *MockShareHolderService) TransferShares(marketChannelID uuid.UUID, fromUserID uuid.UUID, toUserID uuid.UUID, shareCount int) error {
	ret := _m.Called(marketChannelID, fromUserID, toUserID, shareCount)

	if len(ret) == 0 {
		panic("no return value specified for TransferShares")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID, uuid.UUID, int) error); ok {
		r0 = rf(marketChannelID, fromUserID, toUserID, shareCount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockShareHolderService_TransferShares_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransferShares'
type MockShareHolderService_TransferShares_Call struct {
	*mock.Call
}

// TransferShares is a helper method to define mock.On call
//   - marketChannelID uuid.UUID
//   - fromUserID uuid.UUID
//   - toUserID uuid.UUID
//   - shareCount int
func (_e *MockShareHolderService_Expecter) TransferShares(marketChannelID interface{}, fromUserID interface{}, toUserID interface{}, shareCount interface{}) *MockShareHolderService_TransferShares_Call {
	return &MockShareHolderService_TransferShares_Call{Call: _e.mock.On("TransferShares", marketChannelID, fromUserID, toUserID, shareCount)}
}

func (_c *MockShareHolderService_TransferShares_Call) Run(run func(marketChannelID uuid.UUID, fromUserID uuid.UUID, toUserID uuid.UUID, shareCount int)) *MockShareHolderService_TransferShares_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(uuid.UUID), args[2].(uuid.UUID), args[3].(int))
	})
	return _c
}

func (_c *MockShareHolderService_TransferShares_Call) Return(_a0 error) *MockShareHolderService_TransferShares_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockShareHolderService_TransferShares_Call) RunAndReturn(run func(uuid.UUID, uuid.UUID, uuid.UUID, int) error) *MockShareHolderService_TransferShares_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateShareHolder provides a mock function with given fields: shareHolder
func (_m *MockShareHolderService) UpdateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error) {
	ret := _m.Called(shareHolder)

	if len(ret) == 0 {
		panic("no return value specified for UpdateShareHolder")
	}

	var r0 *models.ShareHolder
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.ShareHolder) (*models.ShareHolder, error)); ok {
		return rf(shareHolder)
	}
	if rf, ok := ret.Get(0).(func(*models.ShareHolder) *models.ShareHolder); ok {
		r0 = rf(shareHolder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShareHolder)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.ShareHolder) error); ok {
		r1 = rf(shareHolder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockShareHolderService_UpdateShareHolder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateShareHolder'
type MockShareHolderService_UpdateShareHolder_Call struct {
	*mock.Call
}

// UpdateShareHolder is a helper method to define mock.On call
//   - shareHolder *models.ShareHolder
func (_e *MockShareHolderService_Expecter) UpdateShareHolder(shareHolder interface{}) *MockShareHolderService_UpdateShareHolder_Call {
	return &MockShareHolderService_UpdateShareHolder_Call{Call: _e.mock.On("UpdateShareHolder", shareHolder)}
}

func (_c *MockShareHolderService_UpdateShareHolder_Call) Run(run func(shareHolder *models.ShareHolder)) *MockShareHolderService_UpdateShareHolder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.ShareHolder))
	})
	return _c
}

func (_c *MockShareHolderService_UpdateShareHolder_Call) Return(_a0 *models.ShareHolder, _a1 error) *MockShareHolderService_UpdateShareHolder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShareHolderService_UpdateShareHolder_Call) RunAndReturn(run func(*models.ShareHolder) (*models.ShareHolder, error)) *MockShareHolderService_UpdateShareHolder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockShareHolderService creates a new instance of MockShareHolderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockShareHolderService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockShareHolderService {
	mock := &MockShareHolderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
