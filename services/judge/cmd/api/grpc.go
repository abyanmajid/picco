package main

import (
	"context"
	"time"

	"github.com/abyanmajid/codemore.io/services/judge/proto/compiler"
	cf "github.com/abyanmajid/codemore.io/services/judge/proto/content-fetcher"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (api *Service) getCompilerServiceClient() (*CompilerServiceClient, error) {
	api.Log.Info("Creating new gRPC client", "endpoint", api.CompilerEndpoint)

	conn, err := grpc.NewClient(api.CompilerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		api.Log.Error("Failed to create gRPC client connection", "error", err)
		return nil, err
	}

	client := compiler.NewCompilerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	api.Log.Info("gRPC client created successfully")

	return &CompilerServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}

func (api *Service) getContentFetcherServiceClient() (*ContentFetcherServiceClient, error) {
	api.Log.Info("Creating new gRPC client", "endpoint", api.ContentFetcherEndpoint)

	conn, err := grpc.NewClient(api.ContentFetcherEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		api.Log.Error("Failed to create gRPC client connection", "error", err)
		return nil, err
	}

	client := cf.NewContentFetcherServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	api.Log.Info("gRPC client created successfully")

	return &ContentFetcherServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}
