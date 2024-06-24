package main

import (
	"net/http"

	"github.com/abyanmajid/codemore.io/broker/proto/compiler"
)

func (api *Service) compile(w http.ResponseWriter, r *http.Request, language string) {
	var requestPayload CompileRequest

	err := api.readJSON(w, r, &requestPayload)
	if err != nil {
		api.Log.Error("Error reading JSON request", "error", err)
		api.errorJSON(w, err)
		return
	}

	client, err := api.getCompilerServiceClient()
	if err != nil {
		api.Log.Error("Error getting compiler service client", "error", err)
		api.errorJSON(w, err)
		return
	}

	defer client.Conn.Close()
	defer client.Cancel()

	output, err := client.Client.Execute(client.Ctx, &compiler.SourceCode{
		Code:     requestPayload.Code,
		Language: language,
		Args:     requestPayload.Args,
	})

	if err != nil {
		api.Log.Error("Error compiling user-submitted code", "error", err)
		api.errorJSON(w, err)
		return
	}

	var responsePayload JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully compiled " + language + " code"
	responsePayload.Data = output

	api.writeJSON(w, http.StatusOK, responsePayload)
}

func (api *Service) HandleCompilePython(w http.ResponseWriter, r *http.Request) {
	api.compile(w, r, "python")
}

func (api *Service) HandleCompileJava(w http.ResponseWriter, r *http.Request) {
	api.compile(w, r, "java")
}

func (api *Service) HandleCompileCpp(w http.ResponseWriter, r *http.Request) {
	api.compile(w, r, "cpp")
}

func (api *Service) HandleCompileJavaScript(w http.ResponseWriter, r *http.Request) {
	api.compile(w, r, "javascript")
}
