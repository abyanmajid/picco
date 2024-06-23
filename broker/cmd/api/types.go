package main

import (
	"context"
	"time"

	"github.com/abyanmajid/codemore.io/broker/user"
	"google.golang.org/grpc"
)

type Config struct {
	UserEndpoint string
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type GRPCClient struct {
	Client user.UserServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type UserRequest struct {
	Action string      `json:"action"`
	Data   UserPayload `json:"data"`
}

type UserPayload struct {
	ID           string    `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	Roles        []string  `json:"roles,omitempty"`
	Xp           int64     `json:"xp,omitempty"`
	IsBanned     bool      `json:"is_banned,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Token        string    `json:"token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}
