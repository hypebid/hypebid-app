package services

import (
	"fmt"

	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/internal/validator"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

var _ MarketChannelService = (*marketChannelService)(nil)

type marketChannelService struct {
	marketChannelRepo       repositories.MarketChannelRepository
	marketChannelValidators validator.MarketChannelValidators
}

func NewMarketChannelService(marketChannelRepo repositories.MarketChannelRepository, marketChannelValidators validator.MarketChannelValidators) *marketChannelService {
	return &marketChannelService{
		marketChannelRepo:       marketChannelRepo,
		marketChannelValidators: marketChannelValidators,
	}
}

func (s *marketChannelService) InitializeMarketChannel(instanceId uuid.UUID, channelId uuid.UUID, totalShares int, sharePrice float64) (*models.MarketChannel, error) {
	// verify instanceId exists
	if !s.marketChannelValidators.IsMarketInstanceExists(instanceId) {
		return nil, fmt.Errorf("market instance does not exist")
	}

	// verify channelId exists
	if !s.marketChannelValidators.IsChannelExists(channelId) {
		return nil, fmt.Errorf("channel does not exist")
	}

	marketChannel := &models.MarketChannel{
		MarketInstanceID: instanceId,
		ChannelID:        channelId,
		TotalShares:      totalShares,
		SharePrice:       sharePrice,
	}
	marketChannel, err := s.marketChannelRepo.CreateMarketChannel(marketChannel)
	if err != nil {
		return nil, fmt.Errorf("failed to create market channel: %w", err)
	}

	return marketChannel, nil
}

func (s *marketChannelService) GetMarketChannelByID(marketChannelID uuid.UUID) (*models.MarketChannel, error) {
	marketChannel, err := s.marketChannelRepo.GetMarketChannelByID(marketChannelID)
	if err != nil {
		return nil, fmt.Errorf("failed to get market channel: %w", err)
	}
	return marketChannel, nil
}

func (s *marketChannelService) GetMarketChannelsByInstanceID(instanceId uuid.UUID) ([]models.MarketChannel, error) {
	marketChannels, err := s.marketChannelRepo.GetMarketChannelsByInstanceID(instanceId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market channels: %w", err)
	}
	return marketChannels, nil
}
