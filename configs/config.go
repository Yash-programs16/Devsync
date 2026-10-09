// Package configs stores all application configuration loaded from
// environment variables (.env).
package configs

import "time"

// Config holds the application's runtime configuration.
type Config struct {
	// Application
	AppEnv            string
	HTTPPort          string
	WorkspaceHTTPPort string
	GRPCPort          string
	LogLevel          string

	// Database
	DatabaseURL string
	DBMaxConns  int32
	DBMinConns  int32

	// Authentication
	PasetoKey            string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration

	// Redis
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	// SMTP
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	SMTPFrom     string

	// Frontend
	FrontendURL string

	// File Uploads
	UploadDir     string
	MaxUploadSize int64
}
