// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"log"
	"net/http"

	"github.com/markTward/grpc-demo/examples/db1/client/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Base)
	http.HandleFunc("/db/read", handlers.DBRead)
	http.HandleFunc("/db/upsert", handlers.DBUpsert)
	http.HandleFunc("/healthcheck", handlers.HealthCheck)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
