package main

import "fmt"

// ! Embedding may look like inheritance but it isn't

type Foo struct {
	A int
	b string
}

type Bar struct {
	C string
	// F is of type Foo struct
	F Foo
}

type Baz struct {
	D string
	// This field has no var name, just the type struct
	// Used for embedding
	Foo
}

func main() {
	// Initialize a Foo struct
	f := Foo{A: 10, b: "Hello"}

	// Initialize Bar struct that uses f for F field
	b1 := Bar{C: "Fred", F: f}
	fmt.Println(b1.F.A) //<--- 10, how to reach the field inside of b1

	b2 := Baz{D: "Nancy", Foo: f}
	// Embedding makes it easier to reach this field in b2
	fmt.Println(b2.A)

	var f2 Foo = b2.Foo
	fmt.Println(f2)
}
