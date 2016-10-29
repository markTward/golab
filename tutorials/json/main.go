package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Key   string `json:"label1"`
	Value string `json:"label2"`
}

func main() {
	jstr, _ := json.Marshal("abc")
	fmt.Println(jstr, string(jstr))

	resp1 := &Response1{
		Key:   "abc",
		Value: "123",
	}

	jstruct, _ := json.Marshal(resp1)
	fmt.Println(resp1, jstruct, string(jstruct))
}
