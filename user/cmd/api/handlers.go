package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/abyanmajid/codemore.io/user/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (api *Config) HandleHealth(w http.ResponseWriter, r *http.Request) {
	// payload := jsonResponse{
	// 	Error:   false,
	// 	Message: "Hello, World!",
	// }

	// _ = api.writeJSON(w, http.StatusOK, payload)
}

func (api *Config) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		AuthType string `json:"auth_type"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := api.readJSON(w, r, &requestPayload)
	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		AuthType:  requestPayload.AuthType,
		Name:      requestPayload.Name,
		Email:     requestPayload.Email,
		Password:  sql.NullString{String: requestPayload.Password, Valid: true},
		Level:     1,
		Badges:    []string{},
		IsBanned:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		api.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	api.writeJSON(w, http.StatusCreated, user)
}
