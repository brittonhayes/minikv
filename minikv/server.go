package minikv

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	path    string
	handler http.HandlerFunc
	method  string
}

// NewRoute initializes a new Route struct
func NewRoute(path string, fn http.HandlerFunc, m string) Route {
	return Route{
		path:    path,
		handler: fn,
		method:  m,
	}
}

// Start begins the http server for the KV store
// and returns an error if anything goes wrong
func Start(address string, routes []Route) error {
	ctx, cancel := listenForShutdown()
	router := mux.NewRouter()
	for _, r := range routes {
		router.Handle(r.path, r.handler).Methods(r.method)
	}

	router.Use(loggingMiddleware)
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+v\n", err)
		}
	}()

	log.Printf("Server Started")
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Printf("Server stopped")
	return nil
}

// loggingMiddleware logs each of the http requests
// received by the http server
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// listenForShutdown waits for a shutdown signal and cleans
// up the server gracefully if one is received
func listenForShutdown() (context.Context, context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	return ctx, cancel
}
