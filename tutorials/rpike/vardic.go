package main

import "fmt"

func Min(args ...int) int {
	min := int(^uint(0) >> 1)
	for _, i := range args {
		if min > i {
			min = i
		}
	}
	return min
}

func main() {

	fmt.Println(Min(11, 22, 333))
	fmt.Println(Min())
	slice1 := []int{4, 5, 1, -1}
	fmt.Println("slice1:", slice1)
	fmt.Println("slice1", Min(slice1...))

	slice2 := []int{10, 9, 8}
	slice1 = append(slice1, slice2...)
	fmt.Println("append s1", slice1)
}
