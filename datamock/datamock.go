package datamock

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

// Mockserver is a gRPC server it implements the methods defined by the DataMockService interface
type MockServer struct {
}

//GetMockUserData implements the OrchestratorServiceServer method and returns the User
func (s *MockServer) GetMockUserData(ctx context.Context, p *protos.RequestMessage) (*protos.User, error) {
	log.Printf("Receive name from from GetUser in Mockdata: %s", p.GetName())
	name := p.Name

	var length = len(name)
	//checking if length is lesser than 6
	if length < 6 {
		log.Fatalf("name less than 6 charector")
		return nil, errors.New("name less than 6 charector")
	}
	//storing roll
	roll := length * 10

	User := &protos.User{
		Name:  name,
		Class: int64(length),
		Roll:  int64(roll),
	}
	return User, nil
}

func MockDataServer() {
	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the OrchestratorServiceServer server
	t := MockServer{}

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	protos.RegisterDataMockServiceServer(gs, &t)
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		os.Exit(1)
	}
	fmt.Println("Mock data listening on listening on port 10000")
	gs.Serve(l)

}
