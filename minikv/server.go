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

func Start(address string, routes map[string]http.HandlerFunc) error {

	ctx, cancel := listenForShutdown()
	router := mux.NewRouter()
	for path, handler := range routes {
		router.Handle(path, handler)
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

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
