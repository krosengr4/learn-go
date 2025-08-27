package main

import "fmt"

type Foo struct {
	A int
	B string
}

func main() {
	f := Foo{
		A: 10,
		B: "Hello",
	}
	fmt.Println(f.String())
	f.Double()
	fmt.Println(f.String())

}

// Method declaration where method reciever = (f Foo) (between keyword func and the name of the method)
func (f Foo) String() string {
	// Sprintf is used as a format string
	return fmt.Sprintf("A: %d and B: %s", f.A, f.B)
}

// Method reciever here is a reference reciever(f *Foo)
func (f *Foo) Double() {
	f.A = f.A * 2
}
