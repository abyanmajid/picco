package main

import (
	"context"
	"time"

	"github.com/abyanmajid/codemore.io/broker/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (api *Config) getUserServiceClient() (*GRPCClient, error) {
	conn, err := grpc.Dial("user:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	client := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return &GRPCClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}
