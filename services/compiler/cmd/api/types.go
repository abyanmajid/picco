package main

import (
	"log/slog"

	compiler "github.com/abyanmajid/codemore.io/services/compiler/proto/compiler"
)

type Service struct {
	compiler.UnimplementedCompilerServiceServer
	Log *slog.Logger
}
