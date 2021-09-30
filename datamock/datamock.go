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

type MockServer struct {
}

func (s *MockServer) GetMockUserData(ctx context.Context, p *protos.RequestMessage) (*protos.User, error) {
	log.Printf("Receive name from from GetUser in Mockdata: %s", p.GetName())
	name := p.Name
	var length = len(name)
	if length < 6 {
		log.Fatalf("name less than 6 charector")
		return nil, errors.New("name less than 6 charector")
	}
	roll := length * 10
	User := &protos.User{
		Name:  name,
		Class: int64(length),
		Roll:  int64(roll),
	}
	return User, nil
}

func MockDataServer() {

	gs := grpc.NewServer()

	t := MockServer{}

	protos.RegisterDataMockServiceServer(gs, &t)
	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		os.Exit(1)
	}
	fmt.Println("Mock data listening on listening on port 10000")
	gs.Serve(l)

}
