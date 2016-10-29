package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func Hello(name string) string {
	// acquire/release worker via buffered tokens channel
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	rpc, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	return fmt.Sprintf("HelloWorld: %v\n", rpc.Message)

}

func main() {

	fmt.Println(Hello("mtw"))
	fmt.Println(Hello("mtward"))
	fmt.Println(Hello("marktward"))

}
