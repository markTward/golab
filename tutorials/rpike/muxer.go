package main

import "fmt"

// type reply struct{}
type request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

func run(op binOp, req *request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service <-chan *request) {
	for {
		req := <-service
		go run(op, req)
	}
}

func startServer(op binOp) chan<- *request {
	service := make(chan *request)
	go server(op, service)
	return service
}

func (r *request) String() string {
	return fmt.Sprintf("%d+%d=%d",
		r.a, r.b, <-r.replyc)
}

func main() {

	adderChan := startServer(
		func(a, b int) int { return a + b },
	)

	subberChan := startServer(
		func(a, b int) int { return a - b },
	)

	req1 := &request{7, 8, make(chan int)}
	req2 := &request{17, 18, make(chan int)}
	adderChan <- req1
	adderChan <- req2
	fmt.Println(req2, req1)

	subberChan <- req1
	subberChan <- req2
	fmt.Printf("%d-%d=%d\n", req1.a, req1.b, <-req1.replyc)
	fmt.Printf("%d-%d=%d\n", req2.a, req2.b, <-req2.replyc)

}
