package main

import (
	"log/slog"

	"github.com/abyanmajid/codemore.io/user/internal/database"
	user "github.com/abyanmajid/codemore.io/user/proto"
)

type Service struct {
	user.UnimplementedUserServiceServer
	DB  database.Queries
	Log *slog.Logger
}

type CreateUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserByIdPayload struct {
	Id string `json:"id"`
}

type GetUserByEmailPayload struct {
	Email string `json:"email"`
}

type UpdateUserByIdPayload struct {
	Id       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
	Xp       int32    `json:"xp"`
	IsBanned bool     `json:"is_banned"`
}

type DeleteUserByIdPayload struct {
	Id string `json:"id"`
}
