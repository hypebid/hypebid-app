package services

import (
	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

var _ ChannelService = (*channelService)(nil)

type channelService struct {
	channelRepo repositories.ChannelRepository
}

func NewChannelService(channelRepo repositories.ChannelRepository) *channelService {
	return &channelService{channelRepo: channelRepo}
}

func (s *channelService) CreateChannel(name string, sharesTotal int) error {
	channel := &models.Channel{
		ChannelID: uuid.New(),
		Name:      name,
	}

	return s.channelRepo.CreateChannel(channel)
}

func (s *channelService) GetAllChannels() ([]models.Channel, error) {
	channels, err := s.channelRepo.GetAllChannels()
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func (s *channelService) GetChannelByName(name string) (*models.Channel, error) {
	return s.channelRepo.GetChannelByName(name)
}

func (s *channelService) GetChannelByID(channelID uuid.UUID) (*models.Channel, error) {
	return s.channelRepo.GetChannelByID(channelID)
}

func (s *channelService) IsChannelNameTaken(name string) bool {
	_, err := s.channelRepo.GetChannelByName(name)
	return err == nil
}

func (s *channelService) IsChannelExists(channelID uuid.UUID) bool {
	_, err := s.channelRepo.GetChannelByID(channelID)
	return err == nil
}
