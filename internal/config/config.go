package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type Config struct {
	ServerConfig  ServerConfig
	LoggingConfig LoggingConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type LoggingConfig struct {
	LogLevel string
}

func ErrInvalidEnvVar(varName, value, reason string) error {
	return errors.New("invalid environment variable value: " + varName +
		"='" + value + "' (" + reason + ")")
}

// LoadEnv loads environment variables from .env file
// It's optional - if .env doesn't exist, it uses system env vars
func LoadEnv() {
	// Load .env file if it exists (ignore error if not found)
	_ = godotenv.Load()
}

// New creates a new Config from environment variables
func New() *Config {
	// Load .env file first
	LoadEnv()

	cfg := Config{
		ServerConfig: ServerConfig{
			Host: GetEnv("SHADIS_HOST", "0.0.0.0"),
			Port: GetEnv("SHADIS_PORT", 6379),
		},
		LoggingConfig: LoggingConfig{
			LogLevel: GetEnv("SHADIS_LOG_LEVEL", string(LogLevelInfo)),
		},
	}
	return &cfg
}

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError:
		return true
	}
	return false
}

func (cfg *Config) Validate() error {
	if cfg.ServerConfig.Port < 1 || cfg.ServerConfig.Port > 65535 {
		return ErrInvalidEnvVar("SHADIS_PORT", strconv.Itoa(cfg.ServerConfig.Port), "invalid port")
	}
	if !LogLevel(cfg.LoggingConfig.LogLevel).IsValid() {
		return ErrInvalidEnvVar("SHADIS_LOG_LEVEL", cfg.LoggingConfig.LogLevel, "invalid log level")
	}
	return nil
}

// GetEnv retrieves environment variable with type inference and default value
func GetEnv[T any](key string, defaultValue T) T {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	var anyValue any

	switch any(defaultValue).(type) {
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return defaultValue
		}
		anyValue = i
	case string:
		anyValue = val
	case bool:
		b, err := strconv.ParseBool(val)
		if err != nil {
			return defaultValue
		}
		anyValue = b
	}
	return anyValue.(T)
}
