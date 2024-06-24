package main

import (
	"context"
	"log/slog"

	"github.com/abyanmajid/codemore.io/broker/proto/compiler"
	"github.com/abyanmajid/codemore.io/broker/proto/judge"
	"github.com/abyanmajid/codemore.io/broker/proto/user"
	"google.golang.org/grpc"
)

type Service struct {
	UserEndpoint     string
	CompilerEndpoint string
	JudgeEndpoint    string
	Log              *slog.Logger
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ServiceClient interface {
	SourceCode(ctx context.Context, in *compiler.SourceCode, opts ...grpc.CallOption) (*compiler.Output, error)
}

type UserServiceClient struct {
	Client user.UserServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type CompilerServiceClient struct {
	Client compiler.CompilerServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type JudgeServiceClient struct {
	Client judge.JudgeServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserByIdRequest struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
	Xp       int32    `json:"xp"`
	IsBanned bool     `json:"is_banned"`
}

type CompileRequest struct {
	Code string   `json:"code"`
	Args []string `json:"args"`
}

type TestResult struct {
	Id             string `json:"_id"`
	Passed         bool   `json:"passed"`
	Output         string `json:"output"`
	ExpectedOutput string `json:"expected_output"`
}

type TestCase struct {
	Inputs         []string `json:"inputs"`
	ExpectedOutput string   `json:"expected_output"`
}

type RunTestsRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}
