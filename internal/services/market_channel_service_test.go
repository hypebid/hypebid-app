package services

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	repoMocks "github.com/hypebid/hypebid-app/internal/mocks/repositories"
	validatorMocks "github.com/hypebid/hypebid-app/internal/mocks/validator"
	"github.com/hypebid/hypebid-app/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestInitializeMarketChannel(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketChannelRepository)
	mockValidator := new(validatorMocks.MockMarketChannelValidators)
	service := NewMarketChannelService(mockRepo, mockValidator)

	instanceID := uuid.New()
	channelID := uuid.New()
	totalShares := 100
	sharePrice := 10.0

	mockValidator.On("IsMarketInstanceExists", instanceID).Return(true)
	mockValidator.On("IsChannelExists", channelID).Return(true)

	marketChannel := &models.MarketChannel{
		MarketInstanceID: instanceID,
		ChannelID:        channelID,
		TotalShares:      totalShares,
		SharePrice:       sharePrice,
	}

	mockRepo.On("CreateMarketChannel", marketChannel).Return(marketChannel, nil)

	result, err := service.InitializeMarketChannel(instanceID, channelID, totalShares, sharePrice)
	assert.NoError(t, err)
	assert.Equal(t, marketChannel, result)

	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestInitializeMarketChannel_InstanceNotExists(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketChannelRepository)
	mockValidator := new(validatorMocks.MockMarketChannelValidators)
	service := NewMarketChannelService(mockRepo, mockValidator)

	instanceID := uuid.New()
	channelID := uuid.New()
	totalShares := 100
	sharePrice := 10.0

	mockValidator.On("IsMarketInstanceExists", instanceID).Return(false)

	result, err := service.InitializeMarketChannel(instanceID, channelID, totalShares, sharePrice)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "market instance does not exist", err.Error())

	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestInitializeMarketChannel_ChannelNotExists(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketChannelRepository)
	mockValidator := new(validatorMocks.MockMarketChannelValidators)
	service := NewMarketChannelService(mockRepo, mockValidator)

	instanceID := uuid.New()
	channelID := uuid.New()
	totalShares := 100
	sharePrice := 10.0

	mockValidator.On("IsMarketInstanceExists", instanceID).Return(true)
	mockValidator.On("IsChannelExists", channelID).Return(false)

	result, err := service.InitializeMarketChannel(instanceID, channelID, totalShares, sharePrice)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "channel does not exist", err.Error())

	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

func TestGetMarketChannelByID(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketChannelRepository)
	service := NewMarketChannelService(mockRepo, nil)

	marketChannelID := uuid.New()
	expectedMarketChannel := &models.MarketChannel{MarketInstanceID: marketChannelID}

	mockRepo.On("GetMarketChannelByID", marketChannelID).Return(expectedMarketChannel, nil)

	result, err := service.GetMarketChannelByID(marketChannelID)
	assert.NoError(t, err)
	assert.Equal(t, expectedMarketChannel, result)

	mockRepo.AssertExpectations(t)
}

func TestGetMarketChannelByID_Error(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketChannelRepository)
	service := NewMarketChannelService(mockRepo, nil)

	marketChannelID := uuid.New()

	mockRepo.On("GetMarketChannelByID", marketChannelID).Return(nil, fmt.Errorf("not found"))

	result, err := service.GetMarketChannelByID(marketChannelID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to get market channel: not found", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetMarketChannelsByInstanceID(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketChannelRepository)
	service := NewMarketChannelService(mockRepo, nil)

	instanceID := uuid.New()
	expectedMarketChannels := []models.MarketChannel{
		{MarketInstanceID: instanceID, ChannelID: uuid.New(), TotalShares: 100, SharePrice: 10.0},
	}

	mockRepo.On("GetMarketChannelsByInstanceID", instanceID).Return(expectedMarketChannels, nil)

	result, err := service.GetMarketChannelsByInstanceID(instanceID)
	assert.NoError(t, err)
	assert.Equal(t, expectedMarketChannels, result)

	mockRepo.AssertExpectations(t)
}

func TestGetMarketChannelsByInstanceID_Error(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketChannelRepository)
	service := NewMarketChannelService(mockRepo, nil)

	instanceID := uuid.New()

	mockRepo.On("GetMarketChannelsByInstanceID", instanceID).Return(nil, fmt.Errorf("not found"))

	result, err := service.GetMarketChannelsByInstanceID(instanceID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to get market channels: not found", err.Error())

	mockRepo.AssertExpectations(t)
}
