# common-base-api

This is an 'Hello World' API.

Mainly built for me to learn Golang.

Using various go modules:

## Gorilla handlers

Package handlers is a collection of handlers (aka "HTTP middleware") for use with Go's net/http package (or any framework supporting http.Handler) :

github.com/gorilla/handlers

## Gorilla mux

Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler :

github.com/gorilla/mux

## Build

To run tests

```bash

go test main_test.go

```

To build on local machine

```bash

go build main.go
# run ./main or #main.exe

```

This should also work with docker, you may need to play with [Dockerfile](Dockerfile)

Build it via docker

```bash
docker build -t common-base-api .
```

Then run it via docker

```bash
docker run -p 8080:8080 common-base-api:latest
```

Then point your browser to see some output:

<http://localhost:8080/hello>

<http://localhost:8080/status>

/hello outputs in plain text

/status returns a JSON object

the example also uses LoggingHandler from the Gorilla modules to output to stdout any response.

Has tests as well, as an example of bringing it together.
