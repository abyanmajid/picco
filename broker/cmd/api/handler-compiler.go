package main

import (
	"net/http"

	"github.com/abyanmajid/codemore.io/broker/proto/compiler"
)

func (api *Service) HandleCompilePython(w http.ResponseWriter, r *http.Request) {
	var requestPayload CompileRequest

	err := api.readJSON(w, r, &requestPayload)

	if err != nil {
		api.Log.Error("Error reading JSON request", "error", err)
		api.errorJSON(w, err)
		return
	}

	client, err := api.getCompilerServiceClient()
	if err != nil {
		api.Log.Error("Error getting user service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	output, err := client.Client.Execute(client.Ctx, &compiler.SourceCode{
		Code:     requestPayload.Code,
		Language: "python",
		Args:     requestPayload.Args,
	})

	if err != nil {
		api.Log.Error("Error creating user", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully compiled python code"
	responsePayload.Data = output

	api.writeJSON(w, http.StatusOK, responsePayload)
}
