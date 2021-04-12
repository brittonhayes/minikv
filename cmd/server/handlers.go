package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/asdine/storm/v3"
	"github.com/gorilla/mux"
)

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

// getKey fetches the requested key from the URI path
// from the database and returns it in the http response
func getKey(w http.ResponseWriter, r *http.Request) {
	// Handle path variables
	vars := mux.Vars(r)

	// Open database
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fetch the key from the database
	var bytes []byte
	err = db.Get("kv", vars["key"], &bytes)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
		return
	}

	if bytes != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

// postValue accepts the user provided request body and enters
// the value into the database
func postValue(w http.ResponseWriter, r *http.Request) {
	// Handle path variables
	vars := mux.Vars(r)

	// Open database
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Post the key to the database
	if err := db.Set("kv", vars["key"], &body); err != nil {
		http.Error(w, "failed to set key", http.StatusInternalServerError)
		log.Println("failed to set key", err)
		return
	}

	// Return the success message
	w.WriteHeader(http.StatusCreated)
	log.Println(http.StatusCreated, "Success")
}

// healthCheckHandler checks if the server is alive and okay
func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"alive": "true",
	})
}
