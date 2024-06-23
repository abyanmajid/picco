package main

import (
	"context"
	"time"

	"github.com/abyanmajid/codemore.io/user/internal/database"
	user "github.com/abyanmajid/codemore.io/user/proto"
	utils "github.com/abyanmajid/codemore.io/user/utils"
	"github.com/google/uuid"
)

func (s *UserServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	payload := CreateUserPayload{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	u, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		Username:  payload.Username,
		Email:     payload.Email,
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

	return utils.EncodeProtoUser(u), nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {
	users, err := s.DB.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	res := &user.GetAllUsersResponse{}

	for _, u := range users {
		res.Users = append(res.Users, utils.EncodeProtoUser(u))
	}

	return res, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *user.GetUserByIdRequest) (*user.User, error) {
	payload := GetUserByIdPayload{
		Id: req.GetId(),
	}

	parsedId, err := uuid.Parse(payload.Id)
	if err != nil {
		return nil, err
	}

	u, err := s.DB.GetUserById(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	return utils.EncodeProtoUser(u), nil
}

func (s *UserServer) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.User, error) {
	payload := GetUserByEmailPayload{
		Email: req.GetEmail(),
	}

	u, err := s.DB.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		return nil, err
	}

	return utils.EncodeProtoUser(u), nil
}

func (s *UserServer) UpdateUserById(ctx context.Context, req *user.UpdateUserByIdRequest) (*user.User, error) {
	payload := UpdateUserByIdPayload{
		Id:       req.GetId(),
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Roles:    req.GetRoles(),
		Xp:       req.GetXp(),
		IsBanned: req.GetIsBanned(),
	}

	parsedId, err := uuid.Parse(payload.Id)
	if err != nil {
		return nil, err
	}

	u, err := s.DB.UpdateUserById(ctx, database.UpdateUserByIdParams{
		ID:       parsedId,
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Roles:    payload.Roles,
		Xp:       payload.Xp,
		IsBanned: payload.IsBanned,
	})

	if err != nil {
		return nil, err
	}

	return utils.EncodeProtoUser(u), nil
}

func (s *UserServer) DeleteUserById(ctx context.Context, req *user.DeleteUserByIdRequest) (*user.DeleteUserByIdResponse, error) {
	payload := DeleteUserByIdPayload{
		Id: req.GetId(),
	}

	parsedId, err := uuid.Parse(payload.Id)
	if err != nil {
		return nil, err
	}

	err = s.DB.DeleteUserById(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	return &user.DeleteUserByIdResponse{
		Message: "User deleted successfully",
	}, nil
}
