package handlers

import (
	"context"
	"net/http"
	"time"

	cf "github.com/abyanmajid/codemore.io/services/broker/proto/content-fetcher"
	"github.com/abyanmajid/codemore.io/services/broker/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ContentFetcherService struct {
	Endpoint string
}

type ContentFetcherServiceClient struct {
	Client cf.ContentFetcherServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type GetContentRequest struct {
	Path string `json:"path"`
}

func (s *ContentFetcherService) getContentFetcherServiceClient() (*ContentFetcherServiceClient, error) {

	conn, err := grpc.NewClient(s.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := cf.NewContentFetcherServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return &ContentFetcherServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}

func (s *ContentFetcherService) HandleGetContent(w http.ResponseWriter, r *http.Request) {
	var requestPayload GetContentRequest

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	client, err := s.getContentFetcherServiceClient()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	res, err := client.Client.GetContent(client.Ctx, &cf.GetContentRequest{
		Path: requestPayload.Path,
	})

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := utils.JsonResponse{
		Error:   false,
		Message: "Successfully fetched content",
		Data:    res.Data,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}
