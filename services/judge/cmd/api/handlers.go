package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abyanmajid/codemore.io/services/judge/proto/compiler"
	cf "github.com/abyanmajid/codemore.io/services/judge/proto/content-fetcher"
	judge "github.com/abyanmajid/codemore.io/services/judge/proto/judge"
	utils "github.com/abyanmajid/codemore.io/services/judge/utils"
)

func (api *Service) GetTestCases(ctx context.Context, req *judge.GetTestCasesRequest) (*judge.GetTestCasesResponse, error) {
	path := req.GetPath()

	client, err := api.getContentFetcherServiceClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get compiler service client: %v", err)
	}
	defer client.Conn.Close()
	defer client.Cancel()

	res, err := client.Client.GetContent(client.Ctx, &cf.GetContentRequest{
		Path: path,
	})

	jsonTests, err := utils.DecodeBase64JSON(res.Data)
	if err != nil {
		return nil, err
	}

	testCases := make([]*judge.TestCase, len(jsonTests))
	for i, testCase := range jsonTests {
		testCases[i] = &judge.TestCase{
			Label:          testCase.Label,
			Inputs:         testCase.Inputs,
			ExpectedOutput: testCase.ExpectedOutput,
		}
	}

	return &judge.GetTestCasesResponse{
		Testcases: testCases,
	}, nil
}

func (api *Service) RunTests(ctx context.Context, req *judge.RunTestsRequest) (*judge.RunTestsResponse, error) {
	code := req.GetCode()
	language := req.GetLanguage()
	path := req.GetPath()

	testCasesResponse, err := api.GetTestCases(ctx, &judge.GetTestCasesRequest{
		Path: path,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get test cases: %v", err)
	}

	client, err := api.getCompilerServiceClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get compiler service client: %v", err)
	}
	defer client.Conn.Close()
	defer client.Cancel()

	var results []*judge.TestResult

	for _, testCase := range testCasesResponse.GetTestcases() {

		inputs, err := json.Marshal(testCase.GetInputs())
		if err != nil {
			return nil, err
		}

		// Communicate with the execution microservice
		output, err := client.Client.Execute(ctx, &compiler.SourceCode{
			Code:     code,
			Language: language,
			Args:     []string{string(inputs)},
		})
		if err != nil {
			return nil, err
		}

		result := &judge.TestResult{
			Label:          testCase.GetLabel(),
			Passed:         output.GetOutput() == testCase.GetExpectedOutput(),
			Output:         output.GetOutput(),
			ExpectedOutput: testCase.GetExpectedOutput(),
		}

		results = append(results, result)
	}

	return &judge.RunTestsResponse{Results: results}, nil
}
