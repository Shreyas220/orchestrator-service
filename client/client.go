package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shreyas220/orchestrator-service/protos"
	"google.golang.org/grpc"
)

//this functions returns nothing but displace the error or response from the request
func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protos.NewOrchestratorServiceClient(conn)

	//Asking for name from user
	fmt.Println("Enter Your  Name: ")
	var name string
	// Taking input from user
	fmt.Scanln(&name)

	//sending request to the server
	response, err := c.GetUserByName(context.Background(), &protos.RequestMessage{
		Name: name,
	})
	if err != nil {
		log.Fatalf("Error when calling GetUserByName: %s", err)
	}
	log.Printf("Response from server: %s", response)

}
