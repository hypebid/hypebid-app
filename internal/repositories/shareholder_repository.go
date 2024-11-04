package repositories

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ ShareHolderRepository = (*shareHolderRepository)(nil)

type shareHolderRepository struct {
	db *gorm.DB
}

func NewShareHolderRepository(db *gorm.DB) *shareHolderRepository {
	return &shareHolderRepository{db: db}
}

func (r *shareHolderRepository) CreateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error) {
	err := r.db.Create(shareHolder).Error
	return shareHolder, err
}

func (r *shareHolderRepository) GetShareHolderByUserIDAndMarketChannelID(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error) {
	var shareHolder models.ShareHolder
	err := r.db.Where("user_id = ? AND market_channel_id = ?", userID, marketChannelID).First(&shareHolder).Error
	return &shareHolder, err
}

func (r *shareHolderRepository) UpdateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error) {
	err := r.db.Save(shareHolder).Error
	return shareHolder, err
}
