package main

import "time"

type Config struct {
	UserEndpoint string
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type UserRequest struct {
	Action string      `json:"action"`
	Data   UserPayload `json:"data"`
}

type UserPayload struct {
	ID        string    `json:"id,omitempty"`
	AuthType  string    `json:"auth_type,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Level     int       `json:"level,omitempty"`
	Badges    []string  `json:"ranks,omitempty"`
	IsBanned  bool      `json:"is_banned,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
