package main

import (
	"fmt"
	"time"
)

type Worker struct {
	ID          int
	Job         chan Job
	WorkerQueue chan chan Job
	Quit        chan bool
}

func NewWorker(id int, workerQueue chan chan Job) Worker {
	worker := Worker{
		ID:          id,
		Job:         make(chan Job),
		WorkerQueue: workerQueue,
		Quit:        make(chan bool)}
	return worker
}

func (w *Worker) Start() {
	fmt.Printf("Worker%d Starting\n", w.ID)
	go func() {
		for {
			w.WorkerQueue <- w.Job
			select {
			case job := <-w.Job:
				// simulate worker working
				fmt.Printf("%v Worker%v Starting Job %v for %v\n", time.Now(), w.ID, job.ID, job.Delay)
				time.Sleep(job.Delay)
				fmt.Printf("%v Worker%v Finishing Job %v for %v\n", time.Now(), w.ID, job.ID, job.Delay)
				StatChan <- JobStat{job.Delay}

			case <-w.Quit:
				fmt.Printf("Worker%v Stopping\n", w.ID)
				return
			}
		}
	}()
}
