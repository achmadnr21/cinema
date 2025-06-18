package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
}
type Token struct {
	AccessSecret  string
	RefreshSecret string
}
type Config struct {
	Database Database `json:"database"`
	TokenCFG Token    `json:"token"`
	Port     int      `json:"port"`
}

// LoadConfig loads the configuration from a file or environment variables.
func LoadConfig() (*Config, error) {
	godotenv.Load()
	return &Config{
		Database: Database{
			Host:     os.Getenv("PG_HOST"),
			Port:     func() int { p, _ := strconv.Atoi(os.Getenv("PG_PORT")); return p }(),
			Username: os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Database: os.Getenv("PG_DATABASE"),
			SSLMode:  os.Getenv("PG_SSLMODE"),
		},
		TokenCFG: Token{
			AccessSecret:  os.Getenv("ACCESS_SECRET"),
			RefreshSecret: os.Getenv("REFRESH_SECRET"),
		},
		Port: func() int { p, _ := strconv.Atoi(os.Getenv("PORT")); return p }(),
	}, nil
}
