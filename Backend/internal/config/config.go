// internal/config/config.go
package config

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	CORS     CORSConfig
	Logging  LoggingConfig
	Storage  StorageConfig
	Email    EmailConfig
	Auth     AuthConfig
	Cron     CronConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port      string
	Mode      string
	Debug     bool
	APIPrefix string
}

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	URL      string
}

// RedisConfig holds Redis connection configuration
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	URL      string
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins string
	AllowedMethods string
	AllowedHeaders string
	ExposedHeaders string
	MaxAge         int
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
}

// StorageConfig holds MinIO/S3 storage configuration
type StorageConfig struct {
	Provider             string
	Endpoint             string
	AccessKey            string
	SecretKey            string
	Bucket               string
	Prefix               string
	Region               string
	UseSSL               bool
	UsePathStyle         bool
	PublicBaseURL        string
	URLDownloadExpiryMins int
	URLPreviewExpiryMins  int
	URLUploadExpiryMins   int
}

// EmailConfig holds SMTP email configuration
type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	UseTLS   bool
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	SecretKey                 string
	JWTAccessTokenExpiryHours int
	JWTRefreshTokenExpiryDays int
}

// CronConfig holds cron job configuration
type CronConfig struct {
	Schedule string
}
