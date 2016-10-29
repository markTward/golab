package main

import "fmt"

var WorkerQueue chan chan WorkRequest

func StartDispatcher(n int) {
	WorkerQueue := make(chan chan WorkRequest, n)

	for i := 0; i < n; i++ {
		fmt.Println("attempt to start worker", i)
		worker := NewWorker(i, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("WorkQueue:", WorkQueue, len(WorkQueue))
				go func() {
					fmt.Println("WorkerQueue:", WorkerQueue, len(WorkerQueue))
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request worker:", worker)
					worker <- work
					fmt.Println("Sending worker<-work")
				}()
			}
		}
	}()
}
