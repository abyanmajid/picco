package main

import (
	"errors"
	"net/http"
)

func (api *Config) HandleUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload UserRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "create-user":
		api.CreateUser(w, requestPayload.Data)
	case "login":
		api.Login(w, requestPayload.Data)
	case "refresh":
		api.Refresh(w, requestPayload.Data)
	case "logout":
		api.Logout(w, requestPayload.Data)
	default:
		api.errorJSON(w, errors.New("unknown action"))
	}
}
