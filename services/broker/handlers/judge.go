package handlers

import (
	"context"
	"net/http"
	"time"

	judge "github.com/abyanmajid/codemore.io/services/broker/proto/judge"
	"github.com/abyanmajid/codemore.io/services/broker/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type JudgeServiceClient struct {
	Client judge.JudgeServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type JudgeService struct {
	Endpoint string
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

func (api *JudgeService) getJudgeServiceClient() (*JudgeServiceClient, error) {

	conn, err := grpc.NewClient(api.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := judge.NewJudgeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return &JudgeServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}

func (api *JudgeService) HandleRunTests(w http.ResponseWriter, r *http.Request) {

	var requestPayload RunTestsRequest
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	client, err := api.getJudgeServiceClient()
	if err != nil {
		utils.ErrorJSON(w, err)
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
		utils.ErrorJSON(w, err)
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

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully ran tests",
		Data:    testResults,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}
