//pointer recivers vs value receivers

package main

import "fmt"

type employee struct {
	name string
	age  int
}

func (e employee) chnageName(newname string) {
	e.name = newname
}

func (e *employee) changeAge(newage int) {
	e.age = newage
}

func main() {
	e := employee{
		name: "anna",
		age:  27,
	}
	e.chnageName("ani")
	fmt.Println("name", e.name)
	(&e).changeAge(30)
	fmt.Println("age", e.age)
}
