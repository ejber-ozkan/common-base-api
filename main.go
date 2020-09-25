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

	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {

	cfg, err := jaegercfg.FromEnv()

	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}

	//jLogger := jaegerlog.StdLogger
	//jMetricsFactory := metrics.NullFactory
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer()

	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)

	router := routes.NewRouter()

	loggedRouter := utils.APILoggingHandler(router)

	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
