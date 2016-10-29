package main

import (
	"fmt"
	"sync"
	"time"
)

func makeTakes(wg *sync.WaitGroup, name string, secs time.Duration) {
	defer wg.Done()
	time.Sleep(secs * time.Second)
	fmt.Printf("%v is ready\n", name)
	return
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	// waitgroup arg is POINTER!  passed as value failed (of course)
	go makeTakes(&wg, "dinner", 10)

	wg.Add(1)
	go makeTakes(&wg, "lunch", 5)

	wg.Add(1)
	go makeTakes(&wg, "breakfast", 3)

	fmt.Println("waiting ...")
	wg.Wait()

}
