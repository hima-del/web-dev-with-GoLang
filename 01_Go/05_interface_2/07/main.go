package main

import "fmt"

type tank interface {
	Tarea() float64
	Volume() float64
}

func main() {

	var t tank
	fmt.Println("Value of the tank interface is: ", t)
	fmt.Printf("Type of the tank interface is: %T ", t)
}
