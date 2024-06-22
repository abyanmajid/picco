package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func (api *Config) HandleHealth(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hello, World!",
	}

	api.writeJSON(w, http.StatusOK, payload)
}
