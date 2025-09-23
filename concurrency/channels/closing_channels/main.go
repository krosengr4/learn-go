package main

import "fmt"

func main() {
	// Create in and out channel
	in := make(chan int, 10)
	out := make(chan int)

	for i := 0; i < 10; i++ {
		// Writing i into the in channel
		in <- i
	}

	// ? Closing a channel means that channel will not have anymore values written to it
	// ? If a buffered channel is closed, any values in the buffer are still available to be read
	close(in)

	go func() {
		for {
			// If the value of ok is true, the channel is still open. If false, the channel was closed
			i, ok := <-in

			if !ok {
				close(out)
				break
			}
			out <- i * 2
		}
	}()

	// ? You can have a for range loop with channels
	for v := range out {
		fmt.Println(v)
	}
}
