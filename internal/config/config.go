package config

import (
	"errors"
	"os"
)

type Config struct {
	DatabaseURL string
	RedisAddr   string
	JWTSecret   string
}

func LoadConfig() (*Config, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL is required")
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		return nil, errors.New("REDIS_ADDR is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET is required")
	}

	return &Config{
		DatabaseURL: databaseURL,
		RedisAddr:   redisAddr,
		JWTSecret:   jwtSecret,
	}, nil

}
