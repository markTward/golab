package main

import (
	"log"
	"time"
)

var WorkQueue = make(chan WorkRequest, 100)

func Collector(id int, delay time.Duration) {
	work := WorkRequest{ID: id, Delay: delay}
	log.Println("Collector received WorkRequest:", work)
	WorkQueue <- work
	return
}
