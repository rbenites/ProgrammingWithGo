package main

import "fmt"

/*
-create two custom struct types called triangle and square - DONE
-square should be a struct with a field called sideLength of type float64 - DONE
-triangle should be a struct with a field called height of type float64 and a field of type base of type float64 - DONE
-both types should have a function called getArea that returns the calculated area of the square or triangle

-Area of triangle = 0.5*base*height
-Are of square = sideLength*sideLength

-Add a shape interface tht defines a function called printArea.
-this function should calculate the area of the shape and print it to the terminal
*/

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{height: 2, base: 2}
	s := square{sideLength: 2}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("area of the shape is", s.getArea())
}
