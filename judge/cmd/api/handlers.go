package main

import (
	"context"
	"fmt"

	judge "github.com/abyanmajid/codemore.io/judge/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (api *Service) GetAllTestCases(ctx context.Context, req *judge.GetAllTestCasesRequest) (*judge.GetAllTestCasesResponse, error) {
	collection := api.Mongo.Database("testcases").Collection("testcases")

	// Create a filter to match the task_id
	filter := bson.M{"task_id": req.GetTaskId()}

	// Find all documents that match the filter
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find test cases: %v", err)
	}
	defer cursor.Close(ctx)

	var testCases []TestCase
	for cursor.Next(ctx) {
		var testCase TestCase
		if err := cursor.Decode(&testCase); err != nil {
			return nil, fmt.Errorf("failed to decode test case: %v", err)
		}
		testCases = append(testCases, testCase)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	// Construct the response
	var response judge.GetAllTestCasesResponse
	for _, tc := range testCases {
		testCase := &judge.TestCase{
			TestCaseId:     tc.TestCaseId,
			HasInput:       tc.HasInput,
			ExpectedOutput: tc.ExpectedOutput,
		}
		if tc.HasInput && tc.Input != nil {
			testCase.Input = *tc.Input
		}
		response.TestCases = append(response.TestCases, testCase)
	}

	return &response, nil
}

func (api *Service) GetTestCase(ctx context.Context, req *judge.GetTestCaseRequest) (*judge.GetTestCaseResponse, error) {
	collection := api.Mongo.Database("testcases").Collection("testcases")

	// Create a filter to match the test_case_id
	filter := bson.M{"test_case_id": req.GetTestCaseId()}

	// Find the document that matches the filter
	var testCase TestCase
	err := collection.FindOne(ctx, filter).Decode(&testCase)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("test case not found: %v", err)
		}
		return nil, fmt.Errorf("failed to find test case: %v", err)
	}

	// Construct the response
	response := &judge.GetTestCaseResponse{
		TestCase: &judge.TestCase{
			TestCaseId:     testCase.TestCaseId,
			HasInput:       testCase.HasInput,
			ExpectedOutput: testCase.ExpectedOutput,
		},
	}
	if testCase.HasInput && testCase.Input != nil {
		response.TestCase.Input = *testCase.Input
	}

	return response, nil
}

func (api *Service) UpdateTestCase(ctx context.Context, req *judge.UpdateTestCaseRequest) (*judge.UpdateTestCaseResponse, error) {
	collection := api.Mongo.Database("testcases").Collection("testcases")

	// Create a filter to match the test_case_id
	filter := bson.M{"test_case_id": req.GetTestCaseId()}

	// Create an update document
	update := bson.M{
		"$set": bson.M{
			"task_id":         req.GetTaskId(),
			"has_input":       req.GetHasInput(),
			"expected_output": req.GetExpectedOutput(),
		},
	}

	if req.GetHasInput() {
		update["$set"].(bson.M)["input"] = req.GetInput()
	} else {
		update["$unset"] = bson.M{"input": ""}
	}

	// Execute the update operation
	result := collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("test case not found: %v", result.Err())
		}
		return nil, fmt.Errorf("failed to update test case: %v", result.Err())
	}

	// Decode the updated document
	var updatedTestCase TestCase
	if err := result.Decode(&updatedTestCase); err != nil {
		return nil, fmt.Errorf("failed to decode updated test case: %v", err)
	}

	// Construct the response
	response := &judge.UpdateTestCaseResponse{
		TestCase: &judge.TestCase{
			TestCaseId:     updatedTestCase.TestCaseId,
			HasInput:       updatedTestCase.HasInput,
			ExpectedOutput: updatedTestCase.ExpectedOutput,
		},
	}
	if updatedTestCase.HasInput && updatedTestCase.Input != nil {
		response.TestCase.Input = *updatedTestCase.Input
	}

	return response, nil
}

func (api *Service) DeleteTestCase(ctx context.Context, req *judge.DeleteTestCaseRequest) (*judge.DeleteTestCaseResponse, error) {
	collection := api.Mongo.Database("testcases").Collection("testcases")

	// Create a filter to match the test_case_id
	filter := bson.M{"test_case_id": req.GetTestCaseId()}

	// Execute the delete operation
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to delete test case: %v", err)
	}

	// Check if a document was actually deleted
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("test case not found")
	}

	// Construct the response
	response := &judge.DeleteTestCaseResponse{
		Success: true,
	}

	return response, nil
}

// func (api *Service) RunTests(ctx context.Context, req *judge.RunTestsRequest) (*judge.RunTestsResponse, error) {

// }
