// internal/validator/market_channel.go
package validator

import (
	"github.com/hypebid/hypebid-app/internal/repositories"

	"github.com/google/uuid"
)

var _ MarketChannelValidator = (*marketChannelValidator)(nil)

type marketChannelValidator struct {
	marketChannelRepo repositories.MarketChannelRepository
}

func NewMarketChannelValidator(marketChannelRepo repositories.MarketChannelRepository) *marketChannelValidator {
	return &marketChannelValidator{marketChannelRepo: marketChannelRepo}
}

type marketChannelValidators struct {
	channel        ChannelValidator
	marketInstance MarketInstanceValidator
	shareHolder    ShareHolderValidator
	marketChannel  *marketChannelValidator
}

func NewMarketChannelValidators(
	marketChannelRepo repositories.MarketChannelRepository,
	channelValidator ChannelValidator,
	instanceValidator MarketInstanceValidator,
	shareHolderValidator ShareHolderValidator,
) MarketChannelValidators {
	return &marketChannelValidators{
		channel:        channelValidator,
		marketInstance: instanceValidator,
		shareHolder:    shareHolderValidator,
		marketChannel:  NewMarketChannelValidator(marketChannelRepo),
	}
}

// Delegate to underlying validators
func (v *marketChannelValidators) IsChannelExists(channelID uuid.UUID) bool {
	return v.channel.IsChannelExists(channelID)
}

func (v *marketChannelValidators) IsMarketInstanceExists(instanceID uuid.UUID) bool {
	return v.marketInstance.IsMarketInstanceExists(instanceID)
}

func (v *marketChannelValidators) IsShareHolderExists(userID uuid.UUID, marketChannelID uuid.UUID) bool {
	return v.shareHolder.IsShareHolderExists(userID, marketChannelID)
}

func (v *marketChannelValidator) IsMarketChannelExists(marketChannelID uuid.UUID) bool {
	_, err := v.marketChannelRepo.GetMarketChannelByID(marketChannelID)
	return err == nil
}

func (v *marketChannelValidators) IsMarketChannelExists(marketChannelID uuid.UUID) bool {
	return v.marketChannel.IsMarketChannelExists(marketChannelID)
}
