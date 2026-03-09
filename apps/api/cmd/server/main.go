package main

import (
	"log"
	"mytodo/apps/api/cmd/server/bootstrap"
)

func main() {

	app, err := bootstrap.NewApp()

	if err != nil {
		log.Fatal("Failed to start app:", err)
	}

	app.Start()
}
