package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegistrationResponse struct {
	UserID   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Token    string    `json:"token"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserID   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Token    string    `json:"token"`
}

type AddCurrencyRequest struct {
	Amount float64 `json:"amount"`
}

type AddCurrencyResponse struct {
	Message  string  `json:"message"`
	Currency float64 `json:"currency"`
}

func RegisterUserHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UserRegistrationRequest

		fmt.Println("request body ", r.Body)

		// Decode the request body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := validateUserRegistrationRequest(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Save the user to the database
		user, err := userService.CreateUser(req.Username, req.Email, req.Password)
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Generate a JWT token
		// token, err := generateToken(user.UserID)
		// if err != nil {
		// 	http.Error(w, "Error generating token", http.StatusInternalServerError)
		// 	return
		// }

		response := UserRegistrationResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
			// Token:    token,
		}

		// Set response header and send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}

func validateUserRegistrationRequest(req UserRegistrationRequest) error {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return errors.New("username, email, and password are required")
	}

	// Validate email format
	if !strings.Contains(req.Email, "@") {
		return errors.New("invalid email format")
	}

	return nil
}

func LoginUserHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UserLoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := userService.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// token, err := auth.GenerateToken(user.UserID.String())
		// if err != nil {
		// 	http.Error(w, "Error generating token", http.StatusInternalServerError)
		// 	return
		// }

		response := UserLoginResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Email:    user.Email,
			// Token:    token,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func LogoutUserHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement user logout logic
	}
}

func GetUserInstancesHandler(marketInstanceService services.MarketInstanceService, userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request URL:", r.URL.String())
		userId := chi.URLParam(r, "userId")
		fmt.Println("userId: ", userId)

		userUUID, err := uuid.Parse(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := userService.GetUserByID(userUUID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("user: ", user)

		instances, err := marketInstanceService.GetAllMarketInstancesByUserID(user.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("instances: ", instances)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GetInstancesResponse{Instances: instances})
	}
}

func AddCurrencyHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for POST method
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userId := chi.URLParam(r, "userId")
		userUUID, err := uuid.Parse(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req AddCurrencyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := userService.AddCurrency(userUUID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add Currency Response
		var response AddCurrencyResponse
		response.Message = "Currency added successfully"
		response.Currency = user.Currency

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
