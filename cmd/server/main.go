package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/asdine/storm"
	"github.com/brittonhayes/minikv/minikv"
	"github.com/gorilla/mux"
)

func main() {
	_ = minikv.Start(":8080", map[string]http.HandlerFunc{
		"/status/health": healthCheckHandler,
		"/{key}":         crudHandler,
	})
}

func crudHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db, err := storm.Open("store.db")
	if err != nil {
		log.Println(err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		var b []byte
		err = db.Get("kv", vars["key"], &b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if b != nil {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(b)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		if err := db.Set("kv", vars["key"], &body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		log.Println(http.StatusCreated, fmt.Sprintf("%q", string(body)))
	}
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = io.WriteString(w, `{"alive": true}`)
}
