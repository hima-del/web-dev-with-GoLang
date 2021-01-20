package main

import "fmt"

type employee struct {
	name     string
	salary   int
	currency string
}

func (e employee) displaySalary() {
	fmt.Printf("salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	emp1 := employee{
		name:     "anna",
		salary:   40000,
		currency: "$",
	}
	emp1.displaySalary()
}
