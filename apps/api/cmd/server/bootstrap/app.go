package bootstrap

import (
	"fmt"
	"log"
	"mytodo/apps/api/internal/auth/interfaces/grpc/pb"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	Router http.Handler
	Logger Logger
	GrpcServer *grpc.Server
	Container  *Container
}

func NewApp() (*App, error) {

	config := LoadConfig()

	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}

	container, err := NewContainer(logger, config)
	if err != nil {
		return nil, err
	}

	router := SetupRouter(container)

	// Setup gRPC server
	grpcServer := grpc.NewServer()
	
	// Register services
	pb.RegisterAuthServiceServer(grpcServer, container.AuthGrpcServer)
	
	// Enable reflection for development (allows tools like grpcurl to introspect services)
	reflection.Register(grpcServer)

	return &App{
		Router: router,
		Logger: logger,
			GrpcServer: grpcServer,
			Container:  container,
	}, nil
}

func (a *App) Start() {
	// Start gRPC server in a goroutine
	go func() {
		grpcPort := ":50051"
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen on %s: %v", grpcPort, err)
		}

		a.Logger.Info(fmt.Sprintf("Starting gRPC server on %s", grpcPort))
		if err := a.GrpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	a.Logger.Info("Starting HTTP server on :8080")

	err := http.ListenAndServe(":8080", a.Router)

	if err != nil {
		panic(fmt.Sprintf("Server failed: %v", err))
	}
}
