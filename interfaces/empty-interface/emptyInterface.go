package main

import "fmt"

func main() {
	// i is declared as an empty interface, which can store any type(string, int, struct, etc)
	var i interface{}

	// We are storing a string in the interface
	i = "Hello"

	//! You can only do a type assertion on an interface
	// Type assertion. Takes the value in i and treats it as a string (if i was not a string, the program would panic at runtime)
	j := i.(string)
	fmt.Println(j)

	// Safe assertion bc of the , ok. i does not hold an int, but the program won't panic.
	// k will default to a 0 value(0 for int) and ok will be false
	k, ok := i.(int)
	fmt.Println(k, ok)

	// This will cause panic at runtime because i does not hold an int and we don't use a safe assertion
	// m := i.(int)
	// fmt.Println(m)

}
