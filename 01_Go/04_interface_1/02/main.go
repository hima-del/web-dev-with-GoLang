package main

import "fmt"

type salaryCalculator interface {
	calculateSalary() int
}

type permenent struct {
	empid    int
	basicpay int
	pf       int
}

type contract struct {
	empid    int
	basicpay int
}

func (p permenent) calculateSalary() int {
	return p.basicpay + p.pf
}

func (c contract) calculateSalary() int {
	return c.basicpay
}

func totalExpense(s []salaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("total expense per month is $%d", expense)
}

func main() {
	e1 := permenent{
		empid:    1,
		basicpay: 5000,
		pf:       340,
	}
	e2 := permenent{
		empid:    2,
		basicpay: 6000,
		pf:       380,
	}
	e3 := permenent{
		empid:    3,
		basicpay: 5000,
		pf:       240,
	}
	e4 := contract{
		empid:    4,
		basicpay: 2000,
	}
	e5 := contract{
		empid:    5,
		basicpay: 3000,
	}
	e := []salaryCalculator{e1, e2, e3, e4, e5}
	totalExpense(e)
}
