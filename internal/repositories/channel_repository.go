package repositories

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ ChannelRepository = (*channelRepository)(nil)

type channelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) *channelRepository {
	return &channelRepository{db: db}
}

func (r *channelRepository) CreateChannel(channel *models.Channel) error {
	return r.db.Create(channel).Error
}

func (r *channelRepository) GetAllChannels() ([]models.Channel, error) {
	var channels []models.Channel
	err := r.db.Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func (r *channelRepository) GetChannelByID(channelID uuid.UUID) (*models.Channel, error) {
	var channel models.Channel
	err := r.db.Where("channel_id = ?", channelID).First(&channel).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}

func (r *channelRepository) GetChannelByName(name string) (*models.Channel, error) {
	var channel models.Channel
	err := r.db.Where("name = ?", name).First(&channel).Error
	if err != nil {
		return nil, err
	}
	return &channel, nil
}
