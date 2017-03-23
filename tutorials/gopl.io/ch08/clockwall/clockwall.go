// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port := flag.Int("port", 8000, "clock port")
	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//!-
