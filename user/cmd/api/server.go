package main

import (
	"fmt"
	"log"
	"net"

	user "github.com/abyanmajid/codemore.io/user/proto"
	"google.golang.org/grpc"
)

func (api *Config) ListenAndServe() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, &UserServer{DB: *api.DB})

	log.Printf("gRPC Server started on port %s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
