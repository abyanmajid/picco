package main

import (
	"context"
	"log/slog"

	compiler "github.com/abyanmajid/codemore.io/services/judge/proto/compiler"
	cf "github.com/abyanmajid/codemore.io/services/judge/proto/content-fetcher"
	judge "github.com/abyanmajid/codemore.io/services/judge/proto/judge"
	"google.golang.org/grpc"
)

type Service struct {
	judge.UnimplementedJudgeServiceServer
	CompilerEndpoint       string
	ContentFetcherEndpoint string
	Log                    *slog.Logger
}

type CompilerServiceClient struct {
	Client compiler.CompilerServiceClient
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

type TestCase struct {
	Label          string   `json:"label"`
	Inputs         []string `json:"inputs"`
	ExpectedOutput string   `json:"expected_output"`
}

type TestCases struct {
	Testcases []TestCase `json:"testcases"`
}
