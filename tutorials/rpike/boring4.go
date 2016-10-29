package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partyIsOver(partyTime time.Duration, quit chan bool) {
	partyTimer := time.NewTimer(partyTime)
	for {
		select {
		case <-partyTimer.C:
			quit <- true
			fmt.Println("partyIsOver")
			return
		default:
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s (%d)", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1 <-chan string, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	quitChan := make(chan bool)

	// how long the party lasts
	timeout := time.Duration(5 * time.Second)
	go partyIsOver(timeout, quitChan)

	// join the party
	john := boring("i'm john and i'm boring!")
	jane := boring("i'm jane and i'm boring!")

	msg := fanIn(john, jane)
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-quitChan:
			return
		}
	}
}
