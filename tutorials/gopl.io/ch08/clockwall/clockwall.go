// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Println(os.Args)
	for _, arg := range os.Args[1:] {
		s := strings.Split(arg, "=")
		tz, port := s[0], s[1]

		fmt.Printf("%s\t", tz)

		go connServer(port)
	}

	select {}

}

func connServer(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	fmt.Println(port, conn)
	if err != nil {
		// log.Fatal(err)
		log.Println("error:", err)
		return
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
