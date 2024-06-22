package main

import "net/http"

func (api *Config) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// Validate user credentials

	// Generate JWT and a refresh token

	// Send token to the user

}

func (api *Config) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Validate the refresh token

	// Generate a new JWT (and possibly a new refresh token)

	// Send token to the user

}

func (api *Config) HandleLogout(w http.ResponseWriter, r *http.Request) {
	// Invalidate the refresh token

	// Send a response confirming logout
}

// func JWTAuthMiddleware(next http.Handler) http.Handler {
// 	// Verify the JWT in the Authorization header

// 	// Allow access if the token is valid; otherwise deny access
// }
