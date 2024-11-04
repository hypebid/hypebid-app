package models

import (
	"time"

	"github.com/google/uuid"
)

type Boost struct {
	BoostID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID           uuid.UUID `gorm:"foreignKey:UserID"`
	Type             string    `gorm:"type:boost_type"`
	IsActive         bool      `gorm:"default:true"`
	ExpiresAt        time.Time
	TradeID          *uuid.UUID
	MarketInstanceID uuid.UUID `gorm:"foreignKey:MarketInstanceID"`
}
