package services

import (
	"fmt"
	"time"

	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

var _ MarketInstanceService = (*marketInstanceService)(nil)

type marketInstanceService struct {
	marketInstanceRepo repositories.MarketInstanceRepository
}

func NewMarketInstanceService(marketInstanceRepo repositories.MarketInstanceRepository) *marketInstanceService {
	return &marketInstanceService{marketInstanceRepo: marketInstanceRepo}
}

func (s *marketInstanceService) CreateMarketInstance(name string, durationDays int, email string, userID uuid.UUID) (*models.MarketInstance, error) {

	// Validate that the name is not already taken
	// TODO: Add validation for naming conventions
	existingInstance, err := s.marketInstanceRepo.GetMarketInstanceByNameAndUserID(name, userID)
	if err != nil {
		return nil, err
	}

	if existingInstance != nil {
		return nil, fmt.Errorf("you've already used that name for a market")
	}

	instance := &models.MarketInstance{
		Name:         name,
		DurationDays: durationDays,
		UserID:       userID,
	}

	if err := s.marketInstanceRepo.CreateMarketInstance(instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (s *marketInstanceService) GetAllMarketInstancesByUserID(userID uuid.UUID) ([]models.MarketInstance, error) {
	instances, err := s.marketInstanceRepo.GetAllMarketInstancesByUserID(userID)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (s *marketInstanceService) GetAllMarketInstances() ([]models.MarketInstance, error) {
	instances, err := s.marketInstanceRepo.GetAllMarketInstances()
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (s *marketInstanceService) GetMarketInstanceByID(instanceID uuid.UUID) (*models.MarketInstance, error) {
	instance, err := s.marketInstanceRepo.GetMarketInstanceByID(instanceID)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (s *marketInstanceService) StartMarketInstance(instanceID uuid.UUID) (*models.MarketInstance, error) {
	marketInstance, err := s.GetMarketInstanceByID(instanceID)
	if err != nil {
		return nil, err
	}

	marketInstance.Status = "active"
	marketInstance.StartTimestamp = time.Now()
	marketInstance.EndTimestamp = time.Now().AddDate(0, 0, marketInstance.DurationDays)

	if err := s.marketInstanceRepo.UpdateMarketInstance(marketInstance); err != nil {
		return nil, err
	}

	return marketInstance, nil
}

func (s *marketInstanceService) IsMarketInstanceExists(instanceID uuid.UUID) bool {
	return s.marketInstanceRepo.IsMarketInstanceExists(instanceID)
}
