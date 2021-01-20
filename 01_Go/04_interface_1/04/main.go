//empty interface

package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("type = %T, value = %v\n", i, i)
}

func main() {
	s := "hello world"
	describe(s)
	i := 30
	describe(i)
	strct := struct {
		name string
	}{
		name: "Himaja",
	}
	describe(strct)
}
