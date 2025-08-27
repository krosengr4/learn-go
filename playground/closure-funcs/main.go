package main

import "fmt"

// ! A closure is a function that can capture variables from its surrounding scope, and retain access even after the function has finished
func main() {
	fmt.Println("Please enter your name:")
	var userName string
	fmt.Scanln(&userName)

	greeting := func(name string) string {
		myGreeting := "Hello " + name + "!"
		return myGreeting
	}

	fmt.Println(greeting(userName))
}
