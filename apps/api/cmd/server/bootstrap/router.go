package bootstrap

import (
	"fmt"
	"net/http"
)

func SetupRouter(container *Container) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		container.Log.Info("Health check endpoint called")

		fmt.Fprint(w, "OK")
	})

	return mux
}