package main

import "fmt"
import "os"

func main() {

	fmt.Println("cmd", os.Args[0])
	fmt.Println(os.Args[1:])

}
