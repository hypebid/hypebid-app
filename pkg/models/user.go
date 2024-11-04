package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string    `gorm:"unique;not null"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash *string
	TwitchID     *string     `gorm:"unique;"`
	Currency     float64     `gorm:"type:numeric(10,2);default:0"`
	Shares       []UserShare `gorm:"foreignKey:UserID"`
	Boosts       []Boost     `gorm:"foreignKey:UserID"`
	TwitchUser   TwitchUser  `gorm:"foreignKey:TwitchID"`
	// Auth Fields
	AccessToken    string `gorm:"type:text"`
	RefreshToken   string `gorm:"type:text"`
	TokenExpiresAt time.Time
	LastLoginAt    time.Time
}

type Member struct {
	MemberID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID           uuid.UUID `gorm:"foreignKey:UserID;primaryKey"`
	MarketInstanceID uuid.UUID `gorm:"foreignKey:MarketInstanceID;primaryKey"`
	Role             string
	JoinedAt         time.Time `gorm:"autoCreateTime"` // Automatically set when the record is created
	Status           string    `gorm:"default:'active'"`
	Currency         float64   `gorm:"type:numeric(14,2);default:0"`
}

type UserShare struct {
	UserShareID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID           uuid.UUID `gorm:"foreignKey:UserID;primaryKey"`
	ChannelID        uuid.UUID `gorm:"foreignKey:ChannelID;primaryKey"`
	MarketInstanceID uuid.UUID `gorm:"foreignKey:MarketInstanceID;primaryKey"`
	ShareCount       int
}
