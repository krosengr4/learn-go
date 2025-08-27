package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-----BASIC FUNCTION-----")
	basicFunction()
	fmt.Println(strings.Repeat("-", 40))

	fmt.Println("-----FUNCTION WITH PARAMETERS-----")
	funcWithParams(7, 3)
	fmt.Println(strings.Repeat("-", 40))

	fmt.Println("-----FUNCTION RETURNING ONE VALUE-----")
	result := returnValue(11, 9)
	fmt.Println("Result:", result)
	fmt.Println(strings.Repeat("-", 40))

	fmt.Println("-----FUNCTION RETURNING MANY VALUES-----")
	addResult, _ := returnMultiple(10, 7)
	// To only get the value of only one return, you can replace the other return with a _
	_, subtractResult := returnMultiple(10, 7)
	fmt.Println("Addition Result:", addResult)
	fmt.Println("Subtract Result:", subtractResult)
	fmt.Println(strings.Repeat("-", 40))

}

func basicFunction() {
	fmt.Println(2 + 3)
}

// Go doesn't have optional parameters or named parameters
func funcWithParams(a, b int) {
	fmt.Println(a + b)
}

// Function that return a single value
func returnValue(a, b int) int {
	return a + b
}

// Go also allows you to return multiple values from a funciton
func returnMultiple(a, b int) (int, int) {
	addition := a + b
	subtraction := a - b
	return addition, subtraction
}
