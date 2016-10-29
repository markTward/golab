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

func boring(msg string, message chan string) {
	for {
		message <- msg
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	msgChan := make(chan string)
	quitChan := make(chan bool)

	// how long the party lasts
	pt := time.Duration(5 * time.Second)
	go partyIsOver(pt, quitChan)

	// join the party
	go boring("boring!", msgChan)

	for {
		select {
		case m := <-msgChan:
			fmt.Println(m)
		case <-quitChan:
			close(msgChan)
			close(quitChan)
			return
		}
	}

}
