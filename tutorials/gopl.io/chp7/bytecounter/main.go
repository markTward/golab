package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	lyrics := []string{"hello", "dolly", "hello", "dolly"}
	for _, word := range lyrics {
		c.Write([]byte(word))
	}
	fmt.Println("count:", c)

	c = 0
	lyrics = []string{"hey", "mister", "tambourine", "man", "..."}
	for _, word := range lyrics {
		c.Write([]byte(word))
	}
	fmt.Println("count:", c)
}
