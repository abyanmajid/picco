package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type TestCase struct {
	Label          string   `json:"label"`
	Inputs         []string `json:"inputs"`
	ExpectedOutput string   `json:"expected_output"`
}

func DecodeBase64JSON(encodedStr string) ([]TestCase, error) {
	// Step 1: Decode the base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 string: %v", err)
	}

	// Step 2: Unmarshal the JSON data into a Go data structure
	var data []TestCase
	err = json.Unmarshal(decodedBytes, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return data, nil
}
