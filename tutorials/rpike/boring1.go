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
		for {
			c <- msg
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	msgChan := make(<-chan string)
	quitChan := make(chan bool)

	// how long the party lasts
	pt := time.Duration(5 * time.Second)
	go partyIsOver(pt, quitChan)

	// join the party
	msgChan = boring("boring!")

	for {
		select {
		case m := <-msgChan:
			fmt.Println(m)
		case <-quitChan:
			close(quitChan)
			return
		}
	}
}
