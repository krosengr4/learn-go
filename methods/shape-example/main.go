package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) AreaRect() float64 {
	return r.height * r.width
}

func (r Rectangle) Parameter() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (c Circle) AreaCirc() float64 {
	return math.Pi * (math.Pow(c.radius, 2))
}

func (c Circle) Circumfrence() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	fmt.Println("1 - Rectangle\n2 - Circle")
	fmt.Println("Choose your shape:")
	var userShape int
	fmt.Scanln(&userShape)

	fmt.Println("1 - Area\n2 - Parameter / Circumfrence")
	fmt.Println("Choose what to calculate")
	var userCalc int
	fmt.Scanln(&userCalc)

	if userShape == 1 {
		fmt.Println("Enter the height:")
		var rectHeight float64
		fmt.Scanln(&rectHeight)

		fmt.Println("Enter the width:")
		var rectWidth float64
		fmt.Scanln(&rectWidth)

		rect := Rectangle{
			height: rectHeight,
			width:  rectWidth,
		}

		switch userCalc {
		case 1:
			result := rect.AreaRect()
			fmt.Println("Area of your rectangle:", result)
		case 2:
			result := rect.Parameter()
			fmt.Println("Parameter of your rectangle:", result)
		default:
			fmt.Println("ERROR: Entered wrong value for calculation!")
		}

	} else if userShape == 2 {
		fmt.Println("Enter the radius:")
		var userRadius float64
		fmt.Scanln(&userRadius)

		circ := Circle{radius: userRadius}

		switch userCalc {
		case 1:
			result := circ.AreaCirc()
			fmt.Println("Area of your circle:", result)
		case 2:
			result := circ.Circumfrence()
			fmt.Println("Circumfrence of your circle:", result)
		default:
			fmt.Println("ERROR: Entered wrong value for calculation!")

		}

	}

}
