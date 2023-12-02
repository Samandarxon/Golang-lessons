package main

import (
	"fmt"
	"time"
)

func Factorial(number chan int, fac int) {
	if fac == 0 {
		number <- 1
		close(number)
		return
	}
	for i := 1; i <= fac; i++ {
		number <- i
	}
	close(number)
}
func main() {

	var (
		number   = 5
		BChannel = make(chan int, number)
		fact     = 1
	)
	x_T := time.Now()
	go Factorial(BChannel, number)
	for v := range BChannel {
		fact *= v
	}

	fmt.Println("\tSum : ", fact, "\n\tTime: ", time.Since(x_T))
}
