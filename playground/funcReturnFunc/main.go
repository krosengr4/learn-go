package main

import "fmt"

func main() {
	ifContinue := true

	for ifContinue {
		fmt.Println("Would you like to add or multiply?\n1 - Add\n2 - Multiply\n3 - Quit")
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
				sum := adder(num1)
				fmt.Println("Sum:", sum(num2))

			case 2:
				product := multiplier(num1)
				fmt.Println("Product:", product(num2))
			}
		}
	}
}

func adder(a int) func(int) int {
	return func(b int) int {
		return b + a
	}
}

func multiplier(a int) func(int) int {
	return func(b int) int {
		return b * a
	}
}
