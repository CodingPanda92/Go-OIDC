package main

import (
	"os"
)

// Config holds the Keycloak configuration
type Config struct {
	Realm          string
	URL            string
	ClientID       string
	ClientSecret   string
	RedirectionURL string
	AllowedClaims  []string
}

// NewConfig creates a new Config instance with values from environment variables
func NewConfig() *Config {
	return &Config{
		Realm:          getEnv("KEYCLOAK_REALM", "demo"),
		URL:            getEnv("KEYCLOAK_URL", "http://localhost:8080/auth"),
		ClientID:       getEnv("KEYCLOAK_CLIENT_ID", "demo"),
		ClientSecret:   getEnv("KEYCLOAK_CLIENT_SECRET", ""),
		RedirectionURL: getEnv("KEYCLOAK_REDIRECTION_URL", "http://localhost:8888/test"),
		AllowedClaims:  getEnvArr("KEYCLOAK_ALLOWED_CLAIMS", []string{"demo"}),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvArr(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return []string{value}
	}
	return defaultValue
}
