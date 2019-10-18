// Base API package that all servcies could follow as a pattern

// My notes for later export GOPATH="/Users/ejberozkan/projects/"
// src/github.com/ejber-ozkan/common-base-api
// dep (install and init-ed)
// bin src pkg folders required
// go mod init (migrated)

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ejber-ozkan/common-base-api/models"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/hello", handler).Methods("GET")
	route.HandleFunc("/status", StatusHandler).Methods("GET")
	return route
}

func main() {
	router := newRouter()
	http.ListenAndServe(":8080", router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// StatusHandler return a status of this common service endpoint in JSON
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	Status := models.Status{}

	Status.Level = "GREEN"
	Status.Description = "Everything is A OK"

	StatusBytes, err := json.Marshal(Status)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(StatusBytes)
}
