package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Shreyas220/orchestrator-service/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type OrchestratorServer struct {
}

func (s *OrchestratorServer) GetUserByName(ctx context.Context, p *protos.RequestMessage) (*protos.User, error) {
	log.Printf("Receive name from from client: %s", p.GetName())
	return nil, errors.New("not implemented yet. Shreyas will implement me")

}

func main() {

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	s := OrchestratorServer{}

	protos.RegisterOrchestratorServiceServer(gs, &s)
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		os.Exit(1)
	}
	fmt.Println("listening on port 9000")
	// listen for requests
	gs.Serve(l)

}
