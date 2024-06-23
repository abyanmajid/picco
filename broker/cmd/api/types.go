package main

import (
	"context"
	"log/slog"

	"github.com/abyanmajid/codemore.io/broker/user"
	"google.golang.org/grpc"
)

type Service struct {
	UserEndpoint string
	Log          *slog.Logger
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

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserByIdRequest struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
	Xp       int32    `json:"xp"`
	IsBanned bool     `json:"is_banned"`
}
