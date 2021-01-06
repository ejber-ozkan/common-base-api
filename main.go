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
	"github.com/ejber-ozkan/common-base-api/utils"
)

func main() {

	router := routes.NewRouter()

	loggedRouter := utils.APILoggingHandler(router)

	log.Fatal(http.ListenAndServe(":8080", loggedRouter))

}
