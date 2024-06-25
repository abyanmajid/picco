package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abyanmajid/codemore.io/judge/proto/compiler"
	judge "github.com/abyanmajid/codemore.io/judge/proto/judge"
	"github.com/abyanmajid/codemore.io/judge/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *Service) CreateTestCase(ctx context.Context, req *judge.CreateTestCaseRequest) (*judge.TestCase, error) {
	testCase := TestCase{
		TaskId:         req.GetTaskId(),
		Inputs:         req.GetInputs(),
		ExpectedOutput: req.GetExpectedOutput(),
	}

	collection := api.Mongo.Database("db").Collection("testcases")
	doc, err := collection.InsertOne(ctx, testCase)
	if err != nil {
		return nil, err
	}

	insertedID, err := utils.ConvertToObjectIDString(doc.InsertedID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	return &judge.TestCase{
		XId:            insertedID,
		TaskId:         testCase.TaskId,
		Inputs:         testCase.Inputs,
		ExpectedOutput: testCase.ExpectedOutput,
	}, nil
}

func (api *Service) GetAllTestCases(ctx context.Context, req *judge.GetAllTestCasesRequest) (*judge.GetAllTestCasesResponse, error) {
	collection := api.Mongo.Database("db").Collection("testcases")

	filter := bson.M{"task_id": req.GetTaskId()}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var testCases []TestCase
	for cursor.Next(ctx) {
		var testCase TestCase
		if err := cursor.Decode(&testCase); err != nil {
			return nil, err
		}
		testCases = append(testCases, testCase)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	res := &judge.GetAllTestCasesResponse{}
	for _, tc := range testCases {
		testCase := &judge.TestCase{
			XId:            tc.Id.Hex(),
			TaskId:         tc.TaskId,
			Inputs:         tc.Inputs,
			ExpectedOutput: tc.ExpectedOutput,
		}
		res.Testcases = append(res.Testcases, testCase)
	}

	return res, nil
}

func (api *Service) GetTestCase(ctx context.Context, req *judge.GetTestCaseRequest) (*judge.TestCase, error) {
	collection := api.Mongo.Database("db").Collection("testcases")

	objectID, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	doc := collection.FindOne(ctx, filter)

	var testCase TestCase
	if err := doc.Decode(&testCase); err != nil {
		return nil, err
	}

	return &judge.TestCase{
		XId:            req.GetXId(),
		TaskId:         testCase.TaskId,
		Inputs:         testCase.Inputs,
		ExpectedOutput: testCase.ExpectedOutput,
	}, nil
}

func (api *Service) UpdateTestCase(ctx context.Context, req *judge.UpdateTestCaseRequest) (*judge.TestCase, error) {
	collection := api.Mongo.Database("db").Collection("testcases")

	// Convert the string ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// Create the update fields
	update := bson.M{
		"$set": bson.M{
			"inputs":          req.GetInputs(),
			"expected_output": req.GetExpectedOutput(),
		},
	}

	// Create the filter with the ObjectID
	filter := bson.M{"_id": objectID}

	// Perform the update
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update test case: %v", err)
	}

	// Check if the update was successful
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("test case not found")
	}

	// Return the updated test case
	return &judge.TestCase{
		XId:            req.GetXId(),
		Inputs:         req.GetInputs(),
		ExpectedOutput: req.GetExpectedOutput(),
	}, nil
}

func (api *Service) DeleteTestCase(ctx context.Context, req *judge.DeleteTestCaseRequest) (*judge.DeleteTestCaseResponse, error) {
	collection := api.Mongo.Database("db").Collection("testcases")
	objectID, err := primitive.ObjectIDFromHex(req.GetXId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &judge.DeleteTestCaseResponse{}, nil
}

func (api *Service) RunTests(ctx context.Context, req *judge.RunTestsRequest) (*judge.RunTestsResponse, error) {
	code := req.GetCode()
	language := req.GetLanguage()
	taskId := req.GetTaskId()

	api.Log.Info("Running tests", "task_id", taskId, "language", language)

	testCasesResponse, err := api.GetAllTestCases(ctx, &judge.GetAllTestCasesRequest{
		TaskId: taskId,
	})
	if err != nil {
		api.Log.Error("Failed to get test cases", "task_id", taskId, "error", err)
		return nil, fmt.Errorf("failed to get test cases: %v", err)
	}

	client, err := api.getCompilerServiceClient()
	if err != nil {
		api.Log.Error("Failed to get compiler service client", "task_id", taskId, "error", err)
		return nil, fmt.Errorf("failed to get compiler service client: %v", err)
	}
	defer client.Conn.Close()
	defer client.Cancel()

	var results []*judge.TestResult

	for _, testCase := range testCasesResponse.GetTestcases() {
		api.Log.Info("Executing test case", "test_case_id", testCase.GetXId())

		inputs, err := json.Marshal(testCase.GetInputs())
		if err != nil {
			api.Log.Error("Failed to marshal inputs", "test_case_id", testCase.GetXId(), "error", err)
			return nil, fmt.Errorf("failed to marshal inputs for test case %s: %v", testCase.GetXId(), err)
		}

		// Communicate with the execution microservice
		output, err := client.Client.Execute(ctx, &compiler.SourceCode{
			Code:     code,
			Language: language,
			Args:     []string{string(inputs)},
		})
		if err != nil {
			api.Log.Error("Failed to execute test case", "test_case_id", testCase.GetXId(), "error", err)
			return nil, fmt.Errorf("failed to execute test case %s: %v", testCase.GetXId(), err)
		}

		result := &judge.TestResult{
			XId:            testCase.GetXId(),
			Passed:         output.GetOutput() == testCase.GetExpectedOutput(),
			Output:         output.GetOutput(),
			ExpectedOutput: testCase.GetExpectedOutput(),
		}

		results = append(results, result)
	}

	api.Log.Info("Successfully ran tests", "task_id", taskId)

	return &judge.RunTestsResponse{Results: results}, nil
}
