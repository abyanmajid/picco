package main

import (
	"context"
	"time"

	"github.com/abyanmajid/codemore.io/services/broker/proto/compiler"
	cf "github.com/abyanmajid/codemore.io/services/broker/proto/content-fetcher"
	"github.com/abyanmajid/codemore.io/services/broker/proto/course"
	"github.com/abyanmajid/codemore.io/services/broker/proto/judge"
	"github.com/abyanmajid/codemore.io/services/broker/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (api *Service) getUserServiceClient() (*UserServiceClient, error) {
	api.Log.Info("Creating new gRPC client", "endpoint", api.UserEndpoint)

	conn, err := grpc.NewClient(api.UserEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		api.Log.Error("Failed to create gRPC client connection", "error", err)
		return nil, err
	}

	client := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	api.Log.Info("gRPC client created successfully")

	return &UserServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}

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

func (api *Service) getJudgeServiceClient() (*JudgeServiceClient, error) {
	api.Log.Info("Creating new gRPC client", "endpoint", api.JudgeEndpoint)

	conn, err := grpc.NewClient(api.JudgeEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		api.Log.Error("Failed to create gRPC client connection", "error", err)
		return nil, err
	}

	client := judge.NewJudgeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	api.Log.Info("gRPC client created successfully")

	return &JudgeServiceClient{
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

func (api *Service) getCourseServiceClient() (*CourseServiceClient, error) {
	api.Log.Info("Creating new gRPC client", "endpoint", api.CourseEndpoint)

	conn, err := grpc.NewClient(api.CourseEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		api.Log.Error("Failed to create gRPC client connection", "error", err)
		return nil, err
	}

	client := course.NewCourseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	api.Log.Info("gRPC client created successfully")

	return &CourseServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}
