package main

import (
	"fmt"
	"strings"
)

// ! In this file we will use an empty interface and demonstrate type switch as opposed to type assertion

func doStuff(i interface{}) {
	// Using a switch statement instead of type assertion
	switch i := i.(type) {
	case int:
		fmt.Println("i is and int! Double i =", i+i)
		fmt.Println("Value:", i)
	case string:
		fmt.Println("i is a string! It is", len(i), "characters long!")
		fmt.Println("Value:", i)
	case bool:
		fmt.Println("i is a boolean!\nValue:", i)
	default:
		fmt.Println("I am not sure what to do with this...")
		fmt.Println("Value:", i)
	}
}

func main() {
	fmt.Println(strings.Repeat("-", 35))

	fmt.Println("===INT===")
	doStuff(10)

	fmt.Println("===STRING===")
	doStuff("Hello world!")

	fmt.Println("===BOOLEAN===")
	doStuff(true)

	fmt.Println("===FLOAT===")
	doStuff(18.2098)

	fmt.Println(strings.Repeat("-", 35))
}
