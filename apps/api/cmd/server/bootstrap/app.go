package bootstrap

type App struct {
	container *Container
	logger    Logger
}

// Package bootstrap handles application initialization.
//
// This file defines the main application struct and lifecycle management.
//
// In production-grade applications, app initialization typically includes:
// - HTTP/gRPC server setup
// - Route registration
// - Middleware configuration
// - Background worker initialization
// - Health check endpoints
// - Graceful shutdown handling
// - Signal handling (SIGTERM, SIGINT)
//
// Example structure:
//   type App struct {
//       config    *Config
//       logger    Logger
//       container *Container
//       httpServer *http.Server
//       grpcServer *grpc.Server
//   }
//
// Example methods:
//   func (a *App) Start() error {
//       // Start HTTP server
//       go func() {
//           if err := a.httpServer.ListenAndServe(); err != nil {
//               a.logger.Fatal("HTTP server failed", err)
//           }
//       }()
//
//       // Wait for shutdown signal
//       a.waitForShutdown()
//       return nil
//   }
//
//   func (a *App) waitForShutdown() {
//       quit := make(chan os.Signal, 1)
//       signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//       <-quit
//
//       a.logger.Info("Shutting down server...")
//       ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//       defer cancel()
//
//       if err := a.httpServer.Shutdown(ctx); err != nil {
//           a.logger.Error("Server shutdown error", err)
//       }
//   }

func NewApp(container *Container, logger Logger) *App {
	return &App{
		container: container,
		logger:    logger,
	}
}

func (a *App) GetAuthHandler() *auth.AuthHandler {
	return a.container.AuthHandler
}

func (a *App) GetUserHandler() *users.UserHandler {
	return a.container.UserHandler
}

func (a *App) GetProjectHandler() *projects.ProjectHandler {
	return a.container.ProjectHandler
}

func (a *App) GetIssueHandler() *issues.IssueHandler {
	return a.container.IssueHandler
}

func (a *App) GetLogger() Logger {
	return a.logger
}
