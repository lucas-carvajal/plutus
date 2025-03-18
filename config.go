package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_ENV = getEnvWithDefault("APP_ENV", "local")

	FINNHUB_API_KEY            = getEnvOrPanic("FINNHUB_API_KEY")
	TWELVE_DATA_API_KEY        = getEnvOrPanic("TWELVE_DATA_API_KEY")
	POSTGRES_CONNECTION_STRING = getEnvOrPanic("POSTGRES_CONNECTION_STRING")
)

func getEnvWithDefault(key, defaultValue string) string {
	value, exists := getEnvVar(key)
	if !exists || value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrPanic(key string) string {
	value, exists := getEnvVar(key)
	if !exists || value == "" {
		panic("Missing env variable " + key)
	}
	return value
}

func getEnvVar(key string) (string, bool) {
	if value := os.Getenv(key); value != "" {
		return value, true
	}
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not loaded: %v", err)
	}
	if value := os.Getenv(key); value != "" {
		return value, true
	}
	return "", false
}
