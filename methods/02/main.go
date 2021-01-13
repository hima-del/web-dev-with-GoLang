package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length int
	width  int
}

type circle struct {
	radius float64
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("area of rectangle is %d\n", r.area())

	c := circle{
		radius: 12,
	}
	fmt.Printf("area of circle is %f\n", c.area())
}
