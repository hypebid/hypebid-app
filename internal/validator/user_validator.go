package validator

import (
	"errors"
	"fmt"

	"github.com/hypebid/hypebid-app/internal/repositories"

	"github.com/google/uuid"
)

var _ UserValidator = (*userValidator)(nil)

type userValidator struct {
	userRepo repositories.UserRepository
}

func NewUserValidator(userRepo repositories.UserRepository) *userValidator {
	return &userValidator{userRepo: userRepo}
}

func (v *userValidator) IsUserExists(userID uuid.UUID) bool {
	return v.userRepo.IsUserExists(userID)
}

// TODO: Currency should be on member not user
func (v *userValidator) ValidateUserBalance(userID uuid.UUID, amount float64) bool {
	user, err := v.userRepo.GetUserByID(userID)
	if err != nil {
		return false
	}

	if user.Currency < amount {
		return false
	}

	return true
}

// TODO: Currency should be on member not user
func (v *userValidator) ValidateUserForBid(userID uuid.UUID, amount float64) error {
	user, err := v.userRepo.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if user.Currency < amount {
		return errors.New("insufficient funds")
	}

	return nil
}
