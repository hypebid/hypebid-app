package services

import (
	"testing"

	"github.com/google/uuid"
	repoMocks "github.com/hypebid/hypebid-app/internal/mocks/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateChannel(t *testing.T) {
	mockRepo := new(repoMocks.MockChannelRepository)
	service := NewChannelService(mockRepo)

	channelName := "Test Channel"
	sharesTotal := 100
	// channelID := uuid.New()

	mockRepo.On("CreateChannel", mock.AnythingOfType("*models.Channel")).Return(nil)

	err := service.CreateChannel(channelName, sharesTotal)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllChannels(t *testing.T) {
	mockRepo := new(repoMocks.MockChannelRepository)
	service := NewChannelService(mockRepo)

	expectedChannels := []models.Channel{
		{ChannelID: uuid.New(), Name: "Channel 1"},
		{ChannelID: uuid.New(), Name: "Channel 2"},
	}

	mockRepo.On("GetAllChannels").Return(expectedChannels, nil)

	channels, err := service.GetAllChannels()
	assert.NoError(t, err)
	assert.Equal(t, expectedChannels, channels)

	mockRepo.AssertExpectations(t)
}

func TestGetChannelByName(t *testing.T) {
	mockRepo := new(repoMocks.MockChannelRepository)
	service := NewChannelService(mockRepo)

	channelName := "Test Channel"
	expectedChannel := &models.Channel{ChannelID: uuid.New(), Name: channelName}

	mockRepo.On("GetChannelByName", channelName).Return(expectedChannel, nil)

	channel, err := service.GetChannelByName(channelName)
	assert.NoError(t, err)
	assert.Equal(t, expectedChannel, channel)

	mockRepo.AssertExpectations(t)
}

func TestIsChannelNameTaken(t *testing.T) {
	mockRepo := new(repoMocks.MockChannelRepository)
	service := NewChannelService(mockRepo)

	channelName := "Test Channel"
	mockRepo.On("GetChannelByName", channelName).Return(&models.Channel{}, nil)

	taken := service.IsChannelNameTaken(channelName)
	assert.True(t, taken)

	mockRepo.AssertExpectations(t)
}

func TestIsChannelExists(t *testing.T) {
	mockRepo := new(repoMocks.MockChannelRepository)
	service := NewChannelService(mockRepo)

	channelID := uuid.New()
	mockRepo.On("GetChannelByID", channelID).Return(&models.Channel{}, nil)

	exists := service.IsChannelExists(channelID)
	assert.True(t, exists)

	mockRepo.AssertExpectations(t)
}
