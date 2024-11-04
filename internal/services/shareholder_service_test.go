package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	repoMocks "github.com/hypebid/hypebid-app/internal/mocks/repositories"
	serviceMocks "github.com/hypebid/hypebid-app/internal/mocks/services"
	validatorMocks "github.com/hypebid/hypebid-app/internal/mocks/validator"
	"github.com/hypebid/hypebid-app/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitializeShareHolder(t *testing.T) {
	mockRepo := new(repoMocks.MockShareHolderRepository)
	mockUserService := new(serviceMocks.MockUserService)
	mockValidator := new(validatorMocks.MockShareHolderValidators)
	service := NewShareHolderService(mockRepo, nil, mockUserService, mockValidator)

	userID := uuid.New()
	marketChannelID := uuid.New()

	user := &models.User{UserID: userID}
	mockUserService.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("CreateShareHolder", mock.AnythingOfType("*models.ShareHolder")).Return(&models.ShareHolder{UserID: userID, MarketChannelID: marketChannelID}, nil)

	shareHolder, err := service.InitializeShareHolder(userID, marketChannelID)
	assert.NoError(t, err)
	assert.Equal(t, userID, shareHolder.UserID)
	assert.Equal(t, marketChannelID, shareHolder.MarketChannelID)

	mockUserService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateShareHolderForChannel(t *testing.T) {
	mockRepo := new(repoMocks.MockShareHolderRepository)
	mockMarketChannelService := new(serviceMocks.MockMarketChannelService)
	mockUserService := new(serviceMocks.MockUserService)
	mockValidator := new(validatorMocks.MockShareHolderValidators)
	service := NewShareHolderService(mockRepo, mockMarketChannelService, mockUserService, mockValidator)

	userID := uuid.New()
	marketChannelID := uuid.New()
	user := &models.User{UserID: userID}
	marketChannel := &models.MarketChannel{ID: marketChannelID, TotalShares: 100}

	mockUserService.On("GetUserByID", userID).Return(user, nil)
	mockMarketChannelService.On("GetMarketChannelByID", marketChannelID).Return(marketChannel, nil)
	mockRepo.On("GetShareHolderByUserIDAndMarketChannelID", userID, marketChannelID).Return(nil, nil)
	mockRepo.On("CreateShareHolder", mock.AnythingOfType("*models.ShareHolder")).Return(&models.ShareHolder{UserID: userID, MarketChannelID: marketChannelID, ShareCount: marketChannel.TotalShares, AcquisitionDate: time.Now()}, nil)

	shareHolder, err := service.CreateShareHolderForChannel(userID, marketChannelID)
	assert.NoError(t, err)
	assert.Equal(t, userID, shareHolder.UserID)
	assert.Equal(t, marketChannelID, shareHolder.MarketChannelID)
	assert.Equal(t, marketChannel.TotalShares, shareHolder.ShareCount)

	mockUserService.AssertExpectations(t)
	mockMarketChannelService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateShareHolderForChannel_UpdateExistingShareHolder(t *testing.T) {
	mockRepo := new(repoMocks.MockShareHolderRepository)
	mockMarketChannelService := new(serviceMocks.MockMarketChannelService)
	mockUserService := new(serviceMocks.MockUserService)
	mockValidator := new(validatorMocks.MockShareHolderValidators)
	service := NewShareHolderService(mockRepo, mockMarketChannelService, mockUserService, mockValidator)

	userID := uuid.New()
	marketChannelID := uuid.New()
	user := &models.User{UserID: userID}
	marketChannel := &models.MarketChannel{ID: marketChannelID, TotalShares: 100}
	existingShareHolder := &models.ShareHolder{UserID: userID, MarketChannelID: marketChannelID, ShareCount: 50}

	mockUserService.On("GetUserByID", userID).Return(user, nil)
	mockMarketChannelService.On("GetMarketChannelByID", marketChannelID).Return(marketChannel, nil)
	mockRepo.On("GetShareHolderByUserIDAndMarketChannelID", userID, marketChannelID).Return(existingShareHolder, nil)
	mockRepo.On("UpdateShareHolder", existingShareHolder).Return(existingShareHolder, nil)

	shareHolder, err := service.CreateShareHolderForChannel(userID, marketChannelID)
	assert.NoError(t, err)
	assert.Equal(t, existingShareHolder, shareHolder)
	assert.Equal(t, marketChannel.TotalShares, shareHolder.ShareCount)

	mockUserService.AssertExpectations(t)
	mockMarketChannelService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestGetShareHolderByUserIDAndMarketChannelID(t *testing.T) {
	mockRepo := new(repoMocks.MockShareHolderRepository)
	mockValidator := new(validatorMocks.MockShareHolderValidators)
	service := NewShareHolderService(mockRepo, nil, nil, mockValidator)

	userID := uuid.New()
	marketChannelID := uuid.New()
	expectedShareHolder := &models.ShareHolder{UserID: userID, MarketChannelID: marketChannelID}

	mockRepo.On("GetShareHolderByUserIDAndMarketChannelID", userID, marketChannelID).Return(expectedShareHolder, nil)

	shareHolder, err := service.GetShareHolderByUserIDAndMarketChannelID(userID, marketChannelID)
	assert.NoError(t, err)
	assert.Equal(t, expectedShareHolder, shareHolder)

	mockRepo.AssertExpectations(t)
}

func TestUpdateShareHolder(t *testing.T) {
	mockRepo := new(repoMocks.MockShareHolderRepository)
	mockValidator := new(validatorMocks.MockShareHolderValidators)
	service := NewShareHolderService(mockRepo, nil, nil, mockValidator)

	shareHolder := &models.ShareHolder{UserID: uuid.New()}

	mockRepo.On("UpdateShareHolder", shareHolder).Return(shareHolder, nil)

	updatedShareHolder, err := service.UpdateShareHolder(shareHolder)
	assert.NoError(t, err)
	assert.Equal(t, shareHolder, updatedShareHolder)

	mockRepo.AssertExpectations(t)
}

// func TestTransferShares(t *testing.T) {
// 	mockRepo := new(repoMocks.MockShareHolderRepository)
// 	mockMarketChannelService := new(serviceMocks.MockMarketChannelService)
// 	mockUserService := new(serviceMocks.MockUserService)
// 	mockValidator := new(validatorMocks.MockShareHolderValidators)
// 	service := NewShareHolderService(mockRepo, mockMarketChannelService, mockUserService, mockValidator)

// 	marketChannelID := uuid.New()
// 	fromUserID := uuid.New()
// 	toUserID := uuid.New()
// 	shareCount := 10
// 	marketChannel := &models.MarketChannel{ID: marketChannelID, TotalShares: 100}
// 	fromShareHolder := &models.ShareHolder{UserID: fromUserID, MarketChannelID: marketChannelID, ShareCount: 50}
// 	toShareHolder := &models.ShareHolder{UserID: toUserID, MarketChannelID: marketChannelID, ShareCount: 20}

// 	mockValidator.On("ValidateTransferSharesInput", marketChannelID, fromUserID, toUserID, shareCount).Return(nil)
// 	mockMarketChannelService.On("GetMarketChannelByID", marketChannelID).Return(marketChannel, nil)
// 	mockRepo.On("GetShareHolderByUserIDAndMarketChannelID", fromUserID, marketChannelID).Return(fromShareHolder, nil)
// 	mockRepo.On("GetShareHolderByUserIDAndMarketChannelID", toUserID, marketChannelID).Return(toShareHolder, nil)
// 	mockValidator.On("ValidateTransferShares", shareCount, fromShareHolder.ShareCount, marketChannel.TotalShares).Return(nil)

// 	err := service.TransferShares(marketChannelID, fromUserID, toUserID, shareCount)

// 	mockRepo.On("UpdateShareHolder", fromShareHolder).Return(fromShareHolder, nil)
// 	mockRepo.On("UpdateShareHolder", toShareHolder).Return(toShareHolder, nil)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 40, fromShareHolder.ShareCount)
// 	assert.Equal(t, 30, toShareHolder.ShareCount)

// 	mockRepo.AssertExpectations(t)
// }

// func TestTransferShares_CreateToShareHolder(t *testing.T) {
// 	mockRepo := new(repoMocks.MockShareHolderRepository)
// 	mockMarketChannelService := new(serviceMocks.MockMarketChannelService)
// 	mockUserService := new(serviceMocks.MockUserService)
// 	mockValidator := new(validatorMocks.MockShareHolderValidators)
// 	service := NewShareHolderService(mockRepo, mockMarketChannelService, mockUserService, mockValidator)

// 	marketChannelID := uuid.New()
// 	fromUserID := uuid.New()
// 	toUserID := uuid.New()
// 	shareCount := 10

// 	// Setup test data
// 	marketChannel := &models.MarketChannel{ID: marketChannelID, TotalShares: 100}
// 	fromShareHolder := &models.ShareHolder{
// 		UserID:          fromUserID,
// 		MarketChannelID: marketChannelID,
// 		ShareCount:      50,
// 		AcquisitionDate: time.Now(),
// 	}
// 	toUser := &models.User{UserID: toUserID}

// 	// The initial shareholder created by InitializeShareHolder
// 	initialToShareHolder := &models.ShareHolder{
// 		UserID:          toUserID,
// 		MarketChannelID: marketChannelID,
// 		ShareCount:      0, // Explicitly set to 0
// 		AcquisitionDate: time.Now(),
// 	}

// 	// Setup mock expectations in the exact order they will be called
// 	mockValidator.On("ValidateTransferSharesInput", marketChannelID, fromUserID, toUserID, shareCount).Return(nil)
// 	fmt.Println("ValidateTransferSharesInput validated: ", marketChannelID, fromUserID, toUserID, shareCount)
// 	mockMarketChannelService.On("GetMarketChannelByID", marketChannelID).Return(marketChannel, nil)
// 	fmt.Println("GetMarketChannelByID called: ", marketChannelID)
// 	mockRepo.On("GetShareHolderByUserIDAndMarketChannelID", fromUserID, marketChannelID).Return(fromShareHolder, nil)
// 	fmt.Println("GetShareHolderByUserIDAndMarketChannelID called: ", fromUserID, marketChannelID)
// 	mockRepo.On("GetShareHolderByUserIDAndMarketChannelID", toUserID, marketChannelID).Return(nil, nil)
// 	fmt.Println("GetShareHolderByUserIDAndMarketChannelID called: ", toUserID, marketChannelID)
// 	mockUserService.On("GetUserByID", toUserID).Return(toUser, nil)
// 	fmt.Println("GetUserByID called: ", toUserID)
// 	mockRepo.On("CreateShareHolder", mock.MatchedBy(func(sh *models.ShareHolder) bool {
// 		return sh.UserID == toUserID && sh.MarketChannelID == marketChannelID
// 	})).Return(initialToShareHolder, nil)
// 	fmt.Println("CreateShareHolder called: ", toUserID, marketChannelID)
// 	mockValidator.On("ValidateTransferShares", shareCount, fromShareHolder.ShareCount, marketChannel.TotalShares).Return(nil)
// 	fmt.Println("ValidateTransferShares called: ", shareCount, fromShareHolder.ShareCount, marketChannel.TotalShares)
// 	// The final state after updates
// 	updatedFromShareHolder := &models.ShareHolder{
// 		UserID:          fromUserID,
// 		MarketChannelID: marketChannelID,
// 		ShareCount:      40, // 50 - 10
// 		AcquisitionDate: fromShareHolder.AcquisitionDate,
// 	}

// 	updatedToShareHolder := &models.ShareHolder{
// 		UserID:          toUserID,
// 		MarketChannelID: marketChannelID,
// 		ShareCount:      shareCount,
// 		AcquisitionDate: initialToShareHolder.AcquisitionDate,
// 	}

// 	mockRepo.On("UpdateShareHolder", mock.MatchedBy(func(sh *models.ShareHolder) bool {
// 		return sh.UserID == fromUserID && sh.ShareCount == 40
// 	})).Return(updatedFromShareHolder, nil)
// 	fmt.Println("UpdateShareHolder called: ", fromUserID, 40)
// 	mockRepo.On("UpdateShareHolder", mock.MatchedBy(func(sh *models.ShareHolder) bool {
// 		return sh.UserID == toUserID && sh.ShareCount == shareCount
// 	})).Return(updatedToShareHolder, nil)
// 	fmt.Println("UpdateShareHolder called: ", toUserID, shareCount)
// 	err := service.TransferShares(marketChannelID, fromUserID, toUserID, shareCount)
// 	fmt.Println("TransferShares called: ", marketChannelID, fromUserID, toUserID, shareCount)
// 	assert.NoError(t, err)

// 	mockRepo.AssertExpectations(t)
// 	mockMarketChannelService.AssertExpectations(t)
// 	mockUserService.AssertExpectations(t)
// 	mockValidator.AssertExpectations(t)
// }
