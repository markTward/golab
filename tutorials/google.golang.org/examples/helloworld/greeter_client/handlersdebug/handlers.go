package handlersdebug

import (
	"fmt"
	"log"
	"net/http"
)

// xyz
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Printf("Greeting: %s", "DEBUG HelloWorld GRPC")

	// w.Write(w, "DEBUG HelloWorld GRPC")
	fmt.Fprint(w, "DEBUG HelloWorld GRPC")

}

func handler(write http.ResponseWriter, req *http.Request) {
	fmt.Fprint(write, "<h1>Hello!</h1>")
}
