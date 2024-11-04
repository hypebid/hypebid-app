package services

import (
	"fmt"
	"testing"
	"time"

	repoMocks "github.com/hypebid/hypebid-app/internal/mocks/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMarketInstance(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	userID := uuid.New()
	instanceName := "Test Market"
	durationDays := 10

	// Mocking the behavior of the repository
	mockRepo.On("GetMarketInstanceByNameAndUserID", instanceName, userID).Return(nil, nil)
	mockRepo.On("CreateMarketInstance", mock.AnythingOfType("*models.MarketInstance")).Return(nil)

	instance, err := service.CreateMarketInstance(instanceName, durationDays, "test@example.com", userID)
	assert.NoError(t, err)
	assert.Equal(t, instanceName, instance.Name)
	assert.Equal(t, durationDays, instance.DurationDays)
	assert.Equal(t, userID, instance.UserID)

	mockRepo.AssertExpectations(t)
}

func TestCreateMarketInstance_NameAlreadyTaken(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	userID := uuid.New()
	instanceName := "Test Market"
	durationDays := 10

	existingInstance := &models.MarketInstance{Name: instanceName, UserID: userID}
	mockRepo.On("GetMarketInstanceByNameAndUserID", instanceName, userID).Return(existingInstance, nil)

	instance, err := service.CreateMarketInstance(instanceName, durationDays, "test@example.com", userID)
	assert.Error(t, err)
	assert.Nil(t, instance)
	assert.Equal(t, "you've already used that name for a market", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetAllMarketInstancesByUserID(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	userID := uuid.New()
	expectedInstances := []models.MarketInstance{
		{Name: "Market 1", UserID: userID},
		{Name: "Market 2", UserID: userID},
	}

	mockRepo.On("GetAllMarketInstancesByUserID", userID).Return(expectedInstances, nil)

	instances, err := service.GetAllMarketInstancesByUserID(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedInstances, instances)

	mockRepo.AssertExpectations(t)
}

func TestGetAllMarketInstances(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	expectedInstances := []models.MarketInstance{
		{Name: "Market 1"},
		{Name: "Market 2"},
	}

	mockRepo.On("GetAllMarketInstances").Return(expectedInstances, nil)

	instances, err := service.GetAllMarketInstances()
	assert.NoError(t, err)
	assert.Equal(t, expectedInstances, instances)

	mockRepo.AssertExpectations(t)
}

func TestGetMarketInstanceByID(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	instanceID := uuid.New()
	expectedInstance := &models.MarketInstance{InstanceID: instanceID}

	mockRepo.On("GetMarketInstanceByID", instanceID).Return(expectedInstance, nil)

	instance, err := service.GetMarketInstanceByID(instanceID)
	assert.NoError(t, err)
	assert.Equal(t, expectedInstance, instance)

	mockRepo.AssertExpectations(t)
}

func TestStartMarketInstance(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	instanceID := uuid.New()
	marketInstance := &models.MarketInstance{
		InstanceID:   instanceID,
		DurationDays: 10,
	}

	mockRepo.On("GetMarketInstanceByID", instanceID).Return(marketInstance, nil)
	mockRepo.On("UpdateMarketInstance", marketInstance).Return(nil)

	updatedInstance, err := service.StartMarketInstance(instanceID)
	assert.NoError(t, err)
	assert.Equal(t, "active", updatedInstance.Status)
	assert.WithinDuration(t, time.Now(), updatedInstance.StartTimestamp, time.Second)
	assert.WithinDuration(t, time.Now().AddDate(0, 0, 10), updatedInstance.EndTimestamp, time.Second)

	mockRepo.AssertExpectations(t)
}

func TestStartMarketInstance_NotFound(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	instanceID := uuid.New()

	mockRepo.On("GetMarketInstanceByID", instanceID).Return(nil, fmt.Errorf("not found"))

	instance, err := service.StartMarketInstance(instanceID)
	assert.Error(t, err)
	assert.Nil(t, instance)
	assert.Equal(t, "not found", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestIsMarketInstanceExists(t *testing.T) {
	mockRepo := new(repoMocks.MockMarketInstanceRepository)
	service := NewMarketInstanceService(mockRepo)

	instanceID := uuid.New()

	mockRepo.On("IsMarketInstanceExists", instanceID).Return(true)

	exists := service.IsMarketInstanceExists(instanceID)
	assert.True(t, exists)

	mockRepo.AssertExpectations(t)
}
