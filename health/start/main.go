package main

import (
	"context"
	"health"
	"log"

	"go.temporal.io/sdk/client"
)



func main(){
	c, err := client.NewClient(client.Options{})
    if err != nil {
        log.Fatalf("Failed to create Temporal client: %v", err)
    }
    defer c.Close()


	// Start a workflow
    options := client.StartWorkflowOptions{
        ID:        "health",
        TaskQueue: "health_queue",
    }


	w,err:=c.ExecuteWorkflow(context.Background(),options,health.Wkflow)
	if err!=nil{
		log.Fatalln("Couldn't execute")
	}

	var res string
	err=w.Get(context.Background(),&res)
	if err!=nil{
		log.Fatalln("failed to open it %v",err)
	}

	log.Println(res)

}