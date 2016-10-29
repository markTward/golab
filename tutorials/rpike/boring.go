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
			fmt.Println("partyIsOver")
			quit <- true
			return
			// default:
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

func gotoParty(input1 <-chan string, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
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

	// invited to the party
	john := boring("i'm john and i'm boring!")
	jane := boring("i'm jane and i'm boring!")

	chitChat := gotoParty(john, jane)

	for {
		select {
		case chit := <-chitChat:
			fmt.Println(chit)
		case <-quitChan:
			return
		}
	}
}
