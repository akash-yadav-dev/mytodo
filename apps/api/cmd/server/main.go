package main

import "mytodo/apps/api/cmd/server/bootstrap"

func main() {
	// Initialize container with all dependencies
	container := bootstrap.NewContainer()

	//initialize app with container
	app := bootstrap.NewApp(container)

	router := bootstrap.NewRouter(app)

}
