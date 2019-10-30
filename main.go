// Base API package that all servcies could follow as a pattern

// My notes for later export GOPATH="/Users/ejberozkan/projects/"
// src/github.com/ejber-ozkan/common-base-api
// dep (install and init-ed)
// bin src pkg folders required
// go mod init (migrated)

package main

import (
	"net/http"

	"github.com/ejber-ozkan/common-base-api/routes"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":8080", router)
}
