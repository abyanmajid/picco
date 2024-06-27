package main

import (
	"net/http"

	judge "github.com/abyanmajid/codemore.io/services/broker/proto/judge"
)

func (api *Service) HandleRunTests(w http.ResponseWriter, r *http.Request) {

	var requestPayload RunTestsRequest
	err := api.readJSON(w, r, &requestPayload)
	if err != nil {
		api.Log.Error("Failed to read JSON", "error", err)
		api.errorJSON(w, err)
		return
	}

	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.Log.Error("Failed to get JudgeService client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	results, err := client.Client.RunTests(client.Ctx, &judge.RunTestsRequest{
		Path:     requestPayload.Path,
		Code:     requestPayload.Code,
		Language: requestPayload.Language,
	})

	if err != nil {
		api.Log.Error("Failed to run tests", "error", err)
		api.errorJSON(w, err)
		return
	}

	var testResults []TestResult
	for _, result := range results.Results {
		testResults = append(testResults, TestResult{
			Passed:         result.Passed,
			Output:         result.Output,
			ExpectedOutput: result.ExpectedOutput,
		})
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully ran tests",
		Data:    testResults,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}
