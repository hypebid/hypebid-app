package services

import (
	"testing"

	"github.com/google/uuid"
	repoMocks "github.com/hypebid/hypebid-app/internal/mocks/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMember(t *testing.T) {
	mockRepo := new(repoMocks.MockMemberRepository)
	service := NewMemberService(mockRepo)

	marketInstanceID := uuid.New()
	userID := uuid.New()

	// Mocking the behavior of the repository
	mockRepo.On("CreateMember", mock.AnythingOfType("*models.Member")).Return(nil)

	err := service.CreateMember(marketInstanceID, userID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllMembersForInstance(t *testing.T) {
	mockRepo := new(repoMocks.MockMemberRepository)
	service := NewMemberService(mockRepo)

	marketInstanceID := uuid.New()
	expectedMembers := []models.Member{
		{MarketInstanceID: marketInstanceID, UserID: uuid.New()},
		{MarketInstanceID: marketInstanceID, UserID: uuid.New()},
	}

	mockRepo.On("GetAllMembersForInstance", marketInstanceID).Return(expectedMembers, nil)

	members, err := service.GetAllMembersForInstance(marketInstanceID)
	assert.NoError(t, err)
	assert.Equal(t, expectedMembers, members)

	mockRepo.AssertExpectations(t)
}

func TestGetAllMembersForInstance_Error(t *testing.T) {
	mockRepo := new(repoMocks.MockMemberRepository)
	service := NewMemberService(mockRepo)

	marketInstanceID := uuid.New()

	mockRepo.On("GetAllMembersForInstance", marketInstanceID).Return(nil, assert.AnError)

	members, err := service.GetAllMembersForInstance(marketInstanceID)
	assert.Error(t, err)
	assert.Nil(t, members)

	mockRepo.AssertExpectations(t)
}
