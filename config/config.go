// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type AppConfig struct {
// 	KachaUsername string
// 	KachaPassword string
// 	KachaBaseURL  string
// 	Port          string
// }

// func Load() (*AppConfig, error) {
// 	if err := godotenv.Load(); err != nil {
// 		log.Printf("No .env file found: %v", err)
// 	}

// 	cfg := &AppConfig{
// 		KachaUsername: os.Getenv("KACHA_USERNAME"),
// 		KachaPassword: os.Getenv("KACHA_PASSWORD"),
// 		KachaBaseURL:  os.Getenv("KACHA_BASE_URL"),
// 		Port:          os.Getenv("PORT"),
// 	}

// 	if cfg.Port == "" {
// 		cfg.Port = "8080"
// 	}

// 	if cfg.KachaUsername == "" || cfg.KachaPassword == "" {
// 		return nil, fmt.Errorf("missing required env vars: KACHA_USERNAME and KACHA_PASSWORD")
// 	}

// 	return cfg, nil
// }

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	KachaBaseURL string
	Port         string
}

func Load() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	cfg := &AppConfig{
		KachaBaseURL: os.Getenv("KACHA_BASE_URL"),
		Port:         os.Getenv("PORT"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg, nil
}
