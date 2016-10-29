package main

import (
	"math/rand"
	"time"
)

func main() {
	StartDispatcher(1000)
	njobs := 10000
	for id := 0; id < njobs; id++ {
		delay := time.Duration(rand.Intn(10)) * time.Second
		Collector(id, delay)
	}
	time.Sleep(time.Duration(60) * time.Second)
}
