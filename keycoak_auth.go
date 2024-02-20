package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
)

type keycloak struct {
	config *Config
	client *gocloak.GoCloak
}

func GetKeycloakClient() *keycloak {
	config := NewConfig()
	client := gocloak.NewClient(config.URL)
	return &keycloak{
		config,
		client,
	}
}

func (keycloak *keycloak) keycloakAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := keycloak.extractTokenFromRequest(r)
		token, _, err := keycloak.client.DecodeAccessToken(context.Background(), tokenString, keycloak.config.Realm)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !containsString(keycloak.config.AllowedClaims, token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func containsString(slice []string, target *jwt.Token) bool {
	jsonData := target.Raw
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	groups, found := data["groups"].([]string)
	if found {
		for _, s := range slice {
			for _, g := range groups {
				if s == g {
					return true
				}
			}
		}
	} else {
		fmt.Println("Error: No Groups Claim in jwt Token.")
		return false
	}
	return false
}

func (keycloak *keycloak) extractTokenFromRequest(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return ""
	}
	authHeaderParts := strings.Split(authorizationHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return ""
	}
	return authHeaderParts[1]
}
