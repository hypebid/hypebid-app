package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/twitch"
	"github.com/hypebid/hypebid-app/pkg/models"
	"gorm.io/gorm"
)

var (
	fileMutex sync.Mutex
	dataFile  = "follower_counts.json"
)

func StartFollowerUpdater(cfg *config.Config, db *gorm.DB, logins []string) {
	interval, err := strconv.Atoi(cfg.FollowerUpdateInterval)
	fmt.Println("Follower update interval:", interval)
	if err != nil {
		log.Fatalf("Invalid follower update interval: %v", err)
	}
	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	go func() {
		for range ticker.C {
			updateFollowerCounts(cfg, db, logins)
		}
	}()
}

// Convert follower counts to metric entries in a batch
func updateFollowerCountsMetricBatch(db *gorm.DB, metricDataPoints []models.MetricDataPoint) error {
	// Create a slice to hold all metric entries
	metrics := make([]models.Metric, len(metricDataPoints))

	// Convert follower counts to metric entries
	for i, mdp := range metricDataPoints {
		metrics[i] = models.Metric{
			Name:       "follower_count",
			MetricID:   mdp.MetricID,
			DataPoints: []models.MetricDataPoint{mdp},
		}
	}

	// Create all records in a single transaction
	return db.CreateInBatches(metrics, 100).Error
}

func updateFollowerCounts(cfg *config.Config, db *gorm.DB, logins []string) {
	client := twitch.NewClient(cfg)
	accessToken, err := client.GetAccessToken()
	if err != nil {
		log.Printf("Error getting access token: %v", err)
		return
	}

	users, err := client.GetUsersByLogin(accessToken, logins)
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return
	}

	now := time.Now()

	// Use a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, user := range users {
		// Create or get channel
		var channel models.Channel
		if err := tx.FirstOrCreate(&channel, models.Channel{
			Name: user.Login,
		}).Error; err != nil {
			tx.Rollback()
			log.Printf("Error creating channel for %s: %v", user.Login, err)
			return
		}

		// Create or get metric
		var metric models.Metric
		if err := tx.FirstOrCreate(&metric, models.Metric{
			ChannelID: channel.ChannelID,
			Name:      "follower_count",
		}).Error; err != nil {
			tx.Rollback()
			log.Printf("Error creating metric for %s: %v", user.Login, err)
			return
		}

		// Get follower count from Twitch
		followerCount, err := client.GetFollowerCount(accessToken, user.ID)
		if err != nil {
			log.Printf("Error getting follower count for %s: %v", user.Login, err)
			continue
		}

		// Create data point
		dataPoint := models.MetricDataPoint{
			MetricID:   metric.MetricID,
			Value:      int(followerCount),
			RecordedAt: now,
		}

		if err := tx.Create(&dataPoint).Error; err != nil {
			tx.Rollback()
			log.Printf("Error creating data point for %s: %v", user.Login, err)
			return
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
	}
}

func writeFollowerCountsToFile(newCounts []models.FollowerCount) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	var allCounts []models.FollowerCount

	data, err := os.ReadFile(dataFile)
	if err == nil {
		json.Unmarshal(data, &allCounts)
	}

	allCounts = append(allCounts, newCounts...)

	data, err = json.MarshalIndent(allCounts, "", "  ")
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return
	}

	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}
