package main

import (
	"context"
	"fmt"
	"time"

	"github.com/abyanmajid/codemore.io/user/internal/database"
	user "github.com/abyanmajid/codemore.io/user/proto"
	"github.com/abyanmajid/codemore.io/user/utils"
	"github.com/google/uuid"
)

func (s *UserServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {

	fmt.Println("Username REq: ", req.GetUsername())

	requestPayload := CreateUserPayload{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	fmt.Println("Username: ", requestPayload.Username)
	fmt.Println("Email: ", requestPayload.Email)
	fmt.Println("Password: ", requestPayload.Password)

	hashedPassword, err := utils.HashPassword(requestPayload.Password)
	if err != nil {
		return nil, err
	}

	u, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		Username:  requestPayload.Username,
		Email:     requestPayload.Email,
		Password:  hashedPassword,
		Roles:     []string{"user"},
		Xp:        0,
		IsBanned:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	res := &user.CreateUserResponse{
		Id:        u.ID.String(),
		Username:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		Roles:     u.Roles,
		Xp:        u.Xp,
		IsBanned:  u.IsBanned,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}

	return res, nil
}
