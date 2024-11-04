package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/services"
	"github.com/hypebid/hypebid-app/internal/validator"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CreateAuctionRequest struct {
	ChannelID string `json:"channelId"`
}

type CreateAuctionResponse struct {
	Message string          `json:"message"`
	Auction *models.Auction `json:"auction"`
}

type StartAuctionRequest struct {
	Duration time.Duration `json:"duration"`
}

type StartAuctionResponse struct {
	Message string          `json:"message"`
	Auction *models.Auction `json:"auction"`
}

type GetAuctionResponse struct {
	Auction *models.Auction `json:"auction"`
}

type PlaceBidRequest struct {
	UserID uuid.UUID `json:"userId"`
	Amount float64   `json:"amount"`
}

func CreateAuctionHandler(auctionService services.AuctionService, auctionValidators validator.AuctionValidators) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		instanceId := chi.URLParam(r, "instanceId")
		err := auctionValidators.ValidateInstanceID(instanceId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req CreateAuctionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate request
		// Need to refactor to use request validator and handler constructors
		if err := validateCreateAuctionRequest(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = auctionValidators.ValidateChannelID(req.ChannelID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create auction
		auction, err := auctionService.CreateAuction(uuid.MustParse(instanceId), uuid.MustParse(req.ChannelID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		createAuctionResponse := CreateAuctionResponse{
			Message: "Auction created",
			Auction: auction,
		}

		// Return success response
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createAuctionResponse)
	}
}

func GetAllAuctionsHandler(auctionService services.AuctionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement auction getting logic
	}
}

func validateCreateAuctionRequest(req CreateAuctionRequest) error {
	// Validate market instance ID and channel ID
	if req.ChannelID == "" {
		return errors.New("channel ID is required")
	}

	return nil
}

func GetCurrentAuctionHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement current auction getting logic
	}
}

func StartAuctionHandler(auctionService services.AuctionService, auctionValidators validator.AuctionValidators) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		instanceId := chi.URLParam(r, "instanceId")
		err := auctionValidators.ValidateInstanceID(instanceId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		auctionId := chi.URLParam(r, "auctionId")
		err = auctionValidators.ValidateAuctionID(auctionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req StartAuctionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate StartAuctionRequest
		if err := validateStartAuctionRequest(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Start auction
		auction, err := auctionService.StartAuction(uuid.MustParse(auctionId), req.Duration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		startAuctionResponse := StartAuctionResponse{
			Message: "Auction started",
			Auction: auction,
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(startAuctionResponse)
	}
}

func validateStartAuctionRequest(req StartAuctionRequest) error {
	if req.Duration <= 0 {
		return errors.New("duration must be greater than 0")
	}
	return nil
}

func PlaceBidHandler(auctionService services.AuctionService, auctionValidators validator.AuctionValidators, userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request
		instanceID := chi.URLParam(r, "instanceId")
		auctionID := chi.URLParam(r, "auctionId")

		err := auctionValidators.ValidateInstanceID(instanceID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = auctionValidators.ValidateAuctionID(auctionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req PlaceBidRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate user exists and has sufficient funds
		err = auctionValidators.ValidateUserForBid(req.UserID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Place bid (auction service handles all other validations)
		err = auctionService.PlaceBid(uuid.MustParse(instanceID), uuid.MustParse(auctionID), req.UserID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := struct {
			Message string  `json:"message"`
			Amount  float64 `json:"amount"`
		}{
			Message: "Bid placed successfully",
			Amount:  req.Amount,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func GetAuctionHandler(auctionService services.AuctionService, auctionValidators validator.AuctionValidators) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for GET method
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		instanceId := chi.URLParam(r, "instanceId")
		err := auctionValidators.ValidateInstanceID(instanceId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		auctionId := chi.URLParam(r, "auctionId")
		err = auctionValidators.ValidateAuctionID(auctionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get auction
		auction, err := auctionService.GetAuctionByID(uuid.MustParse(auctionId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		getAuctionResponse := GetAuctionResponse{
			Auction: auction,
		}

		// Return success response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getAuctionResponse)
	}
}

func GetAuctionHistoryHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement auction history getting logic
	}
}
