package main

import "fmt"

func display(str string) {
	for w := 0; w < 6; w++ {
		fmt.Println(str)
	}
}

func main() {
	go display("welcome")
	display("hello")
}
