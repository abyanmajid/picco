package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cf "github.com/abyanmajid/codemore.io/services/content-fetcher/proto/content-fetcher"
)

func (api *Service) GetContent(ctx context.Context, req *cf.GetContentRequest) (*cf.GetContentResponse, error) {

	filePath := req.GetPath()
	url := GITHUB_API_CONTENTS_ENDPOINT + filePath

	api.Log.Info("Creating HTTP request", "url", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		api.Log.Error("Error creating request", "error", err)
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	if api.Token == "" {
		api.Log.Error("Authorization token is missing")
		return nil, fmt.Errorf("authorization token is missing")
	}

	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("Authorization", "Bearer "+api.Token)
	request.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	api.Log.Info("Sending HTTP request to GitHub API")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		api.Log.Error("Error making request", "error", err)
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer func() {
		resp.Body.Close()
		api.Log.Info("Response body closed")
	}()

	if resp.StatusCode != http.StatusOK {
		api.Log.Error("Received non-200 response status", "status", resp.Status)
		return nil, fmt.Errorf("error: received non-200 response status: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		api.Log.Error("Error reading response body", "error", err)
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	api.Log.Info("Successfully read response body")

	var content GitHubContent
	if err := json.Unmarshal(body, &content); err != nil {
		api.Log.Error("Error parsing JSON response", "error", err)
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}
	api.Log.Info("Successfully parsed JSON response")

	if content.Encoding == "base64" {
		decodedBytes, err := base64.StdEncoding.DecodeString(content.Content)
		if err != nil {
			api.Log.Error("Error decoding base64 content", "error", err)
			return nil, fmt.Errorf("error decoding base64 content: %v", err)
		}
		decodedContent := string(decodedBytes)
		api.Log.Info("Successfully decoded base64 content")

		return &cf.GetContentResponse{
			Mdx: decodedContent,
		}, nil
	}

	api.Log.Error("Unexpected content encoding", "encoding", content.Encoding)
	return nil, fmt.Errorf("unexpected content encoding: %v", content.Encoding)
}
