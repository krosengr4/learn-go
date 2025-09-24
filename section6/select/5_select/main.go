package main

import (
	"fmt"
	"time"
)

// ! The done channel pattern is for graceful shutdown, synchronization, and avoiding Goroutine leaks

func multiples(i int) (chan int, chan struct{}) {

	out := make(chan int)

	// ? done is a signal channel to tell the Goroutine to stop
	// ? The primary purpose of having a chan struct{} is to communicate an event
	done := make(chan struct{})

	currentVal := 0
	go func() {
		for {
			select {
			case out <- currentVal * i:
				currentVal++
			case <-done:
				fmt.Println("Goroutine shutting down...")
				return
			}
		}
	}()

	return out, done
}

func main() {
	twosCh, done := multiples(2)

	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println("V is", v)
	}

	close(done)

	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

	fmt.Println("All done")
}
