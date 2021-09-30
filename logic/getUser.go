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

//this functions returns User struct from Mockdata Service
func GetFromMockdata(name string) (*protos.User, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
		return nil, err
	}
	defer conn.Close()

	c := protos.NewDataMockServiceClient(conn)

	response, err := c.GetMockUserData(context.Background(), &protos.RequestMessage{
		Name: name,
	})
	if err != nil {
		log.Fatalf("Error when calling mockdata: %s", err)
		return nil, err
	}
	log.Printf("Response from MockData: %s", response)

	return response, err

}

//input - Requestmesaage struct defined
//output - User struct defined
//GetUser implements the OrchestratorServiceServer method and returns the User
func (s *OrchestratorServer) GetUser(ctx context.Context, p *protos.RequestMessage) (*protos.User, error) {
	log.Printf("Receive name from from GetUserByName in GetUser: %s", p.GetName())
	name := p.Name
	response, err := GetFromMockdata(name)
	if err != nil {
		log.Fatalf("Error when calling Mockdata: %s", err)
		return nil, err
	}

	return response, nil

}

//starts server
func Server2() {

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	s := OrchestratorServer{}

	protos.RegisterOrchestratorServiceServer(gs, &s)
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		os.Exit(1)
	}
	fmt.Println("GetName listening on port 9001")
	// listen for requests
	gs.Serve(l)

}
