//to create an anonymous goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the main function")
	go func() {
		fmt.Println("welcome!")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("goodbye to main function")
}
