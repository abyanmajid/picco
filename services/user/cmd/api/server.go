package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/abyanmajid/codemore.io/services/user/internal/database"
	user "github.com/abyanmajid/codemore.io/services/user/proto"
	utils "github.com/abyanmajid/codemore.io/services/user/utils"
	"google.golang.org/grpc"
)

func ListenAndServe(conn *sql.DB) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, &Service{
		DB:  *database.New(conn),
		Log: slog.New(utils.StructuredLogHandler(os.Stdout, APP_NAME)),
	})

	log.Printf("gRPC Server started on port %s", PORT)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
