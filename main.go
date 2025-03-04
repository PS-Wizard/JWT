package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PS-Wizard/JWT/internals"
)

func main() {
	// Serve a basic route and protected route
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", protectedHandler)

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// loginHandler is where the JWT is generated after basic auth
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// For demo purposes, we assume basic auth is passed
	username, password, ok := r.BasicAuth()
	if !ok || username != "user" || password != "password" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Set JWT expiration (e.g., 1 hour after login)
	jwtConfig := &internals.JWTConfig{
		Alg:    "HS256",
		Typ:    "JWT",
		Secret: "super secret 69",
	}

	jwtConfig.Claims.Expiration = uint64(time.Now().Add(time.Hour).Unix())
	jwtConfig.Claims.Sub = "123"
	jwt := generateJWT(jwtConfig)

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	})
	w.Write([]byte("Login successful! JWT set in cookie."))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Get JWT from cookies
	cookie, err := r.Cookie("jwt")
	if err != nil {
		fmt.Println("Here?")
		fmt.Println(cookie, err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Validate the JWT (checking expiration and signature)
	valid, err := validateJWT(cookie.Value)
	fmt.Println(cookie.Value)
	if err != nil || !valid {
		fmt.Println(valid, err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// If JWT is valid, allow access
	w.Write([]byte("Welcome to the protected route!"))
}

// generateJWT generates the actual JWT string
func generateJWT(j *internals.JWTConfig) string {
	header := j.GenerateHeader()
	payload := j.GeneratePayload()
	signature := j.GenerateSignature()

	// Combine all parts to form the JWT
	return fmt.Sprintf("%s.%s.%s", header, payload, signature)
}

// validateJWT validates the JWT by checking the signature and expiration
func validateJWT(jwt string) (bool, error) {
	// Split JWT into parts (header, payload, signature)
	parts := splitJWT(jwt)
	if len(parts) != 3 {
		return false, fmt.Errorf("invalid JWT format")
	}

	payload, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("error decoding payload: %v", err)
	}

	// Unmarshal the payload into the claims
	var claims internals.Claims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return false, fmt.Errorf("error unmarshaling payload: %v", err)
	}

	// Check expiration
	if time.Unix(int64(claims.Expiration), 0).Before(time.Now()) {
		return false, fmt.Errorf("token has expired")
	}

	return true, nil
}

// Helper function to split JWT into parts
func splitJWT(jwt string) []string {
	return strings.Split(jwt, ".")
}
