package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Checks to see how many arguments on the CLI
	// Command looks like: go run main.go 10 2
	if len(os.Args) < 3 {
		fmt.Println("Expected two input parameters")
		os.Exit(1)
	}

	// Convert first CLI arg to int
	a, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		// Exits the program with exit code(1)
		os.Exit(1)
	}

	// Convert second CLI arg to int
	b, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		// Exits the program with exit code(1)
		os.Exit(1)
	}

	div, mod, err := divAndMod(int(a), int(b))
	if err != nil {
		fmt.Println(err)
		// Exits the program with exit code(1)
		os.Exit(2)
	}
	fmt.Printf("%d / %d == %d and %d %% %d == %d\n", a, b, div, a, b, mod)
}

func divAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return a / b, a % b, nil
}
