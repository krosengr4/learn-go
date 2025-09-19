package main

import "fmt"

// !  A buffered channel lets you write a certain number of times before there is a read on the channel

// ! Buf channel is never infinite. If you fill up the buffer, your rights will pause your GR until there is a read on the chan from another GR

func main() {
	// Create out channel with a buffer size of 10. The channel can hold up to 10 integer values before a send operation blocks.
	out := make(chan int, 10)

	// For loop that loops 10 times and creates Goroutine each time
	for i := 0; i < 10; i++ {
		go func(localI int) {
			// Writing i time 2 to the out channel for each Goroutine
			out <- localI * 2
		}(i)
	}

	// Create slice of type int to hold our results
	var result []int
	for i := 0; i < 10; i++ {
		// Reading each value from out channel and setting it to val
		val := <-out

		// Appending each val to the result slice
		result = append(result, val)
	}

	fmt.Println(result)
}
