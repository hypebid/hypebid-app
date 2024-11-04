package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/twitch"
	"github.com/hypebid/hypebid-app/pkg/models"
)

var (
	fileMutex sync.Mutex
	dataFile  = "follower_counts.json"
)

func FollowersHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		broadcasterID := r.URL.Query().Get("broadcaster_id")
		login := r.URL.Query().Get("login")
		var userID = ""

		if broadcasterID == "" && login == "" {
			http.Error(w, "Missing broadcaster_id or login parameter", http.StatusBadRequest)
			return
		}

		fmt.Println("broadcasterID: ", broadcasterID)

		client := twitch.NewClient(cfg)
		accessToken, err := client.GetAccessToken()
		if err != nil {
			http.Error(w, "Error getting access token", http.StatusInternalServerError)
			return
		}

		if broadcasterID == "" {
			userID, err = client.GetUserByLogin(accessToken, login)
			if err != nil {
				http.Error(w, "Error getting userID for login", http.StatusInternalServerError)
			}
		}

		// BAD LOGIC FIX THIS
		if broadcasterID == "" && userID != "" {
			fmt.Printf("ID found for login %s, setting broadcasterId to %s ", login, userID)
			broadcasterID = userID
		}

		followerCount, err := client.GetFollowerCount(accessToken, broadcasterID)
		if err != nil {
			http.Error(w, "Error getting follower count", http.StatusInternalServerError)
			return
		}

		response := map[string]int{"follower_count": followerCount}
		fmt.Println(response)
		json.NewEncoder(w).Encode(response)
	}
}

func BulkFollowersHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logins := r.URL.Query()["login"]
		if len(logins) == 0 || len(logins) > 100 {
			http.Error(w, "Please provide between 1 and 100 login parameters", http.StatusBadRequest)
			return
		}

		client := twitch.NewClient(cfg)
		accessToken, err := client.GetAccessToken()
		if err != nil {
			http.Error(w, "Error getting access token", http.StatusInternalServerError)
			return
		}

		users, err := client.GetUsersByLogin(accessToken, logins)
		if err != nil {
			http.Error(w, "Error getting users", http.StatusInternalServerError)
			return
		}

		results := make(map[string]int)
		var followerCounts []models.FollowerCount
		now := time.Now()

		for _, user := range users {
			followerCount, err := client.GetFollowerCount(accessToken, user.ID)
			if err != nil {
				results[user.Login] = -1 // Indicate error for this user
				continue
			}
			results[user.Login] = followerCount
			followerCounts = append(followerCounts, models.FollowerCount{
				UserLogin:     user.Login,
				FollowerCount: followerCount,
				RecordedAt:    now,
			})
		}
		// Write to JSON file
		go writeFollowerCountsToFile(followerCounts)

		json.NewEncoder(w).Encode(results)
	}
}

func writeFollowerCountsToFile(newCounts []models.FollowerCount) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	var allCounts []models.FollowerCount

	// Read existing data
	data, err := os.ReadFile(dataFile)
	if err == nil {
		json.Unmarshal(data, &allCounts)
	}

	// Append new data
	allCounts = append(allCounts, newCounts...)

	// Write all data back to file
	data, err = json.MarshalIndent(allCounts, "", "  ")
	if err != nil {
		// Log the error (consider using a proper logging package)
		println("Error marshalling data:", err)
		return
	}

	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		// Log the error
		println("Error writing to file:", err)
	}
}
