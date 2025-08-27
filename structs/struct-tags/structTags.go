package main

// ! Go allows you to put some metadata on a field in a struct

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"` //<--- Struct tags
	Age  int    `json:"age"`
}

func main() {
	bob := `{ "name": "Bob", "age": 30}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}
