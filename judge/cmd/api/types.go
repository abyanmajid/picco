package main

import (
	"log/slog"

	judge "github.com/abyanmajid/codemore.io/judge/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	judge.UnimplementedJudgeServiceServer
	Mongo *mongo.Client
	Log   *slog.Logger
}

type TestCase struct {
	TaskId         string  `json:"task_id"`
	TestCaseId     string  `json:"test_case_id"`
	HasInput       bool    `json:"has_input"`
	Input          *string `json:"input,omitempty"`
	ExpectedOutput string  `json:"expected_output"`
}
