package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	pbdb "github.com/markTward/grpc-demo/examples/db1/grpc/db"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addressDB = "localhost:50052"
)

var tokens = make(chan struct{}, 100)

// handler echoes the Path component of the requested URL.
func Base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func ServiceInfo(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressDB, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pbdb.NewDBServiceClient(conn)

	// acquire/release worker via buffered tokens channel
	tokens <- struct{}{}
	rpc, err := c.ServiceInfo(context.Background(), &pbdb.ServiceInfoRequest{})
	<-tokens

	fmt.Fprintf(w, "Service Descriptor: %T\t%v\n", rpc.Values, rpc.Values)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	is_alive := `{"is_alive": "true"}`
	io.WriteString(w, is_alive)
}

func DBRead(w http.ResponseWriter, r *http.Request) {
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

func DBUpsert(w http.ResponseWriter, r *http.Request) {
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
