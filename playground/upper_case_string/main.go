package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your name:")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	fmt.Println("\nYour input:", name)

	lowerName := strings.ToLower(name)
	fmt.Println("Lower cased name:", lowerName)

	upperName := strings.ToUpper(name)
	fmt.Println("Upper cased name:", upperName)

}
