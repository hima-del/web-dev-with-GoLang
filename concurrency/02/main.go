package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		fmt.Println(str)
	}
}

func main() {
	go display("welcome")
	go display("hello")
	time.Sleep(1 * time.Second)
}
