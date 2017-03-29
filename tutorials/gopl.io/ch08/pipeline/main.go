// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 228.

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	naturals := make(chan int)
	squares := make(chan int)

	go counter(n, naturals)
	go squarer(naturals, squares)
	printer(squares)
}

// Counter
func counter(n int, out chan<- int) {
	defer close(out)
	for x := 0; x < n; x++ {
		out <- x
	}
}

// Squarer
func squarer(in <-chan int, out chan<- int) {
	defer close(out)
	for x := range in {
		out <- x * x
	}
}

// Printer (in main goroutine)
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
