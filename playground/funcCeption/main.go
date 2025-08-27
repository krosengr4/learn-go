package main

import "fmt"

func main() {
	fmt.Println("ENTER TWO NUMBERS THAT WILL BE ADDED TOGETHER AND THEN DOUBLED")

	fmt.Println("Enter the first number:")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number:")
	var num2 int
	fmt.Scanln(&num2)

	sum := adder(num1)
	fmt.Println("The sum of the two numbers:", sum(num2))

	doubledSum := doubler(sum)
	fmt.Println("The sum doubled:", doubledSum(2))
}

func adder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func doubler(f func(a int) int) func(int) int {
	return func(b int) int {
		a := f(b)
		return a * 2
	}
}
