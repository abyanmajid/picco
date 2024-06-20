package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/abyanmajid/codemore.io/broker/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (api *Config) HandleHealth(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hello, World!",
	}

	_ = api.writeJSON(w, http.StatusOK, payload)
}

func (api *Config) HandleUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload UserRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "create-user":
		api.CreateUserViaGRPC(w, requestPayload.Data)
	default:
		api.errorJSON(w, errors.New("unknown action"))
	}

}

func (api *Config) createUser(w http.ResponseWriter, userData UserPayload) {
	jsonData, _ := json.Marshal(userData)

	requestEndpoint := api.UserEndpoint + "/users"

	request, err := http.NewRequest("POST", requestEndpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("STATUS CODE:", resp.StatusCode)

	if resp.StatusCode != http.StatusCreated {
		api.errorJSON(w, errors.New("error calling user service"))
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Successfully created user #" + userData.ID

	api.writeJSON(w, http.StatusOK, payload)
}

func (api *Config) CreateUserViaGRPC(w http.ResponseWriter, requestPayload UserPayload) {

	conn, err := grpc.Dial("user:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer conn.Close()

	c := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.CreateUser(ctx, &user.CreateUserRequest{
		AuthType: requestPayload.AuthType,
		Name:     requestPayload.Name,
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
	})
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Successfully created user #" + requestPayload.ID + " via GRPC"

	api.writeJSON(w, http.StatusOK, payload)
}
