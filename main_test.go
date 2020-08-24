package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ejber-ozkan/common-base-api/handlers"
	"github.com/ejber-ozkan/common-base-api/models"
	"github.com/ejber-ozkan/common-base-api/routes"
)

func TestHandler(t *testing.T) {
	var err error
	request, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handlerFunction := http.HandlerFunc(handlers.HelloHandler)

	handlerFunction.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status: got %v want %v", status, http.StatusOK)
	}

	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}

}

func TestRouter(t *testing.T) {
	var err error
	r := routes.NewRouter()

	mockServer := httptest.NewServer(r)

	response, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status should ok,got %d", response.StatusCode)
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	responseString := string(bytes)
	expected := "Hello World!"

	if responseString != expected {
		t.Errorf("Response should be %s, got %s", expected, responseString)
	}
}

func TestRouterForNoExistentRoute(t *testing.T) {
	var err error
	router := routes.NewRouter()

	mockServer := httptest.NewServer(router)

	response, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// 405 (method not allowed)
	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", response.StatusCode)
	}

	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	responseString := string(b)
	expected := ""

	if responseString != expected {
		t.Errorf("Response should be %s, got %s", expected, responseString)
	}

}

func TestStatusRouter(t *testing.T) {
	var err error
	router := routes.NewRouter()

	mockServer := httptest.NewServer(router)

	response, err := http.Get(mockServer.URL + "/status")

	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status should ok,got %d", response.StatusCode)
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	status := models.Status{}
	err = json.Unmarshal(bytes, &status)

	if err != nil {
		t.Fatal(err)
	}

	if status.Level != "GREEN" {
		t.Errorf("Status should be GREEN,got %s", status.Level)
	}
}

func TestNotFoundHandler(t *testing.T) {
	var err error
	router := routes.NewRouter()

	mockServer := httptest.NewServer(router)

	response, err := http.Post(mockServer.URL+"/doesnotexist", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// 404 (method not found)
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Status should be 404, got %d", response.StatusCode)
	}

}
