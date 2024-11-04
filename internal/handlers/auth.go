package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/hypebid/hypebid-app/internal/auth"
	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/services"
	"github.com/hypebid/hypebid-app/pkg/models"
)

type AuthHandler struct {
	oauthManager auth.OAuthManager
	userService  services.UserService
	cfg          *config.Config
}

func NewAuthHandler(cfg *config.Config, oauthManager auth.OAuthManager, userService services.UserService) *AuthHandler {
	return &AuthHandler{
		cfg:          cfg,
		oauthManager: oauthManager,
		userService:  userService,
	}
}

func (h *AuthHandler) TwitchLoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if user is already logged in
		// user, err := h.userService.GetUserByID()
		// if err == nil && user != nil {
		// 	http.Redirect(w, r, h.cfg.FrontendURL, http.StatusTemporaryRedirect)
		// 	return
		// }

		log.Println("Twitch Login")
		url, err := h.oauthManager.GetAuthURL()
		if err != nil {
			http.Error(w, "Failed to get auth URL", http.StatusInternalServerError)
			return
		}
		fmt.Println("Auth URL:", url)

		// if running locally, open the URL in the default web browser
		if h.cfg.Environment == "local" {
			err = exec.Command("open", url).Start() // For macOS
			if err != nil {
				log.Println("Failed to open browser:", err)

				http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			}
		}

		// Instead of trying to open the browser, return the URL to the client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"auth_url": url,
		})
	}
}

func (h *AuthHandler) TwitchCallbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		if !h.oauthManager.ValidateState(state) {
			http.Error(w, "Invalid state", http.StatusBadRequest)
			return
		}

		code := r.FormValue("code")
		token, err := h.oauthManager.Exchange(r.Context(), code)
		if err != nil {
			http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
			return
		}

		twitchUser, err := h.oauthManager.GetTwitchUserData(token.AccessToken)
		if err != nil {
			http.Error(w, "Failed to get user data", http.StatusInternalServerError)
			return
		}

		// Create or update user records
		user, err := h.userService.FindOrCreateTwitchUser(&models.TwitchUser{
			ID:    twitchUser.ID,
			Login: twitchUser.Login,
			Email: twitchUser.Email,
		}, token)
		if err != nil {
			http.Error(w, "Failed to process user", http.StatusInternalServerError)
			return
		}

		// Create session/JWT token here
		// Redirect to frontend with token
		// add user to context
		// ctx := context.WithValue(r.Context(), "user", user)

		frontendURL := h.cfg.FrontendURL
		redirectURL := fmt.Sprintf("%s?userID=%s&username=%s&email=%s", frontendURL, user.UserID, user.Username, user.Email)

		http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
	}
}
