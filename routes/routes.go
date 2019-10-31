package routes

import (
	"net/http"

	"github.com/ejber-ozkan/common-base-api/handlers"
	"github.com/ejber-ozkan/common-base-api/utils"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// NewRouter a list of routes to follow on the service
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = utils.APILoggingHandler(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Hello",
		"GET",
		"/hello",
		handlers.HelloHandler,
	},
	Route{
		"Status",
		"GET",
		"/status",
		handlers.StatusHandler,
	},
}
