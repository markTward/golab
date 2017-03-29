package main

import (
	"log"
	"net/http"

	"github.com/markTward/gospace/tutorials/google.golang.org/examples/helloworld/greeter_client/handlers"
)

const (
	address = "localhost:8010"
)

func main() {
	http.HandleFunc("/hw", handlers.HelloWorld)
	log.Fatal(http.ListenAndServe(address, nil))
}
