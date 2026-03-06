package bootstrap

import (
	"os"
	"strconv"
)

// Package bootstrap handles application initialization and dependency injection.
//
// This file manages application configuration loading and validation.
//
// In production-grade applications, config files typically:
// - Load from multiple sources (env vars, files, consul, vault)
// - Support different environments (dev, staging, prod)
// - Validate all required configuration
// - Provide sensible defaults
// - Handle sensitive data (secrets, API keys)
// - Support configuration hot-reloading
//
// Example structure:
//   type Config struct {
//       // Server configuration
//       ServerHost string `env:"SERVER_HOST" default:"0.0.0.0"`
//       ServerPort int    `env:"SERVER_PORT" default:"8080"`
//
//       // Database configuration
//       DBHost     string `env:"DB_HOST" validate:"required"`
//       DBPort     int    `env:"DB_PORT" default:"5432"`
//       DBUser     string `env:"DB_USER" validate:"required"`
//       DBPassword string `env:"DB_PASSWORD" validate:"required"`
//       DBName     string `env:"DB_NAME" validate:"required"`
//
//       // Redis configuration
//       RedisHost string `env:"REDIS_HOST" default:"localhost"`
//       RedisPort int    `env:"REDIS_PORT" default:"6379"`
//
//       // JWT configuration
//       JWTSecret string        `env:"JWT_SECRET" validate:"required"`
//       JWTExpiry time.Duration `env:"JWT_EXPIRY" default:"24h"`
//
//       // Feature flags
//       EnableMetrics bool `env:"ENABLE_METRICS" default:"true"`
//       EnableTracing bool `env:"ENABLE_TRACING" default:"true"`
//   }
//
// Example usage:
//   config := LoadConfig()
//   // Returns: &Config{ServerPort: 8080, DBHost: "postgres",...}

type Config struct {
	// Server
	ServerPort string

	//Database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	//Redis
	RedisHost string
	RedisPort int
	RedisPass string
	RedisDB   int

	//jwt
	JWTSecret string
	JWTExpiry int
}

func LoadConfig() *Config {

	return &Config{
		// Server
		ServerPort: getEnv("PORT", "8080"),

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "todoapp"),

		// Redis
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnvAsInt("REDIS_PORT", 6379),
		RedisPass: getEnv("REDIS_PASSWORD", ""),

		// JWT
		JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiry: getEnvAsInt("JWT_EXPIRY", 24), // hours
	}
}

func (c *Config) GetPort() string {
	return c.ServerPort
}

// Helper functions
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
