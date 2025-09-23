package main

import "fmt"

func main() {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	chan2 <- "Bar"

	// ? When more than one channel are ready to be used, the Select case will pick one at random to use (this prevents certain kinds of deadlocks that can happen)
	select {
	case chan1 <- "Foo":
		fmt.Println("Wrote: Foo to channel 1")
	case v := <-chan2:
		fmt.Println("Read:", v, "to channel 2")
	}
}
