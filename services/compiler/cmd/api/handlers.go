package main

import (
	"context"
	"fmt"
	"log/slog"

	compiler "github.com/abyanmajid/codemore.io/services/compiler/proto/compiler"
)

func (api *Service) Execute(ctx context.Context, req *compiler.SourceCode) (*compiler.Output, error) {
	language := req.GetLanguage()
	code := req.GetCode()
	args := req.GetArgs()

	api.Log.Info("Received compile request", slog.String("language", language))

	var output string
	var err error

	switch language {
	case "python", "python3":
		output, err = executePython(code, args)
	case "java":
		output, err = executeJava(code, args)
	case "cpp", "c++":
		output, err = executeCpp(code, args)
	case "javascript", "js":
		output, err = executeJavaScript(code, args)
	default:
		api.Log.Error("Unsupported language", slog.String("language", language))
		return nil, fmt.Errorf("unsupported language: %s", language)
	}

	if err != nil {
		api.Log.Error("Error executing code", slog.String("error", err.Error()))
		return nil, err
	}

	api.Log.Info("Successfully executed code")
	res := &compiler.Output{
		Output: output,
	}

	return res, nil
}
