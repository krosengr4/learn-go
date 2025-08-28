package main

import "fmt"

//! You can make a function implement an interface by defining a function type and
//! defining a method on the function type

type tester interface {
	test(int) bool
}

// func (runTests) that takes in an int and a slice of the tester interface
// type (funcs or methods that implement tester), and returns a bool
func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

// declare a type func (tester) that takes in an int and returns a bool
type testerFunc func(int) bool

// declare a method on testerFunc() (test) that takes in an int and returns a bool
func (tf testerFunc) test(i int) bool {
	// call testerFunc(), pass in the value this method took in, and this method will return the value from the function
	// * Because our test func type defines a method (test) that takes in an int and return a bool ---> it implements the tester interface
	return tf(i)
}

func main() {
	result := runTests(10, []tester{
		testerFunc(func(i int) bool {
			return i%2 == 0
		}),
		testerFunc(func(i int) bool {
			return i < 20
		}),
	})
	fmt.Println(result)
}
