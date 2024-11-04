package models

import (
	"time"

	"github.com/google/uuid"
)

type MarketChannel struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MarketInstanceID uuid.UUID `gorm:"foreignKey:MarketInstanceID;primaryKey"`
	ChannelID        uuid.UUID `gorm:"foreignKey:ChannelID;primaryKey"`
	TotalShares      int       `gorm:"not null;default:0"`
	SharePrice       float64   `gorm:"not null;default:0"`
	// Add metrics for the channel
}

type ShareHolder struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MarketChannelID uuid.UUID `gorm:"foreignKey:MarketChannelID;primaryKey"`
	UserID          uuid.UUID `gorm:"foreignKey:UserID;primaryKey"`
	ShareCount      int       `gorm:"not null;default:0"`
	AcquisitionDate time.Time `gorm:"autoCreateTime"`
}
