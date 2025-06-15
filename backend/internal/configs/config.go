package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port        int
	Version     string
	JWTSecret   string
	JWTIssuer   string
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

	if os.Getenv("API_VERSION") == "" {
		cfg.Version = "v1"
	} else {
		cfg.Version = os.Getenv("API_VERSION")
	}
	if os.Getenv("JWT_SECRET") == "" {
		return nil, fmt.Errorf("JWT_SECRET is not set")
	} else {
		cfg.JWTSecret = os.Getenv("JWT_SECRET")
	}

	if os.Getenv("JWT_ISSUER") == "" {
		cfg.JWTIssuer = "sanqa-suq"
	} else {
		cfg.JWTIssuer = os.Getenv("JWT_ISSUER")
	}

	return cfg, nil

}
