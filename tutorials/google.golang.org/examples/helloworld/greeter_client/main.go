package main

import (
	"log"
	"net/http"

	"github.com/markTward/gospace/tutorials/google.golang.org/examples/helloworld/greeter_client/handlers"
)

const (
	address = ":8010"
)

func main() {
	http.HandleFunc("/hw", handlers.HelloWorld)
	http.HandleFunc("/healthcheck", handlers.HealthCheck)
	log.Fatal(http.ListenAndServe(address, nil))
}
