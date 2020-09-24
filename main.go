// Base API package that all services could follow as a pattern

// My notes for later export GOPATH="/Users/ejberozkan/projects/"
// src/github.com/ejber-ozkan/common-base-api
// dep (install and init-ed)
// bin src pkg folders required
// go mod init (migrated)

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ejber-ozkan/common-base-api/routes"
	"github.com/ejber-ozkan/common-base-api/utils"

	opentracing "github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func main() {

	cfg, err := jaegercfg.FromEnv()

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
	}

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	router := routes.NewRouter()

	loggedRouter := utils.APILoggingHandler(router)

	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
