package repositories

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ TwitchRepository = (*twitchRepository)(nil)

type twitchRepository struct {
	db *gorm.DB
}

func NewTwitchRepository(db *gorm.DB) *twitchRepository {
	return &twitchRepository{db: db}
}

func (r *twitchRepository) CreateTwitchUser(twitchUser *models.TwitchUser) error {
	return r.db.Create(twitchUser).Error
}

func (r *twitchRepository) GetTwitchUserByID(id string) (*models.TwitchUser, error) {
	var twitchUser models.TwitchUser
	err := r.db.Where("id = ?", id).First(&twitchUser).Error
	if err != nil {
		return nil, err
	}
	return &twitchUser, nil
}

func (r *twitchRepository) UpdateTwitchUser(twitchUser *models.TwitchUser) error {
	return r.db.Save(twitchUser).Error
}

func (r *twitchRepository) GetTwitchUserByUserID(userID uuid.UUID) (*models.TwitchUser, error) {
	var twitchUser models.TwitchUser
	err := r.db.Where("user_id = ?", userID).First(&twitchUser).Error
	if err != nil {
		return nil, err
	}
	return &twitchUser, nil
}
