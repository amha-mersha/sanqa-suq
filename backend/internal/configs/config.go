package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port        int
}

func LoadConfig(envFile string) (*Config, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	cfg := &Config{}
	if os.Getenv("PORT") == "" {
		cfg.Port = 8080
	} else {
		port, errPars := strconv.Atoi(os.Getenv("PORT"))
		if errPars != nil {
			return nil, fmt.Errorf("failed to parse PORT: %w", errPars)
		}
		cfg.Port = port
	}
	if os.Getenv("DATABASE_URL") == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	} else {
		cfg.DatabaseUrl = os.Getenv("DATABASE_URL")
	}
	return cfg, nil
}
