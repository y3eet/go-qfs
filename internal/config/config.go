package config

import (
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	AppEnv      string
	FrontendURL string
	BaseDir     string
	IP          string
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
	ip, err := getLocalIP()
	if err != nil {
		log.Println("Error getting IP")
	}

	Cfg = &Config{
		Port:        getEnv("PORT", "8080"),
		AppEnv:      getEnv("APP_ENV", "dev"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:5173"),
		BaseDir:     getEnv("BASE_DIR", "./"), // default to exe dir, not "./"
		IP:          ip,
	}

	return Cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
