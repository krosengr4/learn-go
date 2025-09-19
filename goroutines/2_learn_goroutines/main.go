package main

import (
	"fmt"
	"time"
)

func runMe() {
	fmt.Println("Hello from inside a go routine")
}

func main() {
	go runMe()
	// ! This is a hacky way of getting the non-main goroutine to complete
	// Pausing the main goroutine for 1 sec so the other goroutine can complete
	time.Sleep(1 * time.Second)
}
