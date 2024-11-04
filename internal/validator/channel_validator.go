package validator

import (
	"errors"

	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

var _ ChannelValidator = (*channelValidator)(nil)

type channelValidator struct {
	channelRepo repositories.ChannelRepository
}

func NewChannelValidator(channelRepo repositories.ChannelRepository) *channelValidator {
	return &channelValidator{channelRepo: channelRepo}
}

func (v *channelValidator) IsChannelExists(channelID uuid.UUID) bool {
	_, err := v.channelRepo.GetChannelByID(channelID)
	return err == nil
}

func (v *channelValidator) GetChannelByID(channelID uuid.UUID) (*models.Channel, error) {
	return v.channelRepo.GetChannelByID(channelID)
}

func (v *channelValidator) ValidateChannelID(channelID string) error {
	channelUUID, err := uuid.Parse(channelID)
	if err != nil {
		return errors.New("invalid channel ID")
	}

	if !v.IsChannelExists(channelUUID) {
		return errors.New("channel does not exist")
	}

	return nil
}
