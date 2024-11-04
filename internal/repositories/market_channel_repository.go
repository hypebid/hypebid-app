package repositories

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ MarketChannelRepository = (*marketChannelRepository)(nil)

type marketChannelRepository struct {
	db *gorm.DB
}

func NewMarketChannelRepository(db *gorm.DB) *marketChannelRepository {
	return &marketChannelRepository{db: db}
}

func (r *marketChannelRepository) CreateMarketChannel(marketChannel *models.MarketChannel) (*models.MarketChannel, error) {
	// create the market channel and return the object
	err := r.db.Create(marketChannel).Error
	return marketChannel, err
}

func (r *marketChannelRepository) GetMarketChannelByID(marketChannelID uuid.UUID) (*models.MarketChannel, error) {
	var marketChannel models.MarketChannel
	err := r.db.Where("id = ?", marketChannelID).First(&marketChannel).Error
	return &marketChannel, err
}

func (r *marketChannelRepository) GetMarketChannelsByInstanceID(instanceId uuid.UUID) ([]models.MarketChannel, error) {
	var marketChannels []models.MarketChannel
	err := r.db.Where("market_instance_id = ?", instanceId).Find(&marketChannels).Error
	return marketChannels, err
}
