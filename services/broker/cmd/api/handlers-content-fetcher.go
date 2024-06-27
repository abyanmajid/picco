package main

import (
	"net/http"

	cf "github.com/abyanmajid/codemore.io/services/broker/proto/content-fetcher"
)

func (api *Service) HandleGetContent(w http.ResponseWriter, r *http.Request) {
	var requestPayload GetContentRequest

	err := api.readJSON(w, r, &requestPayload)
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getContentFetcherServiceClient()
	if err != nil {
		api.Log.Error("Failed to get Content-Fetcher client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	res, err := client.Client.GetContent(client.Ctx, &cf.GetContentRequest{
		Path: requestPayload.Path,
	})

	if err != nil {
		api.Log.Error("Failed to get content", "error", err)
		api.errorJSON(w, err)
		return
	}

	responsePayload := JsonResponse{
		Error:   false,
		Message: "Successfully fetched content",
		Data:    res.Data,
	}

	api.writeJSON(w, http.StatusOK, responsePayload)
}
