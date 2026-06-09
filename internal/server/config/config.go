package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	JWTSecret    string
	AppEnv       string
	FrontendURL  string
	BuildVersion string
	RootKey      string
}

var Cfg *Config

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	if Cfg != nil {
		return Cfg
	}

	Cfg = &Config{
		Port:         getEnv("PORT", "8080"),
		AppEnv:       getEnv("APP_ENV", "dev"),
		FrontendURL:  getEnv("FRONTEND_URL", "http://localhost:5173"),
	}
	return Cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
