package routes

import (
	"github.com/ejber-ozkan/common-base-api/handlers"
	"github.com/gorilla/mux"
)

// NewRouter a list of routes to follow on the service
func NewRouter() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	route.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
	return route
}
