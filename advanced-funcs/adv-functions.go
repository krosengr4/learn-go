package main

import "fmt"

func main() {
	// You can assign a function to a var and treat that var just like the declared function
	varAddOne := addOne
	// Now we can add the parameters to the variable
	fmt.Println("---SETTING FUNCTION TO A VAR---")
	fmt.Println(varAddOne(1))
	fmt.Println(varAddOne(2))

	// Setting a variable to a function
	// To declare a function inside of another function, you need to assign it to a variable
	fmt.Println("---A FUNCTION WITHIN A FUNCITON---")
	varAddTwo := func(a int) int {
		return a + 2
	}
	fmt.Println(varAddTwo(1))

	// Here we are passing the function subtractOne() to the printResult() function
	fmt.Println("---PASSING A FUNCTION TO ANOTHER FUNCTION---")
	printResult(5, subtractOne)
	printResult(4, addOne)

	// ! A closure is a local function that has access to the variables that exist in the env where it was declared
	// ! An inner function that can capture variables from its surrounding scope, and retain access even after the function has finished
	fmt.Println("---USING LOCAL VARIABLE IN FUNCTION---")
	closureFunc()
	closure2()

	fmt.Println("---USING A FUNCTION THAT RETURNS A FUNCTION---")
	//? Here we must include parameters when we call the function makeAdder()
	num1, num2 := 7, 3
	sum := makeAdder(num1)
	// ? And also when we access the variable sum
	fmt.Println("Sum of", num1, "and", num2, ":", sum(num2))

	fmt.Println("---USING A FUNCTION THAT TAKES IN A FUNCTION AND RETURNS A FUNCTION---")
	myAddOne := makeAdder(1)
	doubleAddOne := makeDoubler(myAddOne)

	fmt.Println("Add one:", myAddOne(1))
	fmt.Println("My doubler:", doubleAddOne(1))

}

func addOne(a int) int {
	return a + 1
}

func subtractOne(a int) int {
	return a - 1
}

// Function that takes in another function as a parameter
func printResult(a int, function func(int) int) { //<--- Return type = int, so the functions being passed in also need a return type of int
	fmt.Println(function(a))
}

// We can also use a local variable inside of a function, assign the function to a variable, and then call it
// This is called a closure
// ! An inner function that can capture variables from its surrounding scope, and retain access even after the function has finished
func closureFunc() {
	b := 2
	myAddOne := func(a int) int {
		return a + b
	}
	fmt.Println(myAddOne(1))
}

func closure2() {
	b := 2
	myAddOne := func(a int) {
		b = a + b
	}
	myAddOne(1)
	fmt.Println(b)
}

// We can also have a function that can return a function
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

// Putting everything together here to make a function that takes in a function and returns a function
func makeDoubler(f func(a int) int) func(int) int {

	return func(b int) int {
		a := f(b)
		return a * 2
	}
}
