package config

import "os"

type Config struct{ Port, DatabaseURL, FinnhubAPIKey string }

func Load() Config {
	return Config{Port: value("PORT", "8080"), DatabaseURL: os.Getenv("DATABASE_URL"), FinnhubAPIKey: os.Getenv("FINNHUB_API_KEY")}
}

func value(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
