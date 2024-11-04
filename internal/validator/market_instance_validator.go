package validator

import (
	"errors"

	"github.com/hypebid/hypebid-app/internal/repositories"

	"github.com/google/uuid"
)

var _ MarketInstanceValidator = (*marketInstanceValidator)(nil)

type marketInstanceValidator struct {
	marketInstanceRepo repositories.MarketInstanceRepository
}

func NewMarketInstanceValidator(marketInstanceRepo repositories.MarketInstanceRepository) *marketInstanceValidator {
	return &marketInstanceValidator{marketInstanceRepo: marketInstanceRepo}
}

func (v *marketInstanceValidator) IsMarketInstanceExists(instanceID uuid.UUID) bool {
	_, err := v.marketInstanceRepo.GetMarketInstanceByID(instanceID)
	return err == nil
}

func (v *marketInstanceValidator) ValidateInstanceID(instanceID string) error {
	instanceUUID, err := uuid.Parse(instanceID)
	if err != nil {
		return errors.New("invalid instance ID")
	}

	if !v.IsMarketInstanceExists(instanceUUID) {
		return errors.New("market instance does not exist")
	}

	return nil
}
