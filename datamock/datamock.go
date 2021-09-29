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

type MockDatarServer struct {
}

func (s *MockDatarServer) GetMockUserData(ctx context.Context, p *protos.RequestMessage) (*protos.User, error) {
	name := p.GetName()
	var length = len([]rune(name))
	if length < 6 {
		return nil, errors.New("name less than 6 charector ")
	}
	roll := length * 10
	return &protos.User{
		Name:  name,
		Class: int64(length),
		Roll:  int64(roll),
	}, errors.New("hah")
}

func main() {

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	t := MockDatarServer{}

	protos.RegisterDataMockServiceServer(gs, &t)
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		os.Exit(1)
	}
	fmt.Println("listening on port 10000")
	// listen for requests
	gs.Serve(l)

}
