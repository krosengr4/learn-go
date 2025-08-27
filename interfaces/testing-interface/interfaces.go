package main

import "fmt"

//? An interface is a type that lists a set of methods but provides no implementation.

// Global variable divTest
type divTest int

// global struct with a min and max
type rangeTest struct {
	min int
	max int
}

func main() {
	result := runTests(10, []tester{
		rangeTest{min: 5, max: 20},
		divTest(5),
	})

	fmt.Println("All tests passed:", result)
}

// Declaring an interface.
type tester interface {
	// Method signatures
	test(int) bool
}

// Takes in an int, and a slice of types that implements tester(interface)
func runTests(i int, tests []tester) bool {
	result := true
	// Loop through each tester in the slice, call the test method on the current tester, and combine result with existing result
	for _, test := range tests {
		result = result && test.test(i) //<--- logical and operator(&&) to combine existing result with result of current test
	}
	// return the result
	return result
}

// method that takes in a var of type rangeTest(struct) and return a boolean
func (rt rangeTest) test(i int) bool {
	// Makes sure that i is between rangeTest min and max
	return rt.min <= i && rt.max >= i
}

// method that takes type divTest and returns a bool
func (dt divTest) test(i int) bool {
	// makes sure that i / divTest is a whole number
	return i%int(dt) == 0
}
