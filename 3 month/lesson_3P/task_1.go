package main

import (
	"fmt"
	"time"
)

func print(number chan int) {
	for i := 1; i <= 10; i++ {
		number <- i
	}
	close(number)
}

func main() {
	var bCh = make(chan int, 10)

	timeNow := time.Now()
	go print(bCh)
	for v := range bCh {
		fmt.Println(v)
	}
	fmt.Println("timeNow: ", time.Since(timeNow))
}
