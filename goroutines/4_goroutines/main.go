package main

import (
	"fmt"
	"sync"
)

// ? Pass data into a goroutine by pasing the values into the function you are launching
func printName(name string) {
	fmt.Println("'Hello", name+"!' sincerely, Goroutine :)")
}

func printAge(age int) {
	fmt.Printf("You are %d years old!", age)
}

func main() {
	// Create weight group
	var wg sync.WaitGroup

	// ? Even though our goroutine calls 2 functions, they both execute within the same goroutine, so only 1 wg.Done() is needed
	// ? The number added to the wg corresponds to the number of wg.Done() instances there are
	wg.Add(1)

	fmt.Println("Enter your name:")
	var userName string
	fmt.Scanln(&userName)

	fmt.Println("Enter your age:")
	var userAge int
	fmt.Scanln(&userAge)

	// ? To use data in a goroutine, we also need to change our closure to take in a value
	go func(name string, age int) {
		printName(name)
		printAge(age)
		wg.Done()
	}(userName, userAge) //<--- Pass in the var holding the value of our data here

	wg.Wait()

}
