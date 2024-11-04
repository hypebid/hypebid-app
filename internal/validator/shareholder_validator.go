package validator

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hypebid/hypebid-app/internal/repositories"
)

var _ ShareHolderValidator = (*shareHolderValidator)(nil)

type shareHolderValidator struct {
	shareHolderRepo repositories.ShareHolderRepository
}

func NewShareHolderValidator(shareHolderRepo repositories.ShareHolderRepository) *shareHolderValidator {
	return &shareHolderValidator{shareHolderRepo: shareHolderRepo}
}

func (v *shareHolderValidator) IsShareHolderExists(userID uuid.UUID, marketChannelID uuid.UUID) bool {
	_, err := v.shareHolderRepo.GetShareHolderByUserIDAndMarketChannelID(userID, marketChannelID)
	return err == nil
}

type shareHolderValidators struct {
	shareHolder            *shareHolderValidator
	marketChannelValidator MarketChannelValidator
}

func NewShareHolderValidators(
	shareHolderRepo repositories.ShareHolderRepository,
	marketChannelValidator MarketChannelValidator,
) ShareHolderValidators {
	return &shareHolderValidators{
		shareHolder:            NewShareHolderValidator(shareHolderRepo),
		marketChannelValidator: marketChannelValidator,
	}
}

func (v *shareHolderValidators) IsShareHolderExists(userID uuid.UUID, marketChannelID uuid.UUID) bool {
	return v.shareHolder.IsShareHolderExists(userID, marketChannelID)
}

func (v *shareHolderValidators) ValidateTransferSharesInput(marketChannelID, fromUserID, toUserID uuid.UUID, shareCount int) error {
	// Check if market channel exists
	if !v.marketChannelValidator.IsMarketChannelExists(marketChannelID) {
		return fmt.Errorf("market channel does not exist")
	}

	// Validate share count
	if shareCount <= 0 {
		return fmt.Errorf("share count must be positive")
	}

	// Validate users exist (assuming you have a UserValidator)
	if !v.IsShareHolderExists(fromUserID, marketChannelID) {
		return fmt.Errorf("from user does not exist")
	}
	if !v.IsShareHolderExists(toUserID, marketChannelID) {
		return fmt.Errorf("to user does not exist")
	}

	// Validate users are not the same
	if fromUserID == toUserID {
		return fmt.Errorf("cannot transfer shares to the same user")
	}

	return nil
}

func (v *shareHolderValidators) ValidateTransferShares(shareCount, fromShareHolderShareCount, totalShares int) error {
	if shareCount > totalShares {
		return fmt.Errorf("share count exceeds total shares")
	}
	if shareCount > fromShareHolderShareCount {
		return fmt.Errorf("share count exceeds from user's share count")
	}

	return nil
}
