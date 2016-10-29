package main

import (
	"fmt"
	"time"
)

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	Quit        chan bool
}

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		Quit:        make(chan bool)}
	return worker
}

func (w *Worker) Start() {
	fmt.Printf("Worker%d Starting\n", w.ID)
	go func() {
		for {
			w.WorkerQueue <- w.Work
			select {
			case work := <-w.Work:
				// simulate worker working
				fmt.Printf("Worker%v Starting Job %v for %v\n", w.ID, work.ID, work.Delay)
				time.Sleep(work.Delay)
				fmt.Printf("Worker%v Finishing Job %v for %v\n", w.ID, work.ID, work.Delay)

			case <-w.Quit:
				fmt.Printf("Worker%v Stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}
