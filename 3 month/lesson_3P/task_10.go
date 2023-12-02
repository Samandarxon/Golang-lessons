package main

import (
	"fmt"
	"time"
)

func Square(Channel chan int) {
	defer close(Channel)
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i * i
	}
	Channel <- sum

}
func main() {
	var (
		bch = make(chan int, 1)
	)
	x_time := time.Now()
	go Square(bch)

	fmt.Println("sum of squares of numbers from 1 to 10 concurrently: ", <-bch, "\nx_time: ", time.Since(x_time))

}
