package main

import (
	"fmt"
	"log"
	"mytodo/apps/api/cmd/server/bootstrap"
)

func main() {

	fmt.Println("Starting the API server...")
	app, err := bootstrap.NewApp()

	fmt.Println("running database seeds")

	if err != nil {
		log.Fatal("Failed to start app:", err)
	}

	fmt.Println("starting the application")
	app.Start()
}
