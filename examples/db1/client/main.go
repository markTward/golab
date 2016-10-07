// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"

	pb "github.com/markTward/grpc-demo/examples/db1/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/query", query)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func query(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Query %v\tString %v\n", r.URL.Path, r.URL.Query())
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// examine query string if 'name' key exists, send to server, otherwise default
	name := defaultName
	qsname, ok := r.URL.Query()["name"]
	// qsname, ok := qs["name"]

	if ok {
		name = qsname[0]
	}

	// Contact the server and print out its response.
	rpc, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Fprintf(w, "Greeting: %s\n", rpc.Message)

}

//!-
