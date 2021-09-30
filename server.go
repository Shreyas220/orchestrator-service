package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Shreyas220/orchestrator-service/datamock"
	"github.com/Shreyas220/orchestrator-service/logic"
	"github.com/Shreyas220/orchestrator-service/protos"
)

type OrchestratorServer struct {
}

func (s *OrchestratorServer) GetUserByName(ctx context.Context, p *protos.RequestMessage) (*protos.User, error) {
	log.Printf("Receive name from from client: %s", p.GetName())
	return nil, errors.New("not implemented yet. Shreyas will implement me")

}

func main() {
	fmt.Println("....starting the three servers")
	//starting GetUserByName
	go logic.Server1()
	//starting GetName
	go logic.Server2()
	//starting mock dataservice
	datamock.MockDataServer()

}
