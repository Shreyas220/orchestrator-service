package main

import (
	"fmt"

	"github.com/Shreyas220/orchestrator-service/datamock"
	"github.com/Shreyas220/orchestrator-service/logic"
)

func main() {
	fmt.Println("....starting the three servers")
	//starting GetUserByName
	go logic.Server1()
	//starting GetName
	go logic.Server2()
	//starting mock dataservice
	datamock.MockDataServer()

}
