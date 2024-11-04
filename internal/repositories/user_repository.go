package repositories

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ UserRepository = (*userRepository)(nil)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) IsUserExists(userUUID uuid.UUID) bool {
	var user models.User
	err := r.db.Where("user_id = ?", userUUID).First(&user).Error
	return err == nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByTwitchID(twitchID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("twitch_id = ?", twitchID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	err := r.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
