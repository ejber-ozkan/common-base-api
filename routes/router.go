package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter building a new router from the routes array
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		//handler = utils.APILoggingHandler(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
