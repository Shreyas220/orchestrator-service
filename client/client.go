package main

import (
	"context"
	"log"

	"github.com/Shreyas220/orchestrator-service/protos"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protos.NewOrchestratorServiceClient(conn)

	response, err := c.GetUserByName(context.Background(), &protos.RequestMessage{
		Name: "Shreyas",
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)

}
