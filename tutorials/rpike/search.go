package main

import (
	"fmt"
	"math/rand"
	"time"
)

// define structure of Search func and Result
type Search func(query string) Result
type Result string

// create closure Search func by 'kind' and 'query' returning Result string
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// query all Search replicas and return first result provided
func FI(query string, replicas ...Search) Result {
	c := make(chan Result)
	// create searchReplica func executing fakeSearch('kind') closures on query string
	searchReplica := func(i int) { c <- replicas[i](query) }
	// starup all replicas
	for i := range replicas {
		go searchReplica(i)
	}
	// return first value found across all replicas; all others canceled
	return <-c
}

func Google(query string) (results []Result) {
	// channel for returning search results
	c := make(chan Result)

	// search func closures differentiated by 'kind'
	var (
		Web1, Web2     = fakeSearch("web"), fakeSearch("web")
		Image1, Image2 = fakeSearch("image"), fakeSearch("image")
		Video1, Video2 = fakeSearch("video"), fakeSearch("video")
	)

	// launch search server replicas
	go func() { c <- FI(query, Web1, Web2) }()
	go func() { c <- FI(query, Image1, Image2) }()
	go func() { c <- FI(query, Video1, Video2) }()

	// set search timeout all servers
	timeout := time.After(750 * time.Millisecond)

	// wait for 1st of each kind of search server to return a result totaling 3
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			// recv and accumulate FI replica search results
			results = append(results, result)
		case <-timeout:
			// halt search query after Nms, returning results to the ms
			fmt.Println("TIMEOUT!")
			return results
		}
	}
	// successful accumulation of search results across all kinds
	return results
}

func main() {

	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	// query google search sources on 'golang'
	results := Google("golang")

	elapsed := time.Since(start)
	fmt.Printf("Results: %v\n", results)
	fmt.Printf("Elapsed Time: %v\n", elapsed)
}
