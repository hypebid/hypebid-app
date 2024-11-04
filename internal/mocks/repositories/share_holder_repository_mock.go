// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/hypebid/hypebid-app/pkg/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockShareHolderRepository is an autogenerated mock type for the ShareHolderRepository type
type MockShareHolderRepository struct {
	mock.Mock
}

type MockShareHolderRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockShareHolderRepository) EXPECT() *MockShareHolderRepository_Expecter {
	return &MockShareHolderRepository_Expecter{mock: &_m.Mock}
}

// CreateShareHolder provides a mock function with given fields: shareHolder
func (_m *MockShareHolderRepository) CreateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error) {
	ret := _m.Called(shareHolder)

	if len(ret) == 0 {
		panic("no return value specified for CreateShareHolder")
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

// MockShareHolderRepository_CreateShareHolder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateShareHolder'
type MockShareHolderRepository_CreateShareHolder_Call struct {
	*mock.Call
}

// CreateShareHolder is a helper method to define mock.On call
//   - shareHolder *models.ShareHolder
func (_e *MockShareHolderRepository_Expecter) CreateShareHolder(shareHolder interface{}) *MockShareHolderRepository_CreateShareHolder_Call {
	return &MockShareHolderRepository_CreateShareHolder_Call{Call: _e.mock.On("CreateShareHolder", shareHolder)}
}

func (_c *MockShareHolderRepository_CreateShareHolder_Call) Run(run func(shareHolder *models.ShareHolder)) *MockShareHolderRepository_CreateShareHolder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.ShareHolder))
	})
	return _c
}

func (_c *MockShareHolderRepository_CreateShareHolder_Call) Return(_a0 *models.ShareHolder, _a1 error) *MockShareHolderRepository_CreateShareHolder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShareHolderRepository_CreateShareHolder_Call) RunAndReturn(run func(*models.ShareHolder) (*models.ShareHolder, error)) *MockShareHolderRepository_CreateShareHolder_Call {
	_c.Call.Return(run)
	return _c
}

// GetShareHolderByUserIDAndMarketChannelID provides a mock function with given fields: userID, marketChannelID
func (_m *MockShareHolderRepository) GetShareHolderByUserIDAndMarketChannelID(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
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

// MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetShareHolderByUserIDAndMarketChannelID'
type MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call struct {
	*mock.Call
}

// GetShareHolderByUserIDAndMarketChannelID is a helper method to define mock.On call
//   - userID uuid.UUID
//   - marketChannelID uuid.UUID
func (_e *MockShareHolderRepository_Expecter) GetShareHolderByUserIDAndMarketChannelID(userID interface{}, marketChannelID interface{}) *MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call {
	return &MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call{Call: _e.mock.On("GetShareHolderByUserIDAndMarketChannelID", userID, marketChannelID)}
}

func (_c *MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call) Run(run func(userID uuid.UUID, marketChannelID uuid.UUID)) *MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call) Return(_a0 *models.ShareHolder, _a1 error) *MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call) RunAndReturn(run func(uuid.UUID, uuid.UUID) (*models.ShareHolder, error)) *MockShareHolderRepository_GetShareHolderByUserIDAndMarketChannelID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateShareHolder provides a mock function with given fields: shareHolder
func (_m *MockShareHolderRepository) UpdateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error) {
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

// MockShareHolderRepository_UpdateShareHolder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateShareHolder'
type MockShareHolderRepository_UpdateShareHolder_Call struct {
	*mock.Call
}

// UpdateShareHolder is a helper method to define mock.On call
//   - shareHolder *models.ShareHolder
func (_e *MockShareHolderRepository_Expecter) UpdateShareHolder(shareHolder interface{}) *MockShareHolderRepository_UpdateShareHolder_Call {
	return &MockShareHolderRepository_UpdateShareHolder_Call{Call: _e.mock.On("UpdateShareHolder", shareHolder)}
}

func (_c *MockShareHolderRepository_UpdateShareHolder_Call) Run(run func(shareHolder *models.ShareHolder)) *MockShareHolderRepository_UpdateShareHolder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.ShareHolder))
	})
	return _c
}

func (_c *MockShareHolderRepository_UpdateShareHolder_Call) Return(_a0 *models.ShareHolder, _a1 error) *MockShareHolderRepository_UpdateShareHolder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockShareHolderRepository_UpdateShareHolder_Call) RunAndReturn(run func(*models.ShareHolder) (*models.ShareHolder, error)) *MockShareHolderRepository_UpdateShareHolder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockShareHolderRepository creates a new instance of MockShareHolderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockShareHolderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockShareHolderRepository {
	mock := &MockShareHolderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
