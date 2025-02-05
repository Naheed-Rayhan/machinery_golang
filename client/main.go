package main

import (
	"github.com/Naheed-Rayhan/machinery_golang/connection"
	"github.com/RichardKnop/machinery/v1/tasks"
	"log"
)

func main() {
	srv, err := connection.InitServer()
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
