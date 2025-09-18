package main

import (
	"fmt"
	"os"
	"strconv"
)

// ! Because Error is an interface, we can make our own error types with more information about the error
type MyError struct {
	A       int
	B       int
	Message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("values %d and %d produced error: %s", me.A, me.B, me.Message)
}

func divAndMod(a, b int) (int, int, error) {
	if b == 0 {
		// Using MyError struct to return div by 0 error
		return 0, 0, &MyError{A: a, B: b, Message: "Cannot divide by zero"}
	}
	return a / b, a % b, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(fmt.Println("Expected two input parameters"))
		os.Exit(3)
	}

	a, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, remainder, err := divAndMod(int(a), int(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("Result: %d\nRemainder: %d\n", result, remainder)

}
