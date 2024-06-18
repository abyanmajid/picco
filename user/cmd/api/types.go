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
