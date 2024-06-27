package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

type TestCase struct {
	Label          string   `json:"label"`
	Inputs         []string `json:"inputs"`
	ExpectedOutput string   `json:"expected_output"`
}

func DecodeJSON(encodedStr string) ([]TestCase, error) {
	log.Println("Unmarshalling JSON data")
	var data []TestCase
	err := json.Unmarshal([]byte(encodedStr), &data)
	if err != nil {
		log.Printf("Error unmarshalling JSON data: %v\n", err)
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	log.Printf("Successfully unmarshalled JSON data: %d test cases\n", len(data))
	return data, nil
}
