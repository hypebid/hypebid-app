package models

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	ChannelID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"not null"`
	Metrics   []Metric  `gorm:"foreignKey:ChannelID"`
}

type Metric struct {
	MetricID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ChannelID  uuid.UUID `gorm:"type:uuid"`
	Name       string
	DataPoints []MetricDataPoint `gorm:"foreignKey:MetricID"`
}

type MetricDataPoint struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	MetricID   uuid.UUID `gorm:"type:uuid"`
	Value      int
	RecordedAt time.Time `gorm:"not null;index"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
