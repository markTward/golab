package main

import "fmt"
import "flag"

func main() {
	n := flag.String("n", "you", "your name")
	flag.Parse()
	fmt.Println("hello world from gospace", *n)
}
