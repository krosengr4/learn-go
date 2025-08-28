package main

import "fmt"

//? A method in go is a func that is assoiciated with a specific type.
//? Methods are defined with a reciever argument, which binds the method to and instance of that type.
///? This allows the method to operate on the data encapsulated within that instance.

// ! You cannot call a reference reciever if the struct has not been assigned to a variable
// ! If you call a value reciever on a pointer that is nil, the program will panic at runtime

type Foo struct {
	A int
	B string
}

func (f Foo) fieldCount() int {
	return 2
}

// Method declaration where method reciever (f) is of type Foo
func (f Foo) String() string {
	// Sprintf is used as a format string
	return fmt.Sprintf("A: %d and B: %s", f.A, f.B)
}

// Method reciever here is a reference reciever(f *Foo)
func (f *Foo) Double() {
	f.A = f.A * 2
}

type Bar struct {
	C bool
	Foo
}

func (b Bar) String() string {
	return fmt.Sprintf("%s and C: %v", b.Foo.String(), b.C)
}

func (b Bar) fieldCount() int {
	return 3
}

func main() {
	f := Foo{
		A: 10,
		B: "Hello",
	}

	b := Bar{
		C:   true,
		Foo: f,
	}

	fmt.Println(b.String())
	b.Double()
	fmt.Println(b.String())

}
