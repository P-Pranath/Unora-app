// pkg/logger/logger.go
package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var (
	globalLogger zerolog.Logger
	once         sync.Once
)

// InitLogger initializes the global logger
func InitLogger(env string) {
	once.Do(func() {
		var output zerolog.ConsoleWriter

		if env == "production" {
			// JSON format for production
			globalLogger = zerolog.New(os.Stdout).With().Timestamp().Logger()
		} else {
			// Pretty console output for development
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			}
			globalLogger = zerolog.New(output).With().Timestamp().Logger()
		}

		// Set global log level
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	})
}

// GetLogger returns a logger with a specific component name
func GetLogger(component string) zerolog.Logger {
	return globalLogger.With().Str("component", component).Logger()
}

// Debug returns a debug level event
func Debug() *zerolog.Event {
	return globalLogger.Debug()
}

// Info returns an info level event
func Info() *zerolog.Event {
	return globalLogger.Info()
}

// Warn returns a warning level event
func Warn() *zerolog.Event {
	return globalLogger.Warn()
}

// Error returns an error level event
func Error() *zerolog.Event {
	return globalLogger.Error()
}

// Fatal returns a fatal level event
func Fatal() *zerolog.Event {
	return globalLogger.Fatal()
}

// WithUserID returns a logger with user context
func WithUserID(userID string) zerolog.Logger {
	return globalLogger.With().Str("user_id", userID).Logger()
}

// WithRequestID returns a logger with request context
func WithRequestID(requestID string) zerolog.Logger {
	return globalLogger.With().Str("request_id", requestID).Logger()
}

// WithContext returns a logger with custom context
func WithContext(fields map[string]interface{}) zerolog.Logger {
	ctx := globalLogger.With()
	for k, v := range fields {
		ctx = ctx.Interface(k, v)
	}
	return ctx.Logger()
}
