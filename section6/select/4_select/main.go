package main

import "fmt"

// ! In a long running program, this will create a lot of Goroutines and eventually swamp the Go runtime scheduler

// ! To fix this we need to make a done channel

// Function that takes in integer and returns a channel
func multiples(i int) chan int {
	// Unbuffered channel
	out := make(chan int)
	currentVal := 0

	// Goroutine with infinite for loop
	go func() {
		for {
			// Writing to out channel
			out <- currentVal * i
			currentVal++
		}
	}()

	return out
}

func main() {
	twosCh := multiples(2)

	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}
}
