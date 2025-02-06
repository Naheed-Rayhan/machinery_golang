package main

import (
	"fmt"
	"github.com/Naheed-Rayhan/machinery_golang/connection"
	"log"
)

func main() {
	//srv, err := connection.InitServer_v1()
	srv, err := connection.InitServer_v2()
	if err != nil {
		log.Fatal(err)
	}

	_ = srv.RegisterTasks(map[string]interface{}{
		"add": func(a, b int) error {
			fmt.Printf("Adding %d + %d = %d\n", a, b, a+b)
			return nil
		},
	})

	worker := srv.NewWorker("worker1", 20)
	if err = worker.Launch(); err != nil {
		log.Fatal(err)
	}
}
