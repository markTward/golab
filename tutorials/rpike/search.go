package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func FI(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	var (
		Web1, Web2     = fakeSearch("web"), fakeSearch("web")
		Image1, Image2 = fakeSearch("image"), fakeSearch("image")
		Video1, Video2 = fakeSearch("video"), fakeSearch("video")
	)

	go func() { c <- FI(query, Web1, Web2) }()
	go func() { c <- FI(query, Image1, Image2) }()
	go func() { c <- FI(query, Video1, Video2) }()

	timeout := time.After(750 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("TIMEOUT!")
			return results
		}
	}
	return results
}

func main() {

	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	results := Google("golang")

	elapsed := time.Since(start)
	fmt.Printf("Results: %v\n", results)
	fmt.Printf("Elapsed Time: %v\n", elapsed)
}
