package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brittonhayes/minikv/minikv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	if err := minikv.Start(port, []minikv.Route{
		minikv.NewRoute("/status/health", healthCheckHandler, http.MethodGet),
		minikv.NewRoute("/{key}", getKey, http.MethodGet),
		minikv.NewRoute("/{key}", postValue, http.MethodPost),
	}); err != nil {
		log.Fatalln("Server failed", err)
	}
}
