package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	JobQueueSize        = 1000
	WorkerQueueSize     = 1000
	WorkerPool          = 1000
	StatChanSize        = 1000
	TestMaxJobs         = 10000
	TestDurationSeconds = 15
)

type JobStat struct {
	Duration time.Duration
}

type JobStats struct {
	JobsFinished  int
	TotalDuration time.Duration
}

var StatChan = make(chan JobStat, StatChanSize)

func main() {
	fmt.Println(time.Now(), "StartDispatcher start")
	StartDispatcher(WorkerPool)
	fmt.Println(time.Now(), "StartDispatcher finish")

	for id := 0; id < TestMaxJobs; id++ {
		delay := time.Duration(rand.Intn(10)) * time.Second
		Collector(id, delay)
	}
	js := JobStats{}

	go func() {
		for {
			stat, ok := <-StatChan
			if !ok {
				break
			}
			js.JobsFinished++
			js.TotalDuration += stat.Duration
		}
	}()

	time.Sleep(time.Duration(TestDurationSeconds) * time.Second)
	close(JobQueue)
	close(WorkerQueue)
	close(StatChan)

	fmt.Println("JobStats", js.JobsFinished, js.TotalDuration.Seconds())

}
