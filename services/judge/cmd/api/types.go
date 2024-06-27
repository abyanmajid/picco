package main

import (
	"context"
	"log/slog"

	compiler "github.com/abyanmajid/codemore.io/services/judge/proto/compiler"
	judge "github.com/abyanmajid/codemore.io/services/judge/proto/judge"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Service struct {
	judge.UnimplementedJudgeServiceServer
	CompilerEndpoint string
	Mongo            *mongo.Client
	Log              *slog.Logger
}

type CompilerServiceClient struct {
	Client compiler.CompilerServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type TestCase struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	TaskId         string             `bson:"task_id"`
	Inputs         []string           `bson:"inputs"`
	ExpectedOutput string             `bson:"expected_output"`
}
