package main

import (
	"fmt"
	"net/http"

	"github.com/abyanmajid/codemore.io/broker/proto/judge"
	"github.com/go-chi/chi/v5"
)

func (api *Service) HandleCreateTestCase(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("HandleCreateTestCase called")

	var requestPayload TestCase
	err := api.readJSON(w, r, &requestPayload)
	if err != nil {
		api.Log.Error("Error reading JSON", "error", err)
		api.errorJSON(w, err)
		return
	}
	api.Log.Info("JSON payload read successfully", "payload", requestPayload)

	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.Log.Error("Error getting judge service client", "error", err)
		api.errorJSON(w, err)
		return
	}
	api.Log.Info("Judge service client acquired successfully")

	defer client.Conn.Close()
	defer client.Cancel()

	taskID := chi.URLParam(r, "task_id")
	if taskID == "" {
		err := fmt.Errorf("task_id is required")
		api.Log.Error("task_id missing in URL", "error", err)
		api.errorJSON(w, err)
		return
	}
	api.Log.Info("task_id retrieved from URL", "task_id", taskID)

	var input *string
	if requestPayload.HasInput {
		input = requestPayload.Input
	}

	// Ensure the input is only dereferenced when it is not nil
	var createTestCaseReq *judge.CreateTestCaseRequest
	if requestPayload.HasInput && input != nil {
		createTestCaseReq = &judge.CreateTestCaseRequest{
			TaskId:         taskID,
			HasInput:       requestPayload.HasInput,
			Input:          *input,
			ExpectedOutput: requestPayload.ExpectedOutput,
		}
	} else {
		createTestCaseReq = &judge.CreateTestCaseRequest{
			TaskId:         taskID,
			HasInput:       requestPayload.HasInput,
			ExpectedOutput: requestPayload.ExpectedOutput,
		}
	}

	t, err := client.Client.CreateTestCase(client.Ctx, createTestCaseReq)
	if err != nil {
		api.Log.Error("Error creating test case", "error", err)
		api.errorJSON(w, err)
		return
	}
	api.Log.Info("Test case created successfully", "test_case_id", t.TestCase.TestCaseId)

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully created test case #" + t.TestCase.TestCaseId
	responsePayload.Data = t.TestCase

	api.writeJSON(w, http.StatusCreated, responsePayload)
	api.Log.Info("Response sent successfully", "response_payload", responsePayload)
}

func (api *Service) HandleGetAllTestCases(w http.ResponseWriter, r *http.Request) {

}

func (api *Service) HandleGetTestCase(w http.ResponseWriter, r *http.Request) {

}

func (api *Service) HandleUpdateTestCase(w http.ResponseWriter, r *http.Request) {

}

func (api *Service) HandleDeleteTestCase(w http.ResponseWriter, r *http.Request) {

}

func (api *Service) HandleDeleteAllTestCases(w http.ResponseWriter, r *http.Request) {

}

func (api *Service) HandleRunTests(w http.ResponseWriter, r *http.Request) {

}
