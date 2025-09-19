package main

import (
	"fmt"
	"sync"
	"time"
)

// ! The logic that you run within a goroutine (runMe()) should have no should have anything to do with concurrency
// ! Try to keep concurrency logic seperate from business logic

func runMe() {
	fmt.Println("Hello from inside a go routine")
	// i := 3
	for i := 3; i >= 0; i-- {
		if i == 0 {
			fmt.Print(i)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a wait group
	var weightGroup sync.WaitGroup

	// Adding 1 weight group to the goroutine
	weightGroup.Add(1)

	// Launch a closure as a Goroutine
	go func() {
		runMe()
		weightGroup.Done()
	}()

	// Pauses the main Goroutine until the count of running Goroutines tracked by the weight group is 0
	weightGroup.Wait()
	fmt.Println(" All done!")
}
