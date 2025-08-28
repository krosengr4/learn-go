package main

import "fmt"

type Rectangle struct {
	width, height int
}

// method that is associated with Rectangle struct.
// returns the area of the rectangle passed in
func (r Rectangle) CalculateArea() int {
	return r.width * r.height
}

func main() {
	rectangle1 := Rectangle{
		width:  12,
		height: 3,
	}

	rectangle2 := Rectangle{
		width:  4,
		height: 7,
	}

	fmt.Println("The area of rectangle 1 is:", rectangle1.CalculateArea())
	fmt.Println("The area of rectangle 2 is:", rectangle2.CalculateArea())
}
