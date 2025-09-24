package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int, 1)

	chan2 <- 5

	// ? Select statements act like switch statements
	// ? Select statements are used when you want to skip over the channels that are blocked, and only send or recieve from channels that are ready
	select {
	case chan1 <- 2:
		fmt.Println("Wrote 2 to channel 1")
	case v := <-chan2:
		fmt.Println("read", v, "from channel 2")
	}
}
