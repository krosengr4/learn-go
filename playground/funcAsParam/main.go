package main

import "fmt"

func main() {
	ifContinue := true

	for ifContinue {
		fmt.Println("Would you like to subtract or add?\n1 - Add\n2 - Subtract\n3 - Exit")
		var userChoice int
		fmt.Scanln(&userChoice)

		if userChoice == 3 {
			fmt.Println("Goodbye!")
			ifContinue = false
		} else {
			fmt.Println("Enter the first number:")
			var num1 int
			fmt.Scanln(&num1)

			fmt.Println("Enter the second number:")
			var num2 int
			fmt.Scanln(&num2)

			switch userChoice {
			case 1:
				printResult(num1, num2, add)
			case 2:
				printResult(num1, num2, subtract)

			}
		}
	}
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func printResult(a, b int, f func(int, int) int) {
	fmt.Println("Result:", f(a, b))
}
