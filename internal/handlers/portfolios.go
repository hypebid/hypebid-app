package handlers

import (
	"net/http"

	"github.com/hypebid/hypebid-app/internal/config"
)

func GetUserPortfolioHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement user portfolio getting logic
	}
}
