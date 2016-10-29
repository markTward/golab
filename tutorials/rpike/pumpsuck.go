package main

import (
	"fmt"
	"time"
)

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	stream := pump()
	fmt.Println(<-stream)
	suck(pump())

	time.Sleep(time.Second * 3)

}
