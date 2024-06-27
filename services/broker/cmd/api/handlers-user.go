package main

import (
	"net/http"

	"github.com/abyanmajid/codemore.io/services/broker/proto/user"
	"github.com/abyanmajid/codemore.io/services/broker/utils"
	"github.com/go-chi/chi/v5"
)

func (api *Service) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("Starting to handle CreateUser request")

	var requestPayload CreateUserRequest

	api.Log.Info("Reading JSON request", "handler", "HandleCreateUser")
	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.Log.Error("Error reading JSON request", "error", err)
		api.errorJSON(w, err)
		return
	}

	client, err := api.getUserServiceClient()
	if err != nil {
		api.Log.Error("Error getting user service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	api.Log.Info("Creating user", "username", requestPayload.Username, "email", requestPayload.Email)
	u, err := client.Client.CreateUser(client.Ctx, &user.CreateUserRequest{
		Username: requestPayload.Username,
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
	})

	if err != nil {
		api.Log.Error("Error creating user", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully created user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.Log.Info("User created successfully", "user_id", u.Id)
	api.writeJSON(w, http.StatusCreated, responsePayload)
}

func (api *Service) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("Starting to handle GetAllUsers request")

	client, err := api.getUserServiceClient()
	if err != nil {
		api.Log.Error("Error getting user service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	users, err := client.Client.GetAllUsers(client.Ctx, &user.GetAllUsersRequest{})
	if err != nil {
		api.Log.Error("Error fetching all users", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully fetched all users"
	responsePayload.Data = utils.DecodeMultipleProtoUsers(users.Users)

	api.Log.Info("Successfully fetched all users", "user_count", len(users.Users))
	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("Starting to handle GetUserById request")

	userID := chi.URLParam(r, "id")

	api.Log.Info("Fetching user by ID", "user_id", userID)

	client, err := api.getUserServiceClient()
	if err != nil {
		api.Log.Error("Error getting user service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	u, err := client.Client.GetUserById(client.Ctx, &user.GetUserByIdRequest{
		Id: userID,
	})

	if err != nil {
		api.Log.Error("Error fetching user by ID", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully fetched user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.Log.Info("Successfully fetched user", "user_id", u.Id)
	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("Starting to handle GetUserByEmail request")

	userEmail := chi.URLParam(r, "email")

	api.Log.Info("Fetching user by email", "email", userEmail)

	client, err := api.getUserServiceClient()
	if err != nil {
		api.Log.Error("Error getting user service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	u, err := client.Client.GetUserByEmail(client.Ctx, &user.GetUserByEmailRequest{
		Email: userEmail,
	})

	if err != nil {
		api.Log.Error("Error fetching user by email", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully fetched user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.Log.Info("Successfully fetched user", "user_id", u.Id)
	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("Starting to handle UpdateUserById request")

	userID := chi.URLParam(r, "id")

	var requestPayload UpdateUserByIdRequest
	api.Log.Info("Reading JSON request", "handler", "HandleUpdateUserById")
	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.Log.Error("Error reading JSON request", "error", err)
		api.errorJSON(w, err)
		return
	}

	client, err := api.getUserServiceClient()
	if err != nil {
		api.Log.Error("Error getting user service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	api.Log.Info("Updating user by ID", "user_id", userID)
	u, err := client.Client.UpdateUserById(client.Ctx, &user.UpdateUserByIdRequest{
		Id:       userID,
		Username: requestPayload.Username,
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
		Roles:    requestPayload.Roles,
		Xp:       requestPayload.Xp,
		IsBanned: requestPayload.IsBanned,
	})

	if err != nil {
		api.Log.Error("Error updating user", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully updated user #" + u.Id
	responsePayload.Data = utils.DecodeProtoUser(u)

	api.Log.Info("User updated successfully", "user_id", u.Id)
	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	api.Log.Info("Starting to handle DeleteUserById request")

	userID := chi.URLParam(r, "id")

	api.Log.Info("Deleting user by ID", "user_id", userID)

	client, err := api.getUserServiceClient()
	if err != nil {
		api.Log.Error("Error getting user service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	_, err = client.Client.DeleteUserById(client.Ctx, &user.DeleteUserByIdRequest{
		Id: userID,
	})

	if err != nil {
		api.Log.Error("Error deleting user", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully deleted user #" + userID

	api.Log.Info("User deleted successfully", "user_id", userID)
	api.writeJSON(w, http.StatusOK, responsePayload)
}
