package main

import "time"

type WorkRequest struct {
	ID    int
	Delay time.Duration
}
