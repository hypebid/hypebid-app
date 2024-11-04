// internal/validator/interfaces.go
package validator

import (
	"github.com/google/uuid"
)

type AuctionValidator interface {
	ValidateAuctionID(auctionID string) error
	IsAuctionExists(auctionID uuid.UUID) bool
}

type ChannelValidator interface {
	ValidateChannelID(channelID string) error
	IsChannelExists(channelID uuid.UUID) bool
}

type MarketInstanceValidator interface {
	IsMarketInstanceExists(instanceID uuid.UUID) bool
	ValidateInstanceID(instanceID string) error
}

type ShareHolderValidator interface {
	IsShareHolderExists(userID uuid.UUID, marketChannelID uuid.UUID) bool
}

type MarketChannelValidator interface {
	IsMarketChannelExists(marketChannelID uuid.UUID) bool
}

type AuctionValidators interface {
	IsAuctionExists(auctionID uuid.UUID) bool
	ValidateAuctionID(auctionID string) error
	ValidateChannelID(channelID string) error
	ValidateInstanceID(instanceID string) error
	ValidateUserForBid(userID uuid.UUID, amount float64) error
}

// You might also want a combined interface for external use
type MarketChannelValidators interface {
	IsMarketChannelExists(marketChannelID uuid.UUID) bool
	IsChannelExists(channelID uuid.UUID) bool
	IsShareHolderExists(userID uuid.UUID, marketChannelID uuid.UUID) bool
	IsMarketInstanceExists(instanceID uuid.UUID) bool
}

type ShareHolderValidators interface {
	IsShareHolderExists(userID uuid.UUID, marketChannelID uuid.UUID) bool
	ValidateTransferSharesInput(marketChannelID, fromUserID, toUserID uuid.UUID, shareCount int) error
	ValidateTransferShares(shareCount, fromShareHolderShareCount, totalShares int) error
}

type UserValidator interface {
	IsUserExists(userID uuid.UUID) bool
	ValidateUserForBid(userID uuid.UUID, amount float64) error
	ValidateUserBalance(userID uuid.UUID, amount float64) bool
}
