package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter building a new router from the routes array
// adding static html file folder route
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	staticDir := "/assets/"

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	return router
}
