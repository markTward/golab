package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	// var wg sync.WaitGroup

	ct := make(chan bool)
	inputChan := make(chan string)
	quit := make(chan bool)

	go func() {
		const timeOutInterval = 5 * time.Second
		timer := time.NewTimer(timeOutInterval)

		for {
			select {
			case <-ct:
				fmt.Println("resetting timer")
				timer.Reset(timeOutInterval)
			case <-timer.C:
				fmt.Println("sending quit sig")
				quit <- true
			}
		}
	}()

	go func() {
		for {
			select {
			case line := <-inputChan:
				if in, err := strconv.Atoi(line); err == nil {
					fmt.Println("int test ==>", in, err)
				} else if ft, err := strconv.ParseFloat(line, 64); err == nil {
					fmt.Println("float64 test ==>", ft, err)
				} else {
					fmt.Println("str test ==>", in, err)
				}
			}
			ct <- true
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-quit:
			fmt.Println("preparing to qUIT")
			return
		case inputChan <- scanner.Text():
		}
	}

}
