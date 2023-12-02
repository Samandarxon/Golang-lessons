package main

import (
	"fmt"
	"time"
)

func Sum(number chan int) {
	for i := 1; i < 101; i++ {
		number <- i
	}
	close(number)
}
func main() {

	var (
		BChannel = make(chan int, 100)
		sum      = 0
	)
	x_T := time.Now()
	go Sum(BChannel)
	for v := range BChannel {
		sum += v
	}

	fmt.Println("\tSum : ", sum, "\n\tTime: ", time.Since(x_T))
}
