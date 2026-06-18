package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	AppEnv      string
	FrontendURL string
	BaseDir     string
}

var Cfg *Config

func execDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	// Resolve symlinks (important on Linux)
	resolved, err := filepath.EvalSymlinks(exe)
	if err != nil {
		return filepath.Dir(exe)
	}
	return filepath.Dir(resolved)
}

func Load() *Config {
	if Cfg != nil {
		return Cfg
	}

	dir := execDir()

	// Load .env relative to the binary, not CWD
	if err := godotenv.Load(filepath.Join(dir, ".env")); err != nil {
		log.Println("No .env file found")
	}

	Cfg = &Config{
		Port:        getEnv("PORT", "8080"),
		AppEnv:      getEnv("APP_ENV", "dev"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),
		BaseDir:     getEnv("BASE_DIR", "./"), // default to exe dir, not "./"
	}

	return Cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
