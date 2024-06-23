package main

import (
	"context"
	"time"

	"github.com/abyanmajid/codemore.io/user/internal/database"
	user "github.com/abyanmajid/codemore.io/user/proto"
	utils "github.com/abyanmajid/codemore.io/user/utils"
	"github.com/google/uuid"
)

func (api *Service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	api.Log.Info("Starting CreateUser", "username", req.GetUsername(), "email", req.GetEmail())

	payload := CreateUserPayload{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		api.Log.Error("Error hashing password", "error", err)
		return nil, err
	}

	u, err := api.DB.CreateUser(ctx, database.CreateUserParams{
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
		api.Log.Error("Error creating user in database", "error", err)
		return nil, err
	}

	api.Log.Info("User created successfully", "user_id", u.ID)
	return utils.EncodeProtoUser(u), nil
}

func (api *Service) GetAllUsers(ctx context.Context, req *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {
	api.Log.Info("Starting GetAllUsers")

	users, err := api.DB.GetAllUsers(ctx)
	if err != nil {
		api.Log.Error("Error fetching all users from database", "error", err)
		return nil, err
	}

	res := &user.GetAllUsersResponse{}
	for _, u := range users {
		res.Users = append(res.Users, utils.EncodeProtoUser(u))
	}

	api.Log.Info("Successfully fetched all users", "user_count", len(users))
	return res, nil
}

func (api *Service) GetUserById(ctx context.Context, req *user.GetUserByIdRequest) (*user.User, error) {
	api.Log.Info("Starting GetUserById", "user_id", req.GetId())

	payload := GetUserByIdPayload{
		Id: req.GetId(),
	}

	parsedId, err := uuid.Parse(payload.Id)
	if err != nil {
		api.Log.Error("Error parsing user ID", "error", err)
		return nil, err
	}

	u, err := api.DB.GetUserById(ctx, parsedId)
	if err != nil {
		api.Log.Error("Error fetching user by ID from database", "error", err)
		return nil, err
	}

	api.Log.Info("Successfully fetched user by ID", "user_id", u.ID)
	return utils.EncodeProtoUser(u), nil
}

func (api *Service) GetUserByEmail(ctx context.Context, req *user.GetUserByEmailRequest) (*user.User, error) {
	api.Log.Info("Starting GetUserByEmail", "email", req.GetEmail())

	payload := GetUserByEmailPayload{
		Email: req.GetEmail(),
	}

	u, err := api.DB.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		api.Log.Error("Error fetching user by email from database", "error", err)
		return nil, err
	}

	api.Log.Info("Successfully fetched user by email", "user_id", u.ID)
	return utils.EncodeProtoUser(u), nil
}

func (api *Service) UpdateUserById(ctx context.Context, req *user.UpdateUserByIdRequest) (*user.User, error) {
	api.Log.Info("Starting UpdateUserById", "user_id", req.GetId())

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
		api.Log.Error("Error parsing user ID", "error", err)
		return nil, err
	}

	u, err := api.DB.UpdateUserById(ctx, database.UpdateUserByIdParams{
		ID:       parsedId,
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		Roles:    payload.Roles,
		Xp:       payload.Xp,
		IsBanned: payload.IsBanned,
	})

	if err != nil {
		api.Log.Error("Error updating user in database", "error", err)
		return nil, err
	}

	api.Log.Info("User updated successfully", "user_id", u.ID)
	return utils.EncodeProtoUser(u), nil
}

func (api *Service) DeleteUserById(ctx context.Context, req *user.DeleteUserByIdRequest) (*user.DeleteUserByIdResponse, error) {
	api.Log.Info("Starting DeleteUserById", "user_id", req.GetId())

	payload := DeleteUserByIdPayload{
		Id: req.GetId(),
	}

	parsedId, err := uuid.Parse(payload.Id)
	if err != nil {
		api.Log.Error("Error parsing user ID", "error", err)
		return nil, err
	}

	err = api.DB.DeleteUserById(ctx, parsedId)
	if err != nil {
		api.Log.Error("Error deleting user from database", "error", err)
		return nil, err
	}

	api.Log.Info("User deleted successfully", "user_id", payload.Id)
	return &user.DeleteUserByIdResponse{
		Message: "User deleted successfully",
	}, nil
}
