package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	repoMocks "github.com/hypebid/hypebid-app/internal/mocks/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	username := "testuser"
	email := "test@example.com"
	password := "password123"

	// Hash the password for testing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	expectedUser := &models.User{
		UserID:       uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: new(string),
	}

	*expectedUser.PasswordHash = string(hashedPassword)

	mockRepo.On("GetUserByEmail", email).Return(nil, errors.New("record not found")) // No existing user
	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(expectedUser, nil)

	user, err := service.CreateUser(username, email, password)
	assert.NoError(t, err)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)

	// Verify password hashing
	err = bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password))
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_UserExists(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	email := "test@example.com"
	existingUser := &models.User{Email: email}
	mockRepo.On("GetUserByEmail", email).Return(existingUser, nil) // Existing user

	user, err := service.CreateUser("testuser", email, "password123")
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "user already exists", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetUserByEmail(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	email := "test@example.com"
	expectedUser := &models.User{Email: email}
	mockRepo.On("GetUserByEmail", email).Return(expectedUser, nil)

	user, err := service.GetUserByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockRepo.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	userID := uuid.New()
	expectedUser := &models.User{UserID: userID}
	mockRepo.On("GetUserByID", userID).Return(expectedUser, nil)

	user, err := service.GetUserByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockRepo.AssertExpectations(t)
}

func TestFindOrCreateTwitchUser_UserExists(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	twitchData := &models.TwitchUser{
		ID:    "twitchID",
		Login: "testuser",
		Email: "test@example.com",
	}
	token := &oauth2.Token{}

	// Expected user that will be created
	expectedUser := &models.User{
		UserID:   uuid.New(),
		Username: twitchData.Login,
		Email:    twitchData.Email,
		TwitchID: &twitchData.ID,
	}

	// Mock the Twitch user lookup and update
	existingTwitchUser := &models.TwitchUser{ID: "twitchID"}
	mockTwitchRepo.On("GetTwitchUserByID", twitchData.ID).Return(existingTwitchUser, nil)
	mockTwitchRepo.On("UpdateTwitchUser", mock.MatchedBy(func(tu *models.TwitchUser) bool {
		return tu.ID == "twitchID"
	})).Return(nil)

	// Mock the user lookup
	mockRepo.On("GetUserByTwitchID", twitchData.ID).Return(nil, nil)
	fmt.Println("GetUserByTwitchID called: ", twitchData.ID)

	// Mock the creation of a new user
	mockRepo.On("CreateUser", mock.MatchedBy(func(u *models.User) bool {
		return u.Username == twitchData.Login &&
			u.Email == twitchData.Email &&
			u.TwitchID != nil &&
			*u.TwitchID == twitchData.ID
	})).Return(expectedUser, nil)

	user, err := service.FindOrCreateTwitchUser(twitchData, token)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, twitchData.ID, *user.TwitchID)

	mockRepo.AssertExpectations(t)
	mockTwitchRepo.AssertExpectations(t)
}

// func TestLinkTwitchAccount(t *testing.T) {
// 	mockRepo := new(repoMocks.MockUserRepository)
// 	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
// 	service := NewUserService(mockRepo, mockTwitchRepo)

// 	userID := uuid.New()
// 	twitchData := &models.TwitchUser{ID: "twitchID"}
// 	token := &oauth2.Token{AccessToken: "accessToken", RefreshToken: "refreshToken"}

// 	mockRepo.On("GetUserByTwitchID", twitchData.ID).Return(nil, nil)       // No user found
// 	mockTwitchRepo.On("GetTwitchUserByID", twitchData.ID).Return(nil, nil) // No existing Twitch user
// 	mockRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)
// 	mockTwitchRepo.On("CreateTwitchUser", twitchData).Return(nil)

// 	err := service.LinkTwitchAccount(userID, twitchData, token)
// 	assert.NoError(t, err)

// 	mockRepo.AssertExpectations(t)
// }

func TestAddCurrency(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	userID := uuid.New()
	user := &models.User{UserID: userID, Currency: 100}
	amount := 50.0

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", user).Return(user, nil)

	updatedUser, err := service.AddCurrency(userID, amount)
	assert.NoError(t, err)
	assert.Equal(t, 150.0, updatedUser.Currency)

	mockRepo.AssertExpectations(t)
}

func TestSubtractCurrency(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	userID := uuid.New()
	user := &models.User{UserID: userID, Currency: 100}
	amount := 50.0

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", user).Return(user, nil)

	updatedUser, err := service.SubtractCurrency(userID, amount)
	assert.NoError(t, err)
	assert.Equal(t, 50.0, updatedUser.Currency)

	mockRepo.AssertExpectations(t)
}

func TestValidateUserBalance(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	userID := uuid.New()
	user := &models.User{UserID: userID, Currency: 100}

	mockRepo.On("GetUserByID", userID).Return(user, nil)

	valid := service.ValidateUserBalance(userID, 50)
	assert.True(t, valid)

	invalid := service.ValidateUserBalance(userID, 150)
	assert.False(t, invalid)

	mockRepo.AssertExpectations(t)
}

func TestValidateUserForBid(t *testing.T) {
	mockRepo := new(repoMocks.MockUserRepository)
	mockTwitchRepo := new(repoMocks.MockTwitchRepository)
	service := NewUserService(mockRepo, mockTwitchRepo)

	userID := uuid.New()
	user := &models.User{UserID: userID, Currency: 100}

	mockRepo.On("GetUserByID", userID).Return(user, nil)

	err := service.ValidateUserForBid(userID, 50)
	assert.NoError(t, err)

	err = service.ValidateUserForBid(userID, 150)
	assert.Error(t, err)
	assert.Equal(t, "insufficient funds", err.Error())

	mockRepo.AssertExpectations(t)
}
