package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/abyanmajid/codemore.io/services/broker/proto/compiler"
	"github.com/abyanmajid/codemore.io/services/broker/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CompilerService struct {
	Endpoint string
}

type CompilerServiceClient struct {
	Client compiler.CompilerServiceClient
	Conn   *grpc.ClientConn
	Ctx    context.Context
	Cancel context.CancelFunc
}

type CompileRequest struct {
	Code string   `json:"code"`
	Args []string `json:"args"`
}

func (s *CompilerService) getCompilerServiceClient() (*CompilerServiceClient, error) {

	conn, err := grpc.NewClient(s.Endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := compiler.NewCompilerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return &CompilerServiceClient{
		Client: client,
		Conn:   conn,
		Ctx:    ctx,
		Cancel: cancel,
	}, nil
}

func (s *CompilerService) compile(w http.ResponseWriter, r *http.Request, language string) {
	var requestPayload CompileRequest

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	client, err := s.getCompilerServiceClient()
	if err != nil {
		utils.ErrorJSON(w, err)
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
		utils.ErrorJSON(w, err)
		return
	}

	var responsePayload utils.JsonResponse
	responsePayload.Error = false
	responsePayload.Message = "Successfully compiled " + language + " code"
	responsePayload.Data = output

	utils.WriteJSON(w, http.StatusOK, responsePayload)
}

func (s *CompilerService) HandleCompilePython(w http.ResponseWriter, r *http.Request) {
	s.compile(w, r, "python")
}

func (s *CompilerService) HandleCompileJava(w http.ResponseWriter, r *http.Request) {
	s.compile(w, r, "java")
}

func (s *CompilerService) HandleCompileCpp(w http.ResponseWriter, r *http.Request) {
	s.compile(w, r, "cpp")
}

func (s *CompilerService) HandleCompileJavaScript(w http.ResponseWriter, r *http.Request) {
	s.compile(w, r, "javascript")
}
