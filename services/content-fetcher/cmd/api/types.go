package main

import (
	"log/slog"

	cf "github.com/abyanmajid/codemore.io/services/content-fetcher/proto/content-fetcher"
)

type Service struct {
	cf.UnimplementedContentFetcherServiceServer
	Token string
	Log   *slog.Logger
}

type GitHubContent struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	GitUrl      string `json:"git_url"`
	DownloadUrl string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
}
