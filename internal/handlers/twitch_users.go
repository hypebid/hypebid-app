package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/services"
	"github.com/hypebid/hypebid-app/internal/twitch"
	"github.com/hypebid/hypebid-app/pkg/models"
)

//ID:    twitchUser.ID,
//Login: twitchUser.Login,
//Email: twitchUser.Email,

type TwitchUserPostRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
	ID    string `json:"id"`
}

type TwitchUserResponse struct {
	Message string       `json:"message"`
	User    *models.User `json:"user"`
}

func TwitchUsersHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := r.URL.Query().Get("login")

		client := twitch.NewClient(cfg)
		accessToken, err := client.GetAccessToken()
		if err != nil {
			http.Error(w, "Error getting access token", http.StatusInternalServerError)
			return
		}

		userID, err := client.GetUserByLogin(accessToken, login)
		if err != nil {
			http.Error(w, "Error getting userID for login", http.StatusInternalServerError)
		}

		response := map[string]string{"user_id": userID}
		json.NewEncoder(w).Encode(response)
	}
}

func TwitchUserHandler(cfg *config.Config, userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check Post Request
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request TwitchUserPostRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		twitchUser := &models.TwitchUser{
			ID:    request.ID,
			Login: request.Login,
			Email: request.Email,
		}

		user, err := userService.FindOrCreateTwitchUser(twitchUser, nil)
		if err != nil {
			http.Error(w, "Error finding or creating Twitch user", http.StatusInternalServerError)
			return
		}

		response := TwitchUserResponse{
			Message: "User created successfully",
			User:    user,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func ProtectedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
