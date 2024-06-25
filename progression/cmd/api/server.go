package main

import (
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func ListenAndServe(*mongo.Client) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	// judge.RegisterJudgeServiceServer(s, &Service{
	// 	Mongo: mongoClient,
	// 	Log:   slog.New(utils.StructuredLogHandler(os.Stdout, APP_NAME)),
	// })

	log.Printf("gRPC Server started on port %s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
