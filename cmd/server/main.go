package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/asdine/storm"
	"github.com/brittonhayes/minikv/minikv"
	"github.com/gorilla/mux"
)

func init() {

}

func main() {
	_ = minikv.Start(":8080", map[string]http.HandlerFunc{
		"/status/health": healthCheckHandler,
		"/{key}":         crudHandler,
	})
}

func openDB() (*storm.DB, error) {
	dbPath := ""
	if os.Getenv("ENVIRONMENT") == "docker" {
		dbPath = "/data/"
	}

	db, err := storm.Open(fmt.Sprintf("%s%s", dbPath, "store.db"))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func crudHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(db)

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
		log.Println(http.StatusCreated, "Success")
	}
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = io.WriteString(w, `{"alive": true}`)
}
