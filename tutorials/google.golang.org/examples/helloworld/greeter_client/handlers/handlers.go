package handlers

import (
	"log"
	"net/http"
	"strings"

	pb "github.com/markTward/gospace/tutorials/google.golang.org/examples/helloworld/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addressDB   = "localhost:8000"
	defaultName = "World!"
)

var tokens = make(chan struct{}, 100)

// Respond to HelloWorld request
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

	w.Write([]byte(rpc.Message))
}
