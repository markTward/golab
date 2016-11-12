package main

import (
	"fmt"
	"reflect"
)

func typeofvar(v interface{}) string {

	switch v := v.(type) {
	case int:
		return "int"
	case bool:
		return "bool"
	default:
		_ = v
		return "unknown"
	}
}

func typeof(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	//... etc
	default:
		_ = t
		return "unknown"
	}
}

func main() {
	var b bool
	var i int
	var s string
	fmt.Println("bool", typeofvar(b))
	fmt.Println("string", typeofvar(s))
	fmt.Println("int", reflect.TypeOf(i))
}
