// internal/config/loader.go
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load environment file based on APP_ENV
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "test" {
		if err := godotenv.Load(".env.test"); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not load .env.test file: %v\n", err)
		}
	} else if os.Getenv("DOCKER_ENV") != "true" {
		if err := loadEnvFile(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not load local .env file: %v\n", err)
		}
	}

	cfg := &Config{}

	// Server
	cfg.Server.Port = getEnv("APP_PORT", "8080")
	cfg.Server.Mode = getEnv("APP_ENV", "development")
	cfg.Server.Debug = getEnvAsBool("APP_DEBUG", true)
	cfg.Server.APIPrefix = getEnv("API_PREFIX", "/api/v1")

	// Database (MySQL)
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "3306")
	cfg.Database.User = getEnv("DB_USER", "root")
	cfg.Database.Password = getEnv("DB_PASSWORD", "")
	cfg.Database.Name = getEnv("DB_NAME", "be_db")
	cfg.Database.SSLMode = getEnv("DB_SSL_MODE", "disable")
	cfg.Database.URL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	// Redis
	cfg.Redis.Host = getEnv("REDIS_HOST", "localhost")
	cfg.Redis.Port = getEnv("REDIS_PORT", "6379")
	cfg.Redis.Password = getEnv("REDIS_PASSWORD", "")
	cfg.Redis.DB = getEnvAsInt("REDIS_DB", 0)
	if cfg.Redis.Password != "" {
		cfg.Redis.URL = fmt.Sprintf("redis://:%s@%s:%s/%d", cfg.Redis.Password, cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.DB)
	} else {
		cfg.Redis.URL = fmt.Sprintf("redis://%s:%s/%d", cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.DB)
	}

	// CORS
	cfg.CORS.AllowedOrigins = getEnv("CORS_ALLOWED_ORIGINS", "*")
	cfg.CORS.AllowedMethods = getEnv("CORS_ALLOWED_METHODS", "GET,POST,PUT,DELETE,OPTIONS,PATCH")
	cfg.CORS.AllowedHeaders = getEnv("CORS_ALLOWED_HEADERS", "Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,accept,origin,Cache-Control,X-Requested-With")
	cfg.CORS.ExposedHeaders = getEnv("CORS_EXPOSED_HEADERS", "Content-Length,Content-Type,Authorization")
	cfg.CORS.MaxAge = getEnvAsInt("CORS_MAX_AGE", 86400)

	// Logging
	cfg.Logging.Level = getEnv("LOG_LEVEL", "debug")
	cfg.Logging.Format = getEnv("LOG_FORMAT", "json")

	// MinIO Storage
	cfg.Storage.Provider = getEnv("STORAGE_PROVIDER", "minio")
	cfg.Storage.Endpoint = getEnv("MINIO_ENDPOINT", "localhost:9000")
	cfg.Storage.AccessKey = getEnv("MINIO_ACCESS_KEY", "minioadmin")
	cfg.Storage.SecretKey = getEnv("MINIO_SECRET_KEY", "minioadmin")
	cfg.Storage.Bucket = getEnv("MINIO_BUCKET", "documents")
	cfg.Storage.Prefix = getEnv("MINIO_PREFIX", "uploads")
	cfg.Storage.Region = getEnv("MINIO_REGION", "us-east-1")
	cfg.Storage.UseSSL = getEnvAsBool("MINIO_USE_SSL", false)
	cfg.Storage.UsePathStyle = getEnvAsBool("MINIO_USE_PATH_STYLE", true)
	cfg.Storage.PublicBaseURL = getEnv("MINIO_PUBLIC_BASE_URL", fmt.Sprintf("http://%s", cfg.Storage.Endpoint))
	cfg.Storage.URLDownloadExpiryMins = getEnvAsInt("STORAGE_URL_DOWNLOAD_EXPIRY_MINS", 60)
	cfg.Storage.URLPreviewExpiryMins = getEnvAsInt("STORAGE_URL_PREVIEW_EXPIRY_MINS", 15)
	cfg.Storage.URLUploadExpiryMins = getEnvAsInt("STORAGE_URL_UPLOAD_EXPIRY_MINS", 30)

	// Email (SMTP)
	cfg.Email.Host = getEnv("SMTP_HOST", "")
	cfg.Email.Port = getEnvAsInt("SMTP_PORT", 587)
	cfg.Email.Username = getEnv("SMTP_USERNAME", "")
	cfg.Email.Password = getEnv("SMTP_PASSWORD", "")
	cfg.Email.From = getEnv("SMTP_FROM", "no-reply@example.com")
	cfg.Email.UseTLS = getEnvAsBool("SMTP_USE_TLS", true)

	// Auth
	cfg.Auth.SecretKey = getEnv("AUTH_SECRET_KEY", "default-secret-key-change-in-production")
	cfg.Auth.JWTAccessTokenExpiryHours = getEnvAsInt("JWT_ACCESS_TOKEN_EXPIRY_HOURS", 24)
	cfg.Auth.JWTRefreshTokenExpiryDays = getEnvAsInt("JWT_REFRESH_TOKEN_EXPIRY_DAYS", 7)

	// Cron
	cfg.Cron.Schedule = getEnv("CRON_SCHEDULE", "@daily")

	return cfg, nil
}

func loadEnvFile() error {
	envFile := ".env"
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		return fmt.Errorf("env file not found: %s", envFile)
	}
	return godotenv.Load(envFile)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}
