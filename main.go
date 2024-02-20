package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	keycloak := GetKeycloakClient()
	// Initialize router
	router := http.NewServeMux()

	// Use Keycloak authentication middleware for specific routes
	router.Handle("/secure-endpoint", keycloak.keycloakAuthMiddleware(http.HandlerFunc(secureEndpointHandler)))

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))

}

func secureEndpointHandler(w http.ResponseWriter, r *http.Request) {
	// Your secure endpoint logic goes here
	fmt.Fprintf(w, "Secure endpoint accessed successfully!")
}
