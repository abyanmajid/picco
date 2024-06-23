package main

import (
	"context"

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

type CreateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetUserByIdRequest struct {
	Id string `json:"id"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

type UpdateUserByIdRequest struct {
	Id       string   `json:"id"`
	Username string   `json:"username,omitempty"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Xp       int32    `json:"xp,omitempty"`
	IsBanned bool     `json:"is_banned,omitempty"`
}

type DeleteUserByIdRequest struct {
	Id string `json:"id"`
}
