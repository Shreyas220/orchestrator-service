package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shreyas220/orchestrator-service/protos"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protos.NewOrchestratorServiceClient(conn)

	fmt.Println("Enter Your  Name: ")

	// var then variable name then variable type
	var name string

	// Taking input from user
	fmt.Scanln(&name)

	response, err := c.GetUserByName(context.Background(), &protos.RequestMessage{
		Name: name,
	})
	if err != nil {
		log.Fatalf("Error when calling GetUserByName: %s", err)
	}
	log.Printf("Response from server: %s", response)

}
