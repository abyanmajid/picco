package main

import (
	"log/slog"

	compiler "github.com/abyanmajid/codemore.io/compiler/proto"
)

type Service struct {
	compiler.UnimplementedCompilerServiceServer
	Log *slog.Logger
}
