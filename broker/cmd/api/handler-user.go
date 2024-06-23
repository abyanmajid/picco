package main

import (
	"net/http"
	"time"

	"github.com/abyanmajid/codemore.io/broker/user"
)

func (api *Config) CreateUser(w http.ResponseWriter, requestPayload UserPayload) {
	client, err := api.getUserServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	_, err = client.Client.CreateUser(client.Ctx, &user.CreateUserRequest{
		Username: requestPayload.Username,
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully created user #" + requestPayload.ID

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Config) Login(w http.ResponseWriter, requestPayload UserPayload) {

	client, err := api.getUserServiceClient()

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	token, err := client.Client.Login(client.Ctx, &user.LoginRequest{
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully logged in user #" + requestPayload.ID
	responsePayload.Data = LoginResponse{
		Token: token.Token,
	}

	api.writeJSON(w, http.StatusAccepted, responsePayload)
}

func (api *Config) Refresh(w http.ResponseWriter, requestPayload UserPayload) {
	client, err := api.getUserServiceClient()

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	_, err = client.Client.Refresh(client.Ctx, &user.RefreshRequest{
		RefreshToken: requestPayload.RefreshToken,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully refreshed JWT token for user #" + requestPayload.ID
}

func (api *Config) Logout(w http.ResponseWriter, requestPayload UserPayload) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}
