package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	pb "github.com/markTward/gospace/tutorials/google.golang.org/examples/helloworld/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addressDB   = ":8000"
	defaultName = "World!"
)

var tokens = make(chan struct{}, 100)

// HealthCheck simple
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("healthcheck OK")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	is_alive := `{"is_alive": "true"}`
	io.WriteString(w, is_alive)
}

// HelloWorld grpc request
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressDB, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	// Greeter Client
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	q := r.URL.Query()

	qname, ok := q["name"]
	var name string
	if ok {
		name = strings.Join(qname, ", ")
		// for _, v := range qname {
		// 	name = name + ", " + v
		// }
	} else {
		name = defaultName
	}

	rpc, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", rpc.Message)
	fmt.Fprint(w, rpc.Message)
}
