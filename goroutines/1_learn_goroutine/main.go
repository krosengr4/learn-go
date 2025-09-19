package main

import "fmt"

// ! When running this, nothing happens
// ! This is bc the main goroutine did not know it had to wait for the other goroutine to complete

func main() {
	go runMe()
}

func runMe() {
	fmt.Println("Hello from inside a go routine")
}
