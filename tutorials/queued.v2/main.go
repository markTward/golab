package main

import (
	"fmt"
	"math/rand"
	"time"
)

type JobStat struct {
	Duration time.Duration
}

type JobStats struct {
	JobsFinished  int
	TotalDuration time.Duration
}

var StatChan = make(chan JobStat, 100)

func main() {
	fmt.Println(time.Now(), "StartDispatcher start")
	StartDispatcher(100)
	fmt.Println(time.Now(), "StartDispatcher finish")

	njobs := 1000
	for id := 0; id < njobs; id++ {
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

	time.Sleep(time.Duration(15) * time.Second)
	close(JobQueue)
	close(WorkerQueue)
	close(StatChan)

	fmt.Println("JobStats", js.JobsFinished, js.TotalDuration.Seconds())

}
