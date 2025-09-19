package main

import "fmt"

func main() {
	userName := getName()

	// Make 2 channels (in and out) by using make function, keyword chan, and type of data on the channel
	in := make(chan string)
	out := make(chan string)

	// Launch Goroutine
	go func() {
		// recieve operator (<-var). Channel is on the right side, value on the left.
		// ? This means we are *READING* from the in channel
		name := <-in

		// Recieve operator where value we are *writing* is on the right, and the channel is on the left
		// ? This means that we are *WRITING* a value to the channel
		out <- fmt.Sprintf("Hello, " + name + "!")
	}()
	// Writing the userName that the user input to the in channel
	in <- userName

	// Close the channel means you are done using it
	close(in)

	// Read from the out channel and putting it into a value called message
	message := <-out
	// Print out the value of message that was read one line above
	fmt.Println(message)
}

func getName() string {
	fmt.Println("Please enter your name:")
	var userName string
	fmt.Scan(&userName)

	return userName
}
