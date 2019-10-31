package routes

import (
	"net/http"

	"github.com/ejber-ozkan/common-base-api/handlers"
)

//Route struct to build an array against
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes an array full of routes
type Routes []Route

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
