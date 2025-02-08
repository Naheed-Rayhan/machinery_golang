package main

import (
	"github.com/Naheed-Rayhan/machinery_golang/connection"
	"github.com/RichardKnop/machinery/v2/backends/result"
	"github.com/RichardKnop/machinery/v2/tasks"
	"log"
	"time"
)

func main() {
	//srv, err := connection.InitServer_v1()
	srv, err := connection.InitServer_v2()
	if err != nil {
		log.Fatal(err)
	}

	//_, err = srv.SendTask(&tasks.Signature{
	//	Name: "add",
	//	Args: []tasks.Arg{
	//		{Type: "int", Value: 5},
	//		{Type: "int", Value: 3},
	//	},
	//})

	// Define the time when the task should be executed
	etaTime := time.Now().Add(120 * time.Second) // 60 seconds from now
	var res *result.AsyncResult
	res, err = srv.SendTask(&tasks.Signature{
		UUID: "1234",
		Name: "create_live_classroom",
		Args: []tasks.Arg{
			{Type: "string", Value: "2021-09-01 10:00:00"},
			{Type: "string", Value: "2021-09-01 11:00:00"},
			{Type: "string", Value: "cse-101"},
			{Type: "string", Value: "teacher-1"},
		},
		ETA: &etaTime, // 60 seconds from now
	})

	// Wait for the task to complete and retrieve the result
	taskResult, err := res.Get(time.Second * 1) // Wait for 5 seconds
	if err != nil {
		log.Fatal("Failed to get task result:", err)
	}

	// Print the result on the client side
	// so the queue will be empty
	log.Println("Task result received on the client side:")
	log.Printf("Starting Time: %s\n", taskResult[0])
	log.Printf("Ending Time: %s\n", taskResult[1])
	log.Printf("Course ID: %s\n", taskResult[2])
	log.Printf("Teacher ID: %s\n", taskResult[3])

}
