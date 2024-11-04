package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hypebid/hypebid-app/internal/services"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CreateMarketChannelRequest struct {
	ChannelId   uuid.UUID `json:"channelId"`
	TotalShares int       `json:"totalShares"`
	SharePrice  float64   `json:"sharePrice"`
}

type MarketChannelResponse struct {
	MarketChannel *models.MarketChannel `json:"marketChannel"`
}

type MarketChannelsResponse struct {
	MarketChannels []models.MarketChannel `json:"marketChannels"`
}

type CreateShareHolderRequest struct {
	UserID uuid.UUID `json:"userId"`
}

type ShareHolderResponse struct {
	Message     string              `json:"message"`
	ShareHolder *models.ShareHolder `json:"shareHolder"`
}

func CreateMarketChannelHandler(marketChannelService services.MarketChannelService, marketInstanceService services.MarketInstanceService, channelService services.ChannelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		instanceId := chi.URLParam(r, "instanceId")
		if instanceId == "" {
			http.Error(w, "Instance ID is required", http.StatusBadRequest)
			return
		}

		instanceIdUUID, err := uuid.Parse(instanceId)
		if err != nil {
			http.Error(w, "Invalid instance ID", http.StatusBadRequest)
			return
		}

		var req CreateMarketChannelRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		marketChannel, err := marketChannelService.InitializeMarketChannel(instanceIdUUID, req.ChannelId, req.TotalShares, req.SharePrice)
		if err != nil {
			http.Error(w, "Failed to create market channel", http.StatusInternalServerError)
			return
		}

		// Create MarketChannelResponse
		response := MarketChannelResponse{
			MarketChannel: marketChannel,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func GetMarketChannelsHandler(marketChannelService services.MarketChannelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instanceId := chi.URLParam(r, "instanceId")
		if instanceId == "" {
			http.Error(w, "Instance ID is required", http.StatusBadRequest)
			return
		}

		instanceIdUUID, err := uuid.Parse(instanceId)
		if err != nil {
			http.Error(w, "Invalid instance ID", http.StatusBadRequest)
			return
		}

		marketChannels, err := marketChannelService.GetMarketChannelsByInstanceID(instanceIdUUID)
		if err != nil {
			http.Error(w, "Failed to get market channels", http.StatusInternalServerError)
			return
		}

		response := MarketChannelsResponse{
			MarketChannels: marketChannels,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func CreateShareHolderHandler(shareHolderService services.ShareHolderService, marketChannelService services.MarketChannelService, userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		instanceId := chi.URLParam(r, "instanceId")
		if instanceId == "" {
			http.Error(w, "Instance ID is required", http.StatusBadRequest)
			return
		}

		_, err := uuid.Parse(instanceId)
		if err != nil {
			http.Error(w, "Invalid instance ID", http.StatusBadRequest)
			return
		}

		marketChannelId := chi.URLParam(r, "marketChannelId")
		if marketChannelId == "" {
			http.Error(w, "Market channel ID is required", http.StatusBadRequest)
			return
		}

		marketChannelIdUUID, err := uuid.Parse(marketChannelId)
		if err != nil {
			http.Error(w, "Invalid market channel ID", http.StatusBadRequest)
			return
		}

		var req CreateShareHolderRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		shareHolder, err := shareHolderService.InitializeShareHolder(req.UserID, marketChannelIdUUID)
		if err != nil {
			http.Error(w, "Failed to create share holder", http.StatusInternalServerError)
			return
		}

		response := ShareHolderResponse{
			Message:     "ShareHolder created successfully",
			ShareHolder: shareHolder,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
