#!/bin/bash

export JAEGER_SERVICE_NAME="api"
export JAEGER_AGENT_HOST="jaeger"
export JAEGER_AGENT_PORT="6831"
export JAEGER_ENDPOINT="http://jaeger:14268/api/traces"
export JAEGER_SAMPLER_MANAGER_HOST_PORT="jaeger:5778"

go run main.go
