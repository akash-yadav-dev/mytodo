package bootstrap

import (
	"fmt"
	"net/http"
)

type App struct {
	Router http.Handler
	Logger Logger
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

	return &App{
		Router: router,
		Logger: logger,
	}, nil
}

func (a *App) Start() {

	a.Logger.Info("Starting HTTP server on :8080")

	err := http.ListenAndServe(":8080", a.Router)

	if err != nil {
		panic(fmt.Sprintf("Server failed: %v", err))
	}
}
