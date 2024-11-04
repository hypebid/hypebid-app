package services

import (
	"fmt"
	"time"

	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/internal/validator"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

var _ ShareHolderService = (*shareHolderService)(nil)

type shareHolderService struct {
	shareHolderRepo       repositories.ShareHolderRepository
	marketChannelService  MarketChannelService
	userService           UserService
	shareHolderValidators validator.ShareHolderValidators
}

func NewShareHolderService(shareHolderRepo repositories.ShareHolderRepository, marketChannelService MarketChannelService, userService UserService, shareHolderValidators validator.ShareHolderValidators) *shareHolderService {
	return &shareHolderService{
		shareHolderRepo:       shareHolderRepo,
		marketChannelService:  marketChannelService,
		userService:           userService,
		shareHolderValidators: shareHolderValidators,
	}
}

func (s *shareHolderService) InitializeShareHolder(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
	user, err := s.userService.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	shareHolder := &models.ShareHolder{
		UserID:          user.UserID,
		MarketChannelID: marketChannelID,
	}

	shareHolder, err = s.shareHolderRepo.CreateShareHolder(shareHolder)
	if err != nil {
		return nil, err
	}

	return shareHolder, nil
}

// Create shareholder for channel
func (s *shareHolderService) CreateShareHolderForChannel(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
	user, err := s.userService.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	marketChannel, err := s.marketChannelService.GetMarketChannelByID(marketChannelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get market channel: %w", err)
	}

	// TODO: Add logic to check if user already has a share holder for the channel
	shareHolder, err := s.shareHolderRepo.GetShareHolderByUserIDAndMarketChannelID(userID, marketChannelID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing share holder: %w", err)
	}

	now := time.Now()
	if shareHolder == nil {
		// Create new shareholder if none exists
		shareHolder = &models.ShareHolder{
			UserID:          user.UserID,
			MarketChannelID: marketChannel.ID,
			ShareCount:      marketChannel.TotalShares,
			AcquisitionDate: now,
		}
		return s.shareHolderRepo.CreateShareHolder(shareHolder)
	}

	// Update existing shareholder
	shareHolder.ShareCount = marketChannel.TotalShares
	shareHolder.AcquisitionDate = now
	return s.shareHolderRepo.UpdateShareHolder(shareHolder)
}

func (s *shareHolderService) GetShareHolderByUserIDAndMarketChannelID(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
	shareHolder, err := s.shareHolderRepo.GetShareHolderByUserIDAndMarketChannelID(userID, marketChannelID)
	if err != nil {
		return nil, err
	}
	return shareHolder, nil
}

func (s *shareHolderService) UpdateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error) {
	shareHolder, err := s.shareHolderRepo.UpdateShareHolder(shareHolder)
	if err != nil {
		return nil, err
	}
	return shareHolder, nil
}

func (s *shareHolderService) TransferShares(marketChannelID uuid.UUID, fromUserID uuid.UUID, toUserID uuid.UUID, shareCount int) error {
	fmt.Println("TransferShares called: ", marketChannelID, fromUserID, toUserID, shareCount)
	// validate intial values
	if err := s.shareHolderValidators.ValidateTransferSharesInput(marketChannelID, fromUserID, toUserID, shareCount); err != nil {
		return fmt.Errorf("invalid transfer values: %w", err)
	}

	fmt.Println("ValidateTransferSharesInput validated: ", marketChannelID, fromUserID, toUserID, shareCount)
	marketChannel, err := s.marketChannelService.GetMarketChannelByID(marketChannelID)
	if err != nil {
		return fmt.Errorf("failed to get market channel: %w", err)
	}

	fmt.Println("GetMarketChannelByID called: ", marketChannelID)
	fromShareHolder, err := s.shareHolderRepo.GetShareHolderByUserIDAndMarketChannelID(fromUserID, marketChannelID)
	if err != nil {
		return fmt.Errorf("failed to get share holder: %w", err)
	}
	fmt.Println("GetShareHolderByUserIDAndMarketChannelID called: ", fromUserID, marketChannelID)
	if fromShareHolder == nil {
		return fmt.Errorf("from share holder not found")
	}
	fmt.Println("GetShareHolderByUserIDAndMarketChannelID called: ", fromUserID, marketChannelID)
	toShareHolder, err := s.shareHolderRepo.GetShareHolderByUserIDAndMarketChannelID(toUserID, marketChannelID)
	if err != nil {
		// create new share holder
		toShareHolder, err = s.InitializeShareHolder(toUserID, marketChannelID)
		if err != nil {
			return fmt.Errorf("failed to initialize share holder: %w", err)
		}

		// set initial share count to 0
		toShareHolder.ShareCount = 0
	}
	fmt.Println("GetShareHolderByUserIDAndMarketChannelID called: ", toUserID, marketChannelID)
	if err := s.shareHolderValidators.ValidateTransferShares(shareCount, fromShareHolder.ShareCount, marketChannel.TotalShares); err != nil {
		return fmt.Errorf("invalid transfer: %w", err)
	}

	fmt.Println("ValidateTransferShares validated: ", shareCount, fromShareHolder.ShareCount, marketChannel.TotalShares)
	fromShareHolder.ShareCount -= shareCount
	fmt.Println("ShareCount updated: ", fromShareHolder.UserID, fromShareHolder.ShareCount)
	toShareHolder.ShareCount += shareCount
	fmt.Println("ShareCount updated: ", toShareHolder.UserID, toShareHolder.ShareCount)

	fmt.Println("UpdateShareHolder called: ", fromShareHolder.UserID, fromShareHolder.ShareCount)
	_, err = s.shareHolderRepo.UpdateShareHolder(fromShareHolder)
	if err != nil {
		return fmt.Errorf("failed to update share holder: %w", err)
	}
	fmt.Println("UpdateShareHolder called: ", toShareHolder.UserID, toShareHolder.ShareCount)
	_, err = s.shareHolderRepo.UpdateShareHolder(toShareHolder)
	if err != nil {
		return fmt.Errorf("failed to update share holder: %w", err)
	}

	fmt.Println("TransferShares completed: ", marketChannelID, fromUserID, toUserID, shareCount)
	return nil
}
