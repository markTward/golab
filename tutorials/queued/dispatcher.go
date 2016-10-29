package main

import "fmt"

var WorkerQueue chan chan Job

func StartDispatcher(n int) {
	WorkerQueue = make(chan chan Job, WorkerQueueSize)

	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("Starting Worker", i)
			worker := NewWorker(i, WorkerQueue)
			worker.Start()
		}
	}()

	go func() {
		for {
			select {
			case job := <-JobQueue:
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching worker<-job:", job.ID)
					worker <- job
				}()
			}
		}
	}()
}
