package main

import "fmt"

type tank interface {
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

func (m myvalue) Tarea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank
	t = myvalue{10, 14}
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())
}
