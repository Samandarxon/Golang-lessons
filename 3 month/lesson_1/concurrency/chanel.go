package main

import (
	"fmt"
)

func increment(c chan int, count int) {
	count++
	c <- count
}

func main() {

	var (
		// unbuffered chanel
		c     = make(chan int)
		count = 10
	)

	go increment(c, count)
	count = <-c
	fmt.Println(count)

	var (
		// buffered chanel
		k = make(chan int, 10)
	)

	k <- 20
	go increment(k, 10)

	fmt.Println(<-k)
	fmt.Println(<-k)
	// fmt.Println(<-k)  - error deadlock
}
