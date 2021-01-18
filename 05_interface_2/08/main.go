package main

import "fmt"

func myfun(a interface{}) {
	val := a.(string)
	fmt.Println("Value: ", val)
}
func main() {

	var val interface {
	} = "Himaja"

	myfun(val)
}
