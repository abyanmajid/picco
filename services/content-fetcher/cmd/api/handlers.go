package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	cf "github.com/abyanmajid/codemore.io/services/content-fetcher/proto/content-fetcher"
)

func (api *Service) GetContent(ctx context.Context, req *cf.GetContentRequest) (*cf.GetContentResponse, error) {
	filePath := req.GetPath()
	url := GITHUB_API_CONTENTS_ENDPOINT + filePath

	log.Printf("Creating HTTP request: %s\n", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	if api.Token == "" {
		log.Println("Authorization token is missing")
		return nil, fmt.Errorf("authorization token is missing")
	}

	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("Authorization", "Bearer "+api.Token)
	request.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	log.Println("Sending HTTP request to GitHub API")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("Error making request: %v\n", err)
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer func() {
		resp.Body.Close()
		log.Println("Response body closed")
	}()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-200 response status: %s\n", resp.Status)
		return nil, fmt.Errorf("error: received non-200 response status: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	log.Println("Successfully read response body")

	var content GitHubContent
	if err := json.Unmarshal(body, &content); err != nil {
		log.Printf("Error parsing JSON response: %v\n", err)
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}
	log.Println("Successfully parsed JSON response")

	log.Printf("Content encoding: %s\n", content.Encoding)
	log.Printf("Content: %s\n", content.Content)

	if content.Encoding == "base64" {
		decodedBytes, err := base64.StdEncoding.DecodeString(content.Content)
		if err != nil {
			log.Printf("Error decoding base64 content: %v\n", err)
			return nil, fmt.Errorf("error decoding base64 content: %v", err)
		}
		decodedContent := string(decodedBytes)
		log.Println("Successfully decoded base64 content")

		return &cf.GetContentResponse{
			Data: decodedContent,
		}, nil
	}

	log.Printf("Unexpected content encoding: %s\n", content.Encoding)
	return nil, fmt.Errorf("unexpected content encoding: %v", content.Encoding)
}
