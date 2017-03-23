package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	tz := os.Getenv("TZ")
	loc, _ := time.LoadLocation(tz)
	fmt.Println("Time:", time.Now().In(loc))

	port := flag.Int("port", 8000, "clock port")
	flag.Parse()
	fmt.Println("Port:", *port)
}
