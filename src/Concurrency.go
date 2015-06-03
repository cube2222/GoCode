package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Hello World")
			c <- 1
		}()
	}

	for i := 0; i < 5; i++ {
		<-c
	}

}
