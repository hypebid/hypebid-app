package models

type TwitchUser struct {
	ID string `json:"id" gorm:"primaryKey"`
	// UserID          uuid.UUID `gorm:"unique;type:uuid"` // Foreign key to User
	Login           string `json:"login" gorm:"unique"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
	CreatedAt       string `json:"created_at"`
	// Add OAuth-related fields
	// AccessToken    string `gorm:"type:text"`
	// RefreshToken   string `gorm:"type:text"`
	// TokenExpiresAt time.Time
}
