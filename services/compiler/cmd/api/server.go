package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	compiler "github.com/abyanmajid/codemore.io/services/compiler/proto/compiler"
	"github.com/abyanmajid/codemore.io/services/compiler/utils"
	"google.golang.org/grpc"
)

func ListenAndServe() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	compiler.RegisterCompilerServiceServer(s, &Service{
		Log: slog.New(utils.StructuredLogHandler(os.Stdout, APP_NAME)),
	})

	log.Printf("gRPC Server started on port %s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
