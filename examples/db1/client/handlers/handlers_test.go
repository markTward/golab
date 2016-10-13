package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSanity(t *testing.T) {

	if true {
		log.Println("sanity! nothing to see here ...")
	} else {
		t.Error("big trouble!")
	}

}

func TestHealthCheck(t *testing.T) {
	// thanks to https://elithrar.github.io/article/testing-http-handlers-go/ for tutorial

	// create request for handler
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	// ResponseRecorder satisfies ResponseWriter
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)

	// handler satisfies http.Handler, so ServeHTTP method can be called and
	// Request and ResponseRecorder passed directly
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect response code.  expected %v received %v", status, http.StatusOK)
	}

	// test response against expected value
	expected := `{"is_alive": "true"}`
	if rr.Body.String() != expected {
		t.Errorf("handler return unexpected body: expected %v received %v", expected, rr.Body.String())
	}
}
