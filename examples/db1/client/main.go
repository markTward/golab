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

var tokens = make(chan struct{}, 100)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/helloagain", helloAgain)
	http.HandleFunc("/db/read", dbRead)
	http.HandleFunc("/db/upsert", dbUpsert)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func dbRead(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressDB, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbdb.NewDBServiceClient(conn)

	keys := r.URL.Query()["key"]

	// acquire/release worker via buffered tokens channel
	tokens <- struct{}{}
	rpc, err := c.Read(context.Background(), &pbdb.ReadRequest{Keys: keys})
	<-tokens

	log.Printf("dbRead: Keys: %v\t Values: %v\n", keys, rpc.Values)

	for idx, key := range keys {
		fmt.Fprintf(w, "Key: %v\t Value: %v\n", key, rpc.Values[idx])
	}

}

func dbUpsert(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressDB, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbdb.NewDBServiceClient(conn)

	fmt.Fprintf(w, "DEBUG URL Query(): %v\n", r.URL.Query())

	// TODO: move validation logic to server side, returning bad request / error
	// key/value pair validations:
	// 1 and only 1 key required
	// 0 or 1 value required
	var key, value string

	if len(r.URL.Query()["key"]) != 1 {
		fmt.Fprintf(w, "%v\t1 and only 1 value allowed: %v\n", http.StatusBadRequest, r.URL.Query()["key"])
		return
	} else {
		key = r.URL.Query()["key"][0]
	}

	if len(r.URL.Query()["value"]) > 1 {
		fmt.Fprintf(w, "%v\t0 or only 1 value allowed: %v\n", http.StatusBadRequest, r.URL.Query()["value"])
		return
	} else {
		if len(r.URL.Query()["value"]) == 1 {
			value = r.URL.Query()["value"][0]
		}
	}

	log.Printf("dbUpsert: Key: %v\t Value: %v\n", key, value)

	// acquire/release worker via buffered tokens channel
	tokens <- struct{}{}
	rpc, err := c.Upsert(context.Background(), &pbdb.UpsertRequest{Key: key, Value: value})
	<-tokens

	if err != nil {
		log.Printf("could not upsert record: %v", err)
	}
	fmt.Fprintf(w, "db[%v]=%v\n", key, rpc.Value)

}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// including hello/again as test/example for using grpc 2nd service and multiple protocol buffers
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
