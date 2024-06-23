package main

import (
	"context"
	"time"

	"github.com/abyanmajid/codemore.io/broker/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (api *Service) getUserServiceClient() (*GRPCClient, error) {
	api.Log.Info("Creating new gRPC client", "endpoint", api.UserEndpoint)

	conn, err := grpc.NewClient(api.UserEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		api.Log.Error("Failed to create gRPC client connection", "error", err)

		return nil, err
	}

	client := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	api.Log.Info("gRPC client created successfully")

	return &GRPCClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}
