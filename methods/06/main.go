//Pointer receivers in methods vs Pointer arguments in functions

package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func perimeter(r *rectangle) {
	fmt.Println("perimter function output is", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output is", 2*(r.length+r.width))
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}

	p := &r
	perimeter(p)
	p.perimeter()
	//perimeter(r)
	r.perimeter()
}
