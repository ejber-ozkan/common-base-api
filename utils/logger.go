package utils

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

//APILoggingHandler - an attempt to log something
func APILoggingHandler(h http.Handler) http.Handler {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, h)
}
