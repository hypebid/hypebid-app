package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientID               string
	ClientSecret           string
	DBHost                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBPort                 string
	FrontendURL            string
	Environment            string
	ServerPort             string
	FollowerUpdateInterval string
	TrackedLogins          []string
	HostURL                string
}

func Load() (*Config, error) {
	// Load .env file only if not in production
	fmt.Println("ENVIRONMENT:", os.Getenv("ENVIRONMENT"))
	fmt.Println("os.Getenv:", os.Getenv("ENVIRONMENT"))
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	if _, err := strconv.Atoi(os.Getenv("FOLLOWER_UPDATE_INTERVAL")); err != nil {
		// Default to 15 minutes
		os.Setenv("FOLLOWER_UPDATE_INTERVAL", "15")
	}

	var trackedLogins []string
	if loginStr := os.Getenv("TRACKED_LOGINS"); loginStr != "" {
		trackedLogins = strings.Split(loginStr, ",")
		// Trim spaces from each login
		for i, login := range trackedLogins {
			trackedLogins[i] = strings.TrimSpace(login)
		}
	} else {
		// Default logins if none provided
		trackedLogins = []string{
			"Rubius",
			"KaiCenat",
			"caseoh_",
			"ElMariana",
			"Jynxzi",
			"PirateSoftware",
			"Mictia00",
			"FeirlyGab",
			"brino",
			"GUACAMOLEMOLLY",
			"Lirik",
		}
	}

	return &Config{
		ClientID:               os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret:           os.Getenv("TWITCH_CLIENT_SECRET"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBName:                 os.Getenv("DB_NAME"),
		DBPort:                 os.Getenv("DB_PORT"),
		FrontendURL:            os.Getenv("FRONTEND_URL"),
		Environment:            os.Getenv("ENVIRONMENT"),
		ServerPort:             os.Getenv("SERVER_PORT"),
		FollowerUpdateInterval: os.Getenv("FOLLOWER_UPDATE_INTERVAL"),
		TrackedLogins:          trackedLogins,
		HostURL:                os.Getenv("HOST_URL"),
	}, nil
}
