package config

import "os"

type Config struct {
	AppPort     string
	PostgresDSN string
}

func Load() Config {
	return Config{
		AppPort:     getEnv("APP_PORT", "8080"),
		PostgresDSN: getEnv("POSTGRES_DSN", "postgres://crm:crm@localhost:5432/crm?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
