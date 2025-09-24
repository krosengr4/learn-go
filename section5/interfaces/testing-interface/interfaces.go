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

// global struct with 2 ints
type addTest struct {
	num1, num2 int
}

// ! Declaring an interface.
type tester interface {
	// Method signatures
	test(int) bool
}

func main() {
	result := runTests(10, []tester{
		rangeTest{min: 5, max: 20},
		divTest(5),
		addTest{num1: 10, num2: 11},
	})

	fmt.Println("Did all tests pass:", result)
}

// Takes in an int, and a slice of interface types that implements tester(interface)
func runTests(i int, tests []tester) bool {
	result := true
	// Loop through each tester in the slice, call the test method on the current tester, and combine result with existing result
	for _, test := range tests {
		result = result && test.test(i) //<--- logical and operator(&&) to combine existing result with result of current test
	}
	// return the result
	return result
}

// method that tests that the number passed in is between a min and a max
func (rt rangeTest) test(i int) bool {
	// Makes sure that i is between rangeTest min and max
	return rt.min <= i && rt.max >= i
}

// method that tests that, when dividing our numbers, we get a whole number
func (dt divTest) test(i int) bool {
	// makes sure that i / divTest is a whole number
	return i%int(dt) == 0
}

// method that tests that the sum of 2 numbers is greater than 20
func (at addTest) test(i int) bool {
	return at.num1+at.num2 > 20
}
