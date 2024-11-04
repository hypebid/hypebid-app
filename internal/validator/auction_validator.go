package validator

import (
	"errors"

	"github.com/hypebid/hypebid-app/internal/repositories"

	"github.com/google/uuid"
)

var _ AuctionValidator = (*auctionValidator)(nil)

type auctionValidator struct {
	auctionRepo repositories.AuctionRepository
}

func NewAuctionValidator(auctionRepo repositories.AuctionRepository) *auctionValidator {
	return &auctionValidator{auctionRepo: auctionRepo}
}

type auctionValidators struct {
	auctionValidator        *auctionValidator
	channelValidator        ChannelValidator
	marketInstanceValidator MarketInstanceValidator
	userValidator           UserValidator
}

func NewAuctionValidators(
	auctionRepo repositories.AuctionRepository,
	channelValidator ChannelValidator,
	marketInstanceValidator MarketInstanceValidator,
	userValidator UserValidator,
) *auctionValidators {
	return &auctionValidators{
		auctionValidator:        NewAuctionValidator(auctionRepo),
		channelValidator:        channelValidator,
		marketInstanceValidator: marketInstanceValidator,
		userValidator:           userValidator,
	}
}

func (v *auctionValidator) IsAuctionExists(auctionID uuid.UUID) bool {
	_, err := v.auctionRepo.GetAuctionByID(auctionID)
	return err == nil
}

func (v *auctionValidator) ValidateAuctionID(auctionID string) error {
	auctionUUID, err := uuid.Parse(auctionID)
	if err != nil {
		return errors.New("invalid auction ID")
	}

	if !v.IsAuctionExists(auctionUUID) {
		return errors.New("auction not found")
	}

	return nil
}

func (v *auctionValidators) IsAuctionExists(auctionID uuid.UUID) bool {
	return v.auctionValidator.IsAuctionExists(auctionID)
}

func (v *auctionValidators) ValidateInstanceID(instanceID string) error {
	return v.marketInstanceValidator.ValidateInstanceID(instanceID)
}

func (v *auctionValidators) ValidateChannelID(channelID string) error {
	return v.channelValidator.ValidateChannelID(channelID)
}

func (v *auctionValidators) ValidateAuctionID(auctionID string) error {
	return v.auctionValidator.ValidateAuctionID(auctionID)
}

func (v *auctionValidators) ValidateUserForBid(userID uuid.UUID, amount float64) error {
	return v.userValidator.ValidateUserForBid(userID, amount)
}
