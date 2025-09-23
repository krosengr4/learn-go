package main

import "fmt"

func main() {

	// Make 2 channels for in and out
	in := make(chan int)
	out := make(chan int)

	// Infinite Goroutine loop
	go func() {
		for {
			// read an integer from the in channel
			i := <-in
			// For each integer recieved, it doubles the val and sends it to the out channel
			out <- i * 2
		}
	}()

	// ! Make sure that you read and write from one channel at a time or else the Goroutine will deadlock and panic

	// Writing a 1 to the in channel
	in <- 1
	// Reading the output and storing it in o1
	o1 := <-out

	// Writing a 2 to the in channel
	in <- 2
	// Reading the output and storing it in o2
	o2 := <-out

	fmt.Println(o1, o2)
}
