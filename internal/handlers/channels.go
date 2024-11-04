package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/internal/services"
)

type CreateChannelRequest struct {
	Name        string `json:"name"`
	SharesTotal int    `json:"sharesTotal"`
}

type RecentFollowerCountResponse struct {
	FollowerCounts []FollowerCountResponse `json:"followerCounts"`
}

type FollowerCountResponse struct {
	RecordedAt    time.Time `json:"recordedAt"`
	FollowerCount int       `json:"followerCount"`
}

type PeriodInfo struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Days  int    `json:"days"`
}

type MetaInfo struct {
	TotalDays     int `json:"totalDays"`
	MaxFollowers  int `json:"maxFollowers"`
	MinFollowers  int `json:"minFollowers"`
	OverallGrowth int `json:"overallGrowth"`
}

type FollowerStatsResponse struct {
	Data struct {
		Channel string                     `json:"channel"`
		Period  PeriodInfo                 `json:"period"`
		Metrics []repositories.DailyMetric `json:"metrics"`
	} `json:"data"`
	Meta MetaInfo `json:"meta"`
}

func RegisterChannelHandler(channelService services.ChannelService, twitchClient services.TwitchService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req CreateChannelRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the request
		if err := validateCreateChannelRequest(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Verify channel name is not already taken
		if channelService.IsChannelNameTaken(req.Name) {
			http.Error(w, "Channel name already taken", http.StatusBadRequest)
			return
		}

		// Verify channel shares total is greater than 0
		if req.SharesTotal <= 0 {
			http.Error(w, "Shares total must be greater than 0", http.StatusBadRequest)
			return
		}

		// Verify channel exists on Twitch
		accessToken, err := twitchClient.GetAccessToken()
		if err != nil {
			http.Error(w, "Failed to get access token", http.StatusInternalServerError)
			return
		}

		twitchUserID, err := twitchClient.GetUserByLogin(accessToken, req.Name)
		if err != nil {
			http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
			return
		}

		if twitchUserID == "" {
			http.Error(w, "User not found on Twitch", http.StatusBadRequest)
			return
		}

		if err := channelService.CreateChannel(req.Name, req.SharesTotal); err != nil {
			http.Error(w, "Failed to create channel", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Channel created"})
	}
}

func GetAllChannelsHandler(channelService services.ChannelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for GET method
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		channels, err := channelService.GetAllChannels()
		if err != nil {
			http.Error(w, "Failed to get channels", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(channels)
	}
}

func validateCreateChannelRequest(req CreateChannelRequest) error {
	if req.Name == "" || req.SharesTotal <= 0 {
		return errors.New("invalid request payload")
	}
	return nil
}

func RecentFollowerCountHandler(metricService services.MetricService, metricDataPointService services.MetricDataPointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for GET method
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		channelName := chi.URLParam(r, "login")
		// normalize channel name
		channelName, err := NormalizeURLParam(channelName)
		if err != nil {
			http.Error(w, "Invalid channel name", http.StatusBadRequest)
			return
		}

		metricDataPoints, err := metricDataPointService.GetRecentFollowerCount(channelName)
		if err != nil {
			http.Error(w, "Failed to get follower count", http.StatusInternalServerError)
			return
		}

		var followerCountsResponse RecentFollowerCountResponse
		for _, metricDataPoint := range metricDataPoints {
			followerCountsResponse.FollowerCounts = append(followerCountsResponse.FollowerCounts, FollowerCountResponse{
				RecordedAt:    metricDataPoint.RecordedAt,
				FollowerCount: metricDataPoint.Value,
			})
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(followerCountsResponse)
	}
}

func AverageFollowerCountHandler(metricDataPointService services.MetricDataPointService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		channelName := chi.URLParam(r, "login")
		daysStr := chi.URLParam(r, "days")

		// Normalize channel name
		channelName, err := NormalizeURLParam(channelName)
		if err != nil {
			http.Error(w, "Invalid channel name", http.StatusBadRequest)
			return
		}

		// Parse and validate days parameter
		days, err := strconv.Atoi(daysStr)
		if err != nil || days <= 0 || days > 365 {
			http.Error(w, "Invalid days parameter (must be between 1 and 365)", http.StatusBadRequest)
			return
		}

		stats, err := metricDataPointService.GetFollowerStats(channelName, days)
		if err != nil {
			http.Error(w, "Failed to get follower statistics", http.StatusInternalServerError)
			return
		}

		var response FollowerStatsResponse

		response.Data.Channel = channelName
		response.Data.Period = PeriodInfo{
			Start: stats.StartTime.Format(time.RFC3339),
			End:   stats.EndTime.Format(time.RFC3339),
			Days:  days,
		}
		response.Data.Metrics = stats.DailyMetrics
		response.Meta = MetaInfo{
			TotalDays:     days,
			MaxFollowers:  stats.MaxFollowers,
			MinFollowers:  stats.MinFollowers,
			OverallGrowth: stats.OverallGrowth,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
