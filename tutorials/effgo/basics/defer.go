package main

import "fmt"

func trace(f string) string {
	fmt.Println("entering func", f)
	return f
}

func un(f string) {
	fmt.Println("leaving func", f)
}

func a(a string) {
	defer un(trace("a"))
	fmt.Println("workon", a)
	b("b")
	return
}

func b(b string) {
	defer un(trace("b"))
	fmt.Println("working on", b)
	return
}

func main() {
	a("a")
}
