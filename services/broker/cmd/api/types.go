package main

import (
	"context"
	"log/slog"

	"github.com/abyanmajid/codemore.io/services/broker/proto/compiler"
	cf "github.com/abyanmajid/codemore.io/services/broker/proto/content-fetcher"
	"github.com/abyanmajid/codemore.io/services/broker/proto/judge"
	"github.com/abyanmajid/codemore.io/services/broker/proto/user"
	"google.golang.org/grpc"
)

type Service struct {
	UserEndpoint           string
	CompilerEndpoint       string
	JudgeEndpoint          string
	ContentFetcherEndpoint string
	Log                    *slog.Logger
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

type ContentFetcherServiceClient struct {
	Client cf.ContentFetcherServiceClient
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
	Passed         bool   `json:"passed"`
	Output         string `json:"output"`
	ExpectedOutput string `json:"expected_output"`
}

type TestCase struct {
	Inputs         []string `json:"inputs"`
	ExpectedOutput string   `json:"expected_output"`
}

type RunTestsRequest struct {
	Path     string `json:"path"`
	Code     string `json:"code"`
	Language string `json:"language"`
}

type GetContentRequest struct {
	Path string `json:"path"`
}
