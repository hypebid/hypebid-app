package handlers

import (
	"net/http"

	"github.com/hypebid/hypebid-app/internal/config"
)

func GetBoostsForUserHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement boosts getting logic
	}
}

func GetActiveBoostsForUserHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement active boosts getting logic
	}
}
