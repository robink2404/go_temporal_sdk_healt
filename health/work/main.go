package main

import (
	"health"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)




func main(){
	c, err := client.NewClient(client.Options{})
    if err != nil {
        log.Fatalf("Failed to create Temporal client: %v", err)
    }
    defer c.Close()

   w := worker.New(c,"health_queue",worker.Options{})


   w.RegisterWorkflow(health.Wkflow)
   w.RegisterActivity(health.Healthmeasure)

   err=w.Run(worker.InterruptCh())
   if err!=nil{
	log.Fatalln("work main.go last step %v",err)
   }

   




}