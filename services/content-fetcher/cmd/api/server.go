package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	cf "github.com/abyanmajid/codemore.io/services/content-fetcher/proto/content-fetcher"
	"github.com/abyanmajid/codemore.io/services/content-fetcher/utils"
	"google.golang.org/grpc"
)

func ListenAndServe() {
	token := os.Getenv("GITHUB_PERSONAL_TOKEN")
	if token == "" {
		log.Fatalf("GITHUB_PERSONAL_TOKEN environment variable is missing")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	cf.RegisterContentFetcherServiceServer(s, &Service{
		Token: token,
		Log:   slog.New(utils.StructuredLogHandler(os.Stdout, APP_NAME)),
	})

	log.Printf("gRPC Server started on port %s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
