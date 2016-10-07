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

	pbdb "github.com/markTward/grpc-demo/examples/db1/grpc/db"
	pbhw "github.com/markTward/grpc-demo/examples/db1/grpc/hw"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addressHW   = "localhost:50051"
	addressDB   = "localhost:50052"
	defaultName = "world"
	defaultKey  = ""
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/query", query)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/helloagain", helloAgain)
	http.HandleFunc("/db/read", readRecord)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func readRecord(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressDB, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbdb.NewRecordReaderClient(conn)

	// examine query string if one/many 'name' keys exists
	// if empty, provide default
	qskeys, ok := r.URL.Query()["key"]
	if !ok {
		qskeys = append(qskeys, defaultKey)
	}

	// Contact gRPC helloworld server over range of names
	for _, key := range qskeys {
		rpc, err := c.Read(context.Background(), &pbdb.RecordKey{Key: key})

		if err != nil {
			log.Fatalf("could not find record: %v", err)
		}
		fmt.Fprintf(w, "db[%v]=%v\n", key, rpc.Value)
	}

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
	conn, err := grpc.Dial(addressHW, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbhw.NewGreeterClient(conn)

	// examine query string if one/many 'name' keys exists
	// if empty, provide default
	qsnames, ok := r.URL.Query()["name"]
	if !ok {
		qsnames = append(qsnames, defaultName)
	}

	// Contact gRPC helloworld server over range of names
	for _, name := range qsnames {
		rpc, err := c.SayHello(context.Background(), &pbhw.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Fprintf(w, "Greeting: %s\n", rpc.Message)
	}

}

func helloAgain(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressHW, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbhw.NewGreeterClient(conn)

	// examine query string if one/many 'name' keys exists
	// if empty, provide default
	qsnames, ok := r.URL.Query()["name"]
	if !ok {
		qsnames = append(qsnames, defaultName)
	}

	// Contact gRPC helloworld server over range of names
	for _, name := range qsnames {
		rpc, err := c.SayHelloAgain(context.Background(), &pbhw.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		fmt.Fprintf(w, "Greeting Again from new GRPC: %s\n", rpc.Message)
	}

}

//!-
