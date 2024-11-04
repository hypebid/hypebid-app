package handlers

import (
	"net/url"
	"strings"
)

type BaseHandlerConfig struct {
	// Config *config.Config
	// Logger *slog.Logger
	// Auth *auth.Auth
}

func NormalizeURLParam(param string) (string, error) {
	param, err := url.QueryUnescape(param)
	if err != nil {
		return "", err
	}
	return strings.ToLower(param), nil
}
