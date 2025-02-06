package main

import (
	"github.com/Naheed-Rayhan/machinery_golang/connection"
	"github.com/RichardKnop/machinery/v2/tasks"
	"log"
)

func main() {
	//srv, err := connection.InitServer_v1()
	srv, err := connection.InitServer_v2()
	if err != nil {
		log.Fatal(err)
	}

	_, err = srv.SendTask(&tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{Type: "int", Value: 5},
			{Type: "int", Value: 3},
		},
	})
	if err != nil {
		log.Println("Error sending task:", err)
	}
}
