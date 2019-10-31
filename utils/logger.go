package utils

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

//APILoggingHandler - an attempt to log something to STDOUT
func APILoggingHandler(h http.Handler) http.Handler {

	// to log to file (Commented out for now .. maybe a 'dev' only switch)
	//logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	panic(err)
	//}
	return handlers.LoggingHandler(os.Stdout, h)
}
