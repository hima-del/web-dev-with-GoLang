//type switch
package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("i am string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("i am int and my value is %d\n", i.(int))
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("himaja")
	findType(87)
	findType(78.89)
}
