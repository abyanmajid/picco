package main

import (
	"net/http"

	"github.com/abyanmajid/codemore.io/broker/user"
	utils "github.com/abyanmajid/codemore.io/broker/utils"
)

func (api *Config) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload CreateUserRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getUserServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	u, err := client.Client.CreateUser(client.Ctx, &user.CreateUserRequest{
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
	responsePayload.Message = "Successfully created user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.writeJSON(w, http.StatusCreated, responsePayload)
}

func (api *Config) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	client, err := api.getUserServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	users, err := client.Client.GetAllUsers(client.Ctx, &user.GetAllUsersRequest{})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully fetched all users"
	responsePayload.Data = utils.DecodeMultipleProtoUsers(users.Users)

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Config) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	var requestPayload GetUserByIdRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getUserServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	u, err := client.Client.GetUserById(client.Ctx, &user.GetUserByIdRequest{
		Id: requestPayload.Id,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully fetched user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Config) HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var requestPayload GetUserByEmailRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getUserServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	u, err := client.Client.GetUserByEmail(client.Ctx, &user.GetUserByEmailRequest{
		Email: requestPayload.Email,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully fetched user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Config) HandleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	var requestPayload UpdateUserByIdRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getUserServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	u, err := client.Client.UpdateUserById(client.Ctx, &user.UpdateUserByIdRequest{
		Id:       requestPayload.Id,
		Username: requestPayload.Username,
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
		Roles:    requestPayload.Roles,
		Xp:       requestPayload.Xp,
		IsBanned: requestPayload.IsBanned,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully updated user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Config) HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	var requestPayload DeleteUserByIdRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	client, err := api.getUserServiceClient()
	if err != nil {
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	msg, err := client.Client.DeleteUserById(client.Ctx, &user.DeleteUserByIdRequest{
		Id: requestPayload.Id,
	})

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Data = msg.Message

	api.writeJSON(w, http.StatusNoContent, responsePayload)
}
