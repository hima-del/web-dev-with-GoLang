// Go program to illustrate type assertion
package main

import "fmt"

func myfun(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}
func main() {

	var a1 interface {
	} = 98.09

	myfun(a1)

	var a2 interface {
	} = "GeeksforGeeks"

	myfun(a2)
}
