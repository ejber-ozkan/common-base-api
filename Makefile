#include .env

PROJECTNAME=$(shell basename "$(PWD)")
GOBASE=$(shell pwd)
GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

JAEGER_SERVICE_NAME="api"
JAEGER_AGENT_HOST="jaeger"
JAEGER_AGENT_PORT="6831"
JAEGER_ENDPOINT="jaeger"
JAEGER_SAMPLER_MANAGER_HOST_PORT="jaeger:5778"


build:
	go build -o bin/main main.go

test:
	go test main.go main_test.go -v

run:
	go run main.go

all: test build 
