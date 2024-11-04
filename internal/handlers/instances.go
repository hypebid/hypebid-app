package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/services"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var maxUsers = 10
var maxChannels = 20
var maxChannelShares = 100
var coefficient = 1.04
var initialMetric = "followers"

type CreateInstanceRequest struct {
	Name         string `json:"name"`
	DurationDays int    `json:"durationDays"`
	Email        string `json:"email"`
}

type CreateInstanceResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	DurationDays int    `json:"durationDays"`
}

type GetInstancesResponse struct {
	Instances []models.MarketInstance `json:"instances"`
}

type AddUserToInstanceRequest struct {
	MarketInstanceID string `json:"instanceId"`
	UserID           string `json:"userId"`
}

type StartInstanceResponse struct {
	Message  string                `json:"message"`
	Instance models.MarketInstance `json:"instance"`
}

// CreateInstanceHandler creates a new market instance
func CreateInstanceHandler(marketInstanceService services.MarketInstanceService, userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateInstanceRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := userService.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validate the request
		if err := validateCreateInstanceRequest(req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		instance, err := marketInstanceService.CreateMarketInstance(req.Name, req.DurationDays, req.Email, user.UserID)
		if err != nil {
			if err.Error() == "you've already used that name for a market" {
				http.Error(w, err.Error(), http.StatusBadRequest) // Return 400 for existing instance name
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateInstanceResponse{
			Name:         instance.Name,
			Email:        req.Email,
			Status:       string(instance.Status),
			DurationDays: instance.DurationDays,
		})
	}
}

func validateCreateInstanceRequest(req CreateInstanceRequest) error {
	if req.Name == "" || req.DurationDays <= 0 || req.Email == "" {
		return errors.New("please provide a name, duration, and email")
	}

	return nil
}

func GetAllInstancesHandler(marketInstanceService services.MarketInstanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instances, err := marketInstanceService.GetAllMarketInstances()
		if err != nil {
			log.Printf("Error retrieving market instances: %v", err) // Log the error
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GetInstancesResponse{Instances: instances})
	}
}

func AddUserToInstanceHandler(marketInstanceService services.MarketInstanceService, userService services.UserService, memberService services.MemberService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req AddUserToInstanceRequest
		instanceId := chi.URLParam(r, "instanceId")

		// Decode the request body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the request
		if err := validateAddUserToInstanceRequest(instanceId, req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Parse UUIDs safely
		instanceUUID, _ := uuid.Parse(instanceId)
		userUUID, _ := uuid.Parse(req.UserID)

		// Create member
		if err := memberService.CreateMember(instanceUUID, userUUID); err != nil {
			http.Error(w, "Failed to add user to instance", http.StatusInternalServerError)
			return
		}

		// Respond with success
		w.WriteHeader(http.StatusCreated) // Set status to 201 Created
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "User added to instance"})
	}
}

// validateAddUserToInstanceRequest validates the input for adding a user to an instance.
func validateAddUserToInstanceRequest(instanceId string, req AddUserToInstanceRequest) error {
	if instanceId == "" {
		return fmt.Errorf("instance ID is required")
	}
	if req.UserID == "" {
		return fmt.Errorf("user ID is required")
	}
	if _, err := uuid.Parse(instanceId); err != nil {
		return fmt.Errorf("invalid instance ID format")
	}
	if _, err := uuid.Parse(req.UserID); err != nil {
		return fmt.Errorf("invalid user ID format")
	}
	return nil
}

func GetAllMembersForInstanceHandler(memberService services.MemberService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instanceId := chi.URLParam(r, "instanceId")
		instanceUUID, _ := uuid.Parse(instanceId)

		members, err := memberService.GetAllMembersForInstance(instanceUUID)
		if err != nil {
			http.Error(w, "Failed to get members for instance", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(members)
	}
}

func StartInstanceHandler(marketInstanceService services.MarketInstanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instanceId := chi.URLParam(r, "instanceId")
		instanceUUID, _ := uuid.Parse(instanceId)

		fmt.Println("Starting instance with ID:", instanceUUID)

		instance, err := marketInstanceService.StartMarketInstance(instanceUUID)
		if err != nil {
			http.Error(w, "Failed to start instance", http.StatusInternalServerError)
			return
		}

		fmt.Println("Instance started:", instance)
		// Create the response
		response := StartInstanceResponse{
			Message:  "Instance started",
			Instance: *instance,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // Set appropriate status code
		json.NewEncoder(w).Encode(response)
	}
}

func GetInstanceHandler(marketInstanceService services.MarketInstanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		instanceId := chi.URLParam(r, "instanceId")
		instanceUUID, _ := uuid.Parse(instanceId)

		instance, err := marketInstanceService.GetMarketInstanceByID(instanceUUID)
		if err != nil {
			http.Error(w, "Failed to get instance", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(instance)
	}
}

func GetChannelsHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement channels getting logic
	}
}

func GetChannelHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement channel getting logic
	}
}

func CompleteInstanceHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement instance completion logic
	}
}

func GetInstanceResultsHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement instance results getting logic
	}
}
