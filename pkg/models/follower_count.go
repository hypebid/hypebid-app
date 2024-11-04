package models

import "time"

type FollowerCount struct {
	UserLogin     string    `json:"user_login"`
	FollowerCount int       `json:"follower_count"`
	RecordedAt    time.Time `json:"recorded_at"`
}
