package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// collect input from stdin until "." signal from user or EOF
func getInput(input chan string, quit chan bool) {
	in := bufio.NewReader(os.Stdin)
	for {
		result, err := in.ReadString('\n')
		if err == io.EOF {
			quit <- true
		} else if err != nil {
			log.Fatal(err)
		}

		if result != ".\n" {
			input <- result
		} else {
			fmt.Println("exiting on user signal '.'")
			quit <- true
		}
	}
}

func main() {

	// control and processing channels
	input := make(chan string)
	timeout := make(chan bool)
	quit := make(chan bool)

	// collect word count
	wordcount := make(map[string]int)

	// use timer as timeout generator
	go func() {
		const timeOutInterval = 5 * time.Second
		timer := time.NewTimer(timeOutInterval)

		for {
			select {
			case <-timeout:
				// reset timer on recv on timeout channel
				timer.Reset(timeOutInterval)
			case <-timer.C:
				// quit if timeout interval reached with timer
				fmt.Printf("exiting after %v second inactivity\n", timeOutInterval)
				quit <- true
			}
		}
	}()

	// routine to gather input
	go getInput(input, quit)

	// simple/demo input processing engine: count the words and emit upon quit
	// while loop blocking on receiving a new line of text
	for {
		select {
		case line := <-input:
			for _, word := range strings.Split(strings.TrimSpace(line), " ") {
				wordcount[string(word)]++
			}
		case <-quit:
			jsmap, err := json.Marshal(wordcount)
			if err != nil {
				fmt.Printf("error: %v", err)
			} else {
				fmt.Println(string(jsmap))
			}
			return
		}
		// signal timeout channel to reset after a line of input
		timeout <- true
	}

}
