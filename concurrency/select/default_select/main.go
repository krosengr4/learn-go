package main

import "fmt"

func main() {
	// Create unbuffered in and out channels
	in := make(chan int)
	out := make(chan int)

	// ? We can use default whenever none of the cases can read or write to thier channels
	select {
	case in <- 2:
		fmt.Println("Wrote 2 to in")
	case v := <-out:
		fmt.Println("Wrote", v, "to out")
	default:
		fmt.Println("Nothing else worked")
	}
}
