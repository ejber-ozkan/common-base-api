package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ejber-ozkan/common-base-api/models"
)

// HelloHandler returns hello world!
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello World!")

}

// StatusHandler return a status of this common service endpoint in JSON
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	Status := models.Status{}

	Status.Level = "GREEN"
	Status.Description = "Everything is A OK"

	StatusBytes, err := json.Marshal(Status)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write(StatusBytes)

}
