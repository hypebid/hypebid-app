package models

import (
	"time"

	"github.com/google/uuid"
)

type Auction struct {
	AuctionID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MarketChannelID  uuid.UUID `gorm:"foreignKey:ID"`
	HighestBid       float64   `gorm:"type:numeric(10,2)"`
	HighestBidderID  uuid.UUID
	Status           string    `gorm:"type:auction_status;default:'open'"`
	MarketInstanceID uuid.UUID `gorm:"foreignKey:MarketInstanceID"`
	StartTime        time.Time
	EndTime          time.Time
}
