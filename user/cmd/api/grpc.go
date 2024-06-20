package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/abyanmajid/codemore.io/user/internal/database"
	user "github.com/abyanmajid/codemore.io/user/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
	DB database.Queries
}

func (s *UserServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	payload := CreateUserPayload{
		AuthType: req.GetAuthType(),
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	u, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		AuthType:  payload.AuthType,
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  sql.NullString{String: payload.Password, Valid: true},
		Level:     1,
		Badges:    []string{},
		IsBanned:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	res := &user.CreateUserResponse{
		Id:        u.ID.String(),
		AuthType:  u.AuthType,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password.String,
		Level:     u.Level,
		Badges:    u.Badges,
		IsBanned:  u.IsBanned,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}

	return res, nil
}

func (api *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPC_PORT))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, &UserServer{DB: *api.DB})

	log.Printf("gRPC Server started on port %s", GRPC_PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
