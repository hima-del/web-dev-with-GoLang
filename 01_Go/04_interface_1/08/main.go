package main

import "fmt"

type describer interface {
	describe()
}

type person struct {
	name string
	age  int
}

func (p person) describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case describer:
		v.describe()
	default:
		fmt.Printf("unknown type")
	}
}

func main() {
	findType("himaja")
	p := person{
		name: "himaja",
		age:  25,
	}
	findType(p)
}
