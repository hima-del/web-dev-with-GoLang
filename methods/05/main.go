//Value receivers in methods vs Value arguments in functions

package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("area function result:%d\n", r.length*r.width)
}

func (r rectangle) area() {
	fmt.Printf("area method result:%d\n", r.length*r.width)
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	p.area()
}
