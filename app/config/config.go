package config

import "os"

func getEnvWithFallback(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

var (
	DATABASE_URL = getEnvWithFallback("DATABASE_URL", "mongodb://localhost:27017/test")
	PORT = getEnvWithFallback("PORT", "8081")
	DB_USER = getEnvWithFallback("DB_USER", "admin")
	DB_PASSWORD = getEnvWithFallback("DB_PASSWORD", "password")
	DB_NAME = getEnvWithFallback("DB_NAME", "admin")
	URL_RECORD_COLLECTION = getEnvWithFallback("URL_RECORD_COLLECTION", "url_records")
)