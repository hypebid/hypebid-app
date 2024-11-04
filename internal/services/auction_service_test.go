package services

import (
	"testing"
	"time"

	repoMocks "github.com/hypebid/hypebid-app/internal/mocks/repositories"
	serviceMocks "github.com/hypebid/hypebid-app/internal/mocks/services"
	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ AuctionService = (*auctionService)(nil)
var _ UserService = (*serviceMocks.MockUserService)(nil)
var _ ShareHolderService = (*serviceMocks.MockShareHolderService)(nil)
var _ repositories.AuctionRepository = (*repoMocks.MockAuctionRepository)(nil)

func TestCreateAuction(t *testing.T) {
	// Setup
	mockRepo := new(repoMocks.MockAuctionRepository)
	mockUserSvc := new(serviceMocks.MockUserService)
	mockShareHolderSvc := new(serviceMocks.MockShareHolderService)

	service := NewAuctionService(mockRepo, mockUserSvc, mockShareHolderSvc)

	marketInstanceID := uuid.New()
	marketChannelID := uuid.New()

	expectedAuction := &models.Auction{
		MarketInstanceID: marketInstanceID,
		MarketChannelID:  marketChannelID,
	}

	// Set expectations
	mockRepo.On("CreateAuction", mock.AnythingOfType("*models.Auction")).Return(expectedAuction, nil)

	// Execute
	auction, err := service.CreateAuction(marketInstanceID, marketChannelID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, auction)
	assert.Equal(t, marketInstanceID, auction.MarketInstanceID)
	assert.Equal(t, marketChannelID, auction.MarketChannelID)
	mockRepo.AssertExpectations(t)
}

func TestPlaceBid(t *testing.T) {
	tests := []struct {
		name                  string
		auctionSetup          func(*models.Auction)
		bidAmount             float64
		hasBalance            bool
		expectedError         string
		shouldValidateBalance bool
	}{
		{
			name: "successful bid",
			auctionSetup: func(a *models.Auction) {
				a.Status = "open"
				a.EndTime = time.Now().Add(time.Hour)
				a.HighestBid = 100
			},
			bidAmount:             150,
			hasBalance:            true,
			expectedError:         "",
			shouldValidateBalance: true,
		},
		{
			name: "bid too low",
			auctionSetup: func(a *models.Auction) {
				a.Status = "open"
				a.EndTime = time.Now().Add(time.Hour)
				a.HighestBid = 100
			},
			bidAmount:             90,
			hasBalance:            true,
			expectedError:         "bid amount must be greater than the current highest bid",
			shouldValidateBalance: false,
		},
		{
			name: "insufficient balance",
			auctionSetup: func(a *models.Auction) {
				a.Status = "open"
				a.EndTime = time.Now().Add(time.Hour)
				a.HighestBid = 100
			},
			bidAmount:             150,
			hasBalance:            false,
			expectedError:         "user does not have enough balance",
			shouldValidateBalance: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			mockRepo := new(repoMocks.MockAuctionRepository)
			mockUserSvc := new(serviceMocks.MockUserService)
			mockShareHolderSvc := new(serviceMocks.MockShareHolderService)

			svc := NewAuctionService(mockRepo, mockUserSvc, mockShareHolderSvc)
			service := svc.(*auctionService)

			auctionID := uuid.New()
			instanceID := uuid.New()
			userID := uuid.New()

			auction := &models.Auction{
				AuctionID: auctionID,
			}
			tt.auctionSetup(auction)

			// Store auction in activeAuctions
			service.activeAuctions.Store(auctionID, auction)

			// Set expectations only if we expect validation to occur
			if tt.shouldValidateBalance {
				mockUserSvc.On("ValidateUserBalance", userID, tt.bidAmount).Return(tt.hasBalance)
			}

			if tt.hasBalance && tt.bidAmount > auction.HighestBid {
				mockRepo.On("UpdateAuction", mock.AnythingOfType("*models.Auction")).Return(nil)
			}

			// Execute
			err := service.PlaceBid(instanceID, auctionID, userID, tt.bidAmount)

			// Assert
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.bidAmount, auction.HighestBid)
				assert.Equal(t, userID, auction.HighestBidderID)
			}

			mockRepo.AssertExpectations(t)
			mockUserSvc.AssertExpectations(t)
		})
	}
}
