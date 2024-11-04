package models

import (
	"time"

	"github.com/google/uuid"
)

// MarketInstance represents a market instance
// You can create a new instance with the CreateInstanceHandler
// A new instance doesn't need to be started, it will be started when the auction starts or started manually by ad admin

type MarketInstance struct {
	InstanceID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	Name           string    `gorm:"not null"`
	Status         string    `gorm:"type:market_status;default:'not_started'"`
	DurationDays   int
	CurrentPhase   string
	Channels       []Channel `gorm:"type:jsonb"`
	StartTimestamp time.Time
	EndTimestamp   time.Time
	Metrics        []Metric `gorm:"type:jsonb"`
}
