package repositories

import (
	"errors"
	"log"

	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ MarketInstanceRepository = (*marketInstanceRepository)(nil)

type marketInstanceRepository struct {
	db *gorm.DB
}

func NewMarketInstanceRepository(db *gorm.DB) *marketInstanceRepository {
	return &marketInstanceRepository{db: db}
}

func (r *marketInstanceRepository) CreateMarketInstance(instance *models.MarketInstance) error {
	log.Printf("Creating MarketInstance: %+v", instance) // Log the instance
	return r.db.Create(instance).Error
}

func (r *marketInstanceRepository) GetMarketInstanceByNameAndUserID(name string, userID uuid.UUID) (*models.MarketInstance, error) {
	var instance models.MarketInstance
	err := r.db.Where("name = ? AND user_id = ?", name, userID).First(&instance).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle the case where no record is found, if needed
			return nil, nil // or return a specific error
		}
		return nil, err // Return the error if it's something else
	}
	return &instance, nil
}

func (r *marketInstanceRepository) GetMarketInstanceByID(marketInstanceID uuid.UUID) (*models.MarketInstance, error) {
	var instance models.MarketInstance
	err := r.db.Where("instance_id = ?", marketInstanceID).First(&instance).Error
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func (r *marketInstanceRepository) GetAllMarketInstances() ([]models.MarketInstance, error) {
	var instances []models.MarketInstance
	err := r.db.Find(&instances).Error
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (r *marketInstanceRepository) GetAllMarketInstancesByUserID(userID uuid.UUID) ([]models.MarketInstance, error) {
	var instances []models.MarketInstance
	err := r.db.Where("user_id = ?", userID).Find(&instances).Error
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (r *marketInstanceRepository) GetAllActiveMarketInstances() ([]models.MarketInstance, error) {
	var instances []models.MarketInstance
	err := r.db.Where("status = ?", "active").Find(&instances).Error
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (r *marketInstanceRepository) GetAllMarketInstancesByStatus(status string) ([]models.MarketInstance, error) {
	var instances []models.MarketInstance
	err := r.db.Where("status = ?", status).Find(&instances).Error
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (r *marketInstanceRepository) UpdateMarketInstance(instance *models.MarketInstance) error {
	return r.db.Save(instance).Error
}

func (r *marketInstanceRepository) IsMarketInstanceExists(instanceID uuid.UUID) bool {
	var instance models.MarketInstance
	err := r.db.Where("instance_id = ?", instanceID).First(&instance).Error
	return err == nil
}
