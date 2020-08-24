// Base API package that all services could follow as a pattern

// My notes for later export GOPATH="/Users/ejberozkan/projects/"
// src/github.com/ejber-ozkan/common-base-api
// dep (install and init-ed)
// bin src pkg folders required
// go mod init (migrated)

package main

import (
	"log"
	"net/http"

	"github.com/ejber-ozkan/common-base-api/routes"
)

func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
