package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
		api.createUser(w, requestPayload.Data)
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
