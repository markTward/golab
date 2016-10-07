package main

import "fmt"

func main() {

	sarray := [2]string{"a", "b"}

	fmt.Printf("type / value: %T / %v\n", sarray, sarray)

	var sslice []string = sarray[:]
	sslice = append(sslice, "C")
	fmt.Printf("type / value: %T / %v\n", sslice, sslice)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	s = append(s, 999)
	fmt.Println(s)

}
