package main

import "github.com/abyanmajid/codemore.io/user/internal/database"

type Config struct {
	DB *database.Queries
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type CreateUserPayload struct {
	AuthType string `json:"auth_type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
