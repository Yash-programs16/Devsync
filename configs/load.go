package configs

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Load reads all configuration from environment variables.
func Load() (*Config, error) {
	// Load .env for local development.
	// Ignore the error because production environments
	// usually inject environment variables directly.
	_ = godotenv.Load()

	cfg := &Config{
		// ==========================
		// Application
		// ==========================
		AppEnv:            GetEnv("APP_ENV", "development"),
		HTTPPort:          GetEnv("HTTP_PORT", "8080"),
		WorkspaceHTTPPort: GetEnv("WORKSPACE_HTTP_PORT", "8081"),
		GRPCPort:          GetEnv("GRPC_PORT", "50051"),
		LogLevel:          GetEnv("LOG_LEVEL", "info"),

		// ==========================
		// Database
		// ==========================
		DatabaseURL: GetEnv("DATABASE_URL", ""),
		DBMaxConns:  int32(GetEnvAsInt("DB_MAX_CONNS", 25)),
		DBMinConns:  int32(GetEnvAsInt("DB_MIN_CONNS", 5)),

		// ==========================
		// Authentication
		// ==========================
		PasetoKey:            GetEnv("PASETO_KEY", ""),
		AccessTokenDuration:  GetEnvAsDuration("ACCESS_TOKEN_DURATION", 15*time.Minute),
		RefreshTokenDuration: GetEnvAsDuration("REFRESH_TOKEN_DURATION", 7*24*time.Hour),

		// ==========================
		// Redis
		// ==========================
		RedisHost:     GetEnv("REDIS_HOST", "localhost"),
		RedisPort:     GetEnv("REDIS_PORT", "6379"),
		RedisPassword: GetEnv("REDIS_PASSWORD", ""),
		RedisDB:       GetEnvAsInt("REDIS_DB", 0),

		// ==========================
		// SMTP
		// ==========================
		SMTPHost:     GetEnv("SMTP_HOST", ""),
		SMTPPort:     GetEnv("SMTP_PORT", "587"),
		SMTPUsername: GetEnv("SMTP_USERNAME", ""),
		SMTPPassword: GetEnv("SMTP_PASSWORD", ""),
		SMTPFrom:     GetEnv("SMTP_FROM", ""),

		// ==========================
		// Frontend
		// ==========================
		FrontendURL: GetEnv("FRONTEND_URL", "http://localhost:3000"),

		// ==========================
		// Uploads
		// ==========================
		UploadDir:     GetEnv("UPLOAD_DIR", "./uploads"),
		MaxUploadSize: int64(GetEnvAsInt("MAX_UPLOAD_SIZE", 10485760)), // 10 MB
	}

	// ==========================
	// Validation
	// ==========================

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	if cfg.PasetoKey == "" {
		return nil, fmt.Errorf("PASETO_KEY is required")
	}

	fmt.Println("SMTP HOST:", cfg.SMTPHost)
	fmt.Println("SMTP PORT:", cfg.SMTPPort)
	fmt.Println("SMTP USER:", cfg.SMTPUsername)
	fmt.Println("SMTP PASSWORD LENGTH:", len(cfg.SMTPPassword))
	return cfg, nil
}

// GetEnv returns an environment variable or a fallback value.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// GetEnvAsInt returns an integer environment variable.
func GetEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return i
}

// GetEnvAsDuration returns a time.Duration environment variable.
// Example:
// ACCESS_TOKEN_DURATION=15m
// REFRESH_TOKEN_DURATION=168h
func GetEnvAsDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}

	return duration
}
