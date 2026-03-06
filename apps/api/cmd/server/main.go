package main

import "mytodo/apps/api/cmd/server/bootstrap"

func main() {

	//Initialize logger
	logger := bootstrap.NewLogger()

	//Load configuration
	config, err := bootstrap.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration: ", err)
	}

	// Initialize container with all dependencies
	container := bootstrap.NewContainer(config, logger)

	//initialize app with container
	// app := bootstrap.NewApp(container)

	// router := bootstrap.NewRouter(app)

}

// Package main is the entry point for the API server application.
//
// This file contains the main function that:
// - Initializes the application configuration
// - Sets up logging and telemetry
// - Creates the dependency injection container
// - Starts the HTTP/gRPC servers
// - Handles graceful shutdown
//
// In production-grade applications, main.go typically:
// - Loads configuration from environment variables or config files
// - Initializes infrastructure (database, cache, message queues)
// - Sets up observability (metrics, tracing, logging)
// - Registers signal handlers for graceful shutdown
// - Starts background workers and cron jobs
// - Exposes health check and readiness endpoints
//
// Example structure:
//   func main() {
//       // Load configuration
//       config := bootstrap.LoadConfig()
//
//       // Initialize logger
//       logger := bootstrap.NewLogger(config)
//
//       // Create dependency container
//       container := bootstrap.NewContainer(config, logger)
//       defer container.Close()
//
//       // Initialize application
//       app := bootstrap.NewApp(container)
//
//       // Start server
//       if err := app.Start(); err != nil {
//           logger.Fatal("Failed to start server", err)
//       }
//   }
