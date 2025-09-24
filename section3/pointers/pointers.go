package main

import "fmt"

func main() {
	a := 10
	b := &a //<--- initializing b as a pointer to a (same memory address)
	c := a  //<--- c is it's own variable(w its own mem address), but with the value that a holds

	// Printing out b will just print out a memory address, *b will print out the value that b is pointing to
	fmt.Println(a, b, *b, c)

	// This will change the value of a AND b, but not c
	a = 20
	fmt.Println(a, b, *b, c)

	// This will change the value of a
	*b = 30
	fmt.Println(a, b, *b, c)

	// This will only change the value of c
	c = 40
	fmt.Println(a, b, *b, c)
}
