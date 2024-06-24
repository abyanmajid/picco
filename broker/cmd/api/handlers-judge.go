package main

import (
	"net/http"

	judge "github.com/abyanmajid/codemore.io/broker/proto/judge"
	"github.com/go-chi/chi/v5"
)

func (api *Service) HandleRunTests(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "task_id")

	api.Log.Info("Received RunTests request", "task_id", taskId)

	var requestPayload RunTestsRequest
	err := api.readJSON(w, r, &requestPayload)
	if err != nil {
		api.Log.Error("Failed to read JSON", "error", err)
		api.errorJSON(w, err)
		return
	}

	api.Log.Info("Successfully parsed request payload", "task_id", taskId)

	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.Log.Error("Failed to get JudgeService client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	api.Log.Info("Successfully connected to JudgeService", "task_id", taskId)

	results, err := client.Client.RunTests(client.Ctx, &judge.RunTestsRequest{
		TaskId:   taskId,
		Code:     requestPayload.Code,
		Language: requestPayload.Language,
	})

	if err != nil {
		api.Log.Error("Failed to run tests", "error", err)
		api.errorJSON(w, err)
		return
	}

	api.Log.Info("Successfully ran tests", "task_id", taskId)

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully ran tests",
		Data:    results.Results,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
	api.Log.Info("Response sent", "task_id", taskId)
}

func (api *Service) HandleCreateTestCase(w http.ResponseWriter, r *http.Request) {
	var requestPayload TestCase
	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	t, err := client.Client.CreateTestCase(client.Ctx, &judge.CreateTestCaseRequest{
		TaskId:         chi.URLParam(r, "task_id"),
		Inputs:         requestPayload.Inputs,
		ExpectedOutput: requestPayload.ExpectedOutput,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully created test case #" + t.XId,
		Data:    t,
	}

	api.writeJSON(w, http.StatusCreated, responsePayload)
}

func (api *Service) HandleGetAllTestCases(w http.ResponseWriter, r *http.Request) {
	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	res, err := client.Client.GetAllTestCases(client.Ctx, &judge.GetAllTestCasesRequest{
		TaskId: chi.URLParam(r, "task_id"),
	})
	if err != nil {
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully fetched test cases",
		Data:    res.Testcases,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleGetTestCase(w http.ResponseWriter, r *http.Request) {
	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	t, err := client.Client.GetTestCase(client.Ctx, &judge.GetTestCaseRequest{
		XId: chi.URLParam(r, "test_case_id"),
	})
	if err != nil {
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully fetched test cases",
		Data:    t,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleUpdateTestCase(w http.ResponseWriter, r *http.Request) {
	var requestPayload TestCase
	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	t, err := client.Client.UpdateTestCase(client.Ctx, &judge.UpdateTestCaseRequest{
		XId:            chi.URLParam(r, "test_case_id"),
		Inputs:         requestPayload.Inputs,
		ExpectedOutput: requestPayload.ExpectedOutput,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully updated test case #" + t.XId,
		Data:    t,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleDeleteTestCase(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "test_case_id")

	client, err := api.getJudgeServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	_, err = client.Client.DeleteTestCase(client.Ctx, &judge.DeleteTestCaseRequest{
		XId: id,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully deleted test case #" + id,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}
