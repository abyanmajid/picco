package main

import (
	"context"
	"fmt"

	judge "github.com/abyanmajid/codemore.io/judge/proto"
	"github.com/google/uuid"
)

func (api *Service) CreateTestCase(ctx context.Context, req *judge.CreateTestCaseRequest) (*judge.CreateTestCaseResponse, error) {
	api.Log.Info("CreateTestCase called", "task_id", req.GetTaskId())

	// Generate a new UUID for the test case ID
	testCaseId := uuid.New().String()
	api.Log.Info("Generated new UUID for test case", "test_case_id", testCaseId)

	// Initialize the TestCase
	testCase := TestCase{
		TaskId:         req.GetTaskId(),
		TestCaseId:     testCaseId,
		HasInput:       req.GetHasInput(),
		ExpectedOutput: req.GetExpectedOutput(),
	}

	if req.GetHasInput() {
		input := req.GetInput()
		testCase.Input = &input
		api.Log.Info("Input field set for test case", "input", input)
	}

	// Insert the TestCase into MongoDB
	collection := api.Mongo.Database("testcases").Collection("testcases")
	_, err := collection.InsertOne(ctx, testCase)
	if err != nil {
		api.Log.Error("Failed to insert test case into MongoDB", "error", err)
		return nil, fmt.Errorf("failed to insert test case: %v", err)
	}
	api.Log.Info("Test case inserted into MongoDB successfully", "test_case_id", testCase.TestCaseId)

	// Create the response
	response := &judge.CreateTestCaseResponse{
		TestCase: &judge.TestCase{
			TestCaseId:     testCase.TestCaseId,
			HasInput:       testCase.HasInput,
			Input:          req.GetInput(),
			ExpectedOutput: testCase.ExpectedOutput,
		},
	}
	api.Log.Info("CreateTestCase response created successfully", "response", response)

	return response, nil
}
