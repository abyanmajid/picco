package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (api *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Read the raw request body for debugging
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %w", err)
	}

	// Print the raw request body
	fmt.Printf("Raw request body: %s\n", string(body))

	// Reset the request body after reading it
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Decode the JSON from the request body
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(data)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func (api *Config) errorJSON(w http.ResponseWriter, err error, status int) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(theError)
}

func (api *Config) writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
