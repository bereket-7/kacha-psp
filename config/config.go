package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	KachaAppID   string
	KachaAPIKey  string
	KachaBaseURL string
	Port         string
}

func Load() (*AppConfig, error) {
	// Load .env if present
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	cfg := &AppConfig{
		KachaAppID:   os.Getenv("KACHA_APP_ID"),
		KachaAPIKey:  os.Getenv("KACHA_API_KEY"),
		KachaBaseURL: os.Getenv("KACHA_BASE_URL"),
		Port:         os.Getenv("PORT"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	if cfg.KachaAppID == "" || cfg.KachaAPIKey == "" {
		return nil, fmt.Errorf("missing required env: KACHA_APP_ID and KACHA_API_KEY")
	}

	return cfg, nil
}


