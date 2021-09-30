package logic

import (
	"context"
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

func GetFromGetUser(name string) (*protos.User, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protos.NewOrchestratorServiceClient(conn)

	response, err := c.GetUser(context.Background(), &protos.RequestMessage{
		Name: name,
	})
	if err != nil {
		log.Fatalf("Error when calling mockdata: %s", err)
		return nil, err
	}
	log.Printf("Response from GetUser: %s", response)

	return response, nil

}

func (s *OrchestratorServer) GetUserByName(ctx context.Context, p *protos.RequestMessage) (*protos.User, error) {
	log.Printf("Receive name from from client in GetUserByName: %s", p.Name)

	response, err := GetFromGetUser(p.Name)
	if err != nil {
		log.Fatalf("Error when calling GetFromGetUserClient: %s", err)
		return nil, err
	}

	return response, nil

}

func Server1() {

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	s := OrchestratorServer{}

	protos.RegisterOrchestratorServiceServer(gs, &s)
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		os.Exit(1)
	}
	fmt.Println("GetUserByName listening on port 9000")
	// listen for requests
	gs.Serve(l)

}
