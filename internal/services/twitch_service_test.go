package services

import (
	"testing"

	twitchMocks "github.com/hypebid/hypebid-app/internal/mocks/twitch"
	"github.com/hypebid/hypebid-app/internal/twitch"
	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {
	mockClient := new(twitchMocks.MockClient) // Assuming you have a mock for the twitch.Client
	service := &twitchService{client: mockClient}

	expectedToken := "mockAccessToken"
	mockClient.On("GetAccessToken").Return(expectedToken, nil)

	token, err := service.GetAccessToken()
	assert.NoError(t, err)
	assert.Equal(t, expectedToken, token)

	mockClient.AssertExpectations(t)
}

func TestGetUserByLogin(t *testing.T) {
	mockClient := new(twitchMocks.MockClient) // Assuming you have a mock for the twitch.Client
	service := &twitchService{client: mockClient}

	login := "testUser"
	expectedUserID := "12345"
	mockClient.On("GetUserByLogin", "mockAccessToken", login).Return(expectedUserID, nil)

	userID, err := service.GetUserByLogin("mockAccessToken", login)
	assert.NoError(t, err)
	assert.Equal(t, expectedUserID, userID)

	mockClient.AssertExpectations(t)
}

func TestGetUsersByLogin(t *testing.T) {
	mockClient := new(twitchMocks.MockClient) // Assuming you have a mock for the twitch.Client
	service := &twitchService{client: mockClient}

	logins := []string{"user1", "user2"}
	expectedUsers := []twitch.TwitchUser{
		{ID: "1", Login: "user1"},
		{ID: "2", Login: "user2"},
	}
	mockClient.On("GetUsersByLogin", "mockAccessToken", logins).Return(expectedUsers, nil)

	users, err := service.GetUsersByLogin("mockAccessToken", logins)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)

	mockClient.AssertExpectations(t)
}

func TestGetFollowerCount(t *testing.T) {
	mockClient := new(twitchMocks.MockClient) // Assuming you have a mock for the twitch.Client
	service := &twitchService{client: mockClient}

	broadcasterID := "12345"
	expectedFollowerCount := 100
	mockClient.On("GetFollowerCount", "mockAccessToken", broadcasterID).Return(expectedFollowerCount, nil)

	followerCount, err := service.GetFollowerCount("mockAccessToken", broadcasterID)
	assert.NoError(t, err)
	assert.Equal(t, expectedFollowerCount, followerCount)

	mockClient.AssertExpectations(t)
}
