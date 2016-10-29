package main

import (
	"fmt"
	"time"
)

var JobQueue = make(chan Job, 100)

func Collector(id int, delay time.Duration) {

	job := Job{ID: id, Delay: delay}
	fmt.Println(time.Now(), "Collector JobQueue<-job:", job)
	JobQueue <- job
	return

}
