package models

import (
	"time"

	"github.com/google/uuid"
)

type Trade struct {
	TradeID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MarketInstanceID uuid.UUID `gorm:"foreignKey:MarketInstanceID"`
	InitiatorID      uuid.UUID `gorm:"foreignKey:UserID"`
	RecipientID      uuid.UUID `gorm:"foreignKey:UserID"`
	Status           string    `gorm:"type:trade_status;default:'pending'"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	CompletedAt      time.Time
	BoostID          *uuid.UUID
}

type TradeItem struct {
	TradeItemID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TradeID          uuid.UUID `gorm:"foreignKey:TradeID"`
	MarketChannelID  uuid.UUID `gorm:"foreignKey:MarketChannelID"`
	MarketInstanceID uuid.UUID `gorm:"foreignKey:MarketInstanceID"`
	ShareCount       int
	Currency         float64 `gorm:"type:numeric(14,2)"`
	Direction        string  `gorm:"type:trade_direction"` // 'offer' or 'request'
}
