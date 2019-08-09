// +build integration contract

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"

	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pact-foundation/pact-go/types"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/service"
	"google.golang.org/grpc"
)

var grpcAddress = "localhost:5001"
var httpAddress = "localhost:5002"

type toDoImplStub struct {
}

func (r toDoImplStub) List(limit int32, notCompleted bool) ([]*pb.Todo, error) {
	return nil, nil
}

func (r toDoImplStub) Insert(items *pb.Todo) error {
	return nil
}

func (r toDoImplStub) Get(id string) (*pb.Todo, error) {
	return nil, nil
}

func (r toDoImplStub) Delete(id string) error {
	return nil
}

func startServer() {
	todoRep := &toDoImplStub{}
	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, service.ToDo{ToDoRepo: todoRep})

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal("can not listen tcp grpcAddress ", grpcAddress, " ", err)
	}

	log.Printf("Serving GRPC at %s.\n", grpcAddress)
	go s.Serve(lis)

	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Couldn't contact grpc server")
	}

	mux := grpc_runtime.NewServeMux()
	err = pb.RegisterTodoServiceHandler(context.Background(), mux, conn)
	if err != nil {
		panic("Cannot serve http api")
	}
	log.Printf("Serving http at %s.\n", httpAddress)
	err = http.ListenAndServe(httpAddress, mux)
}

func TestToDoService(t *testing.T) {
	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../consumer/pacts", dir)
	go startServer()

	pact := &dsl.Pact{
		Consumer: "ToDoConsumer",
		Provider: "ToDoService",
	}

	// Verify the Provider using the locally saved Pact Files
	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://" + httpAddress,
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/todoconsumer-todoservice.json", pactDir))},
		StateHandlers: types.StateHandlers{
			// Setup any state required by the test
			// in this case, we ensure there is a "user" in the system
			"state handler": func() error {
				//lastName = "crickets"
				return nil
			},
		},
	})

}
