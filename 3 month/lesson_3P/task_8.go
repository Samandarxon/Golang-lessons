package main

import (
	"fmt"
	"time"
)

func Average(Channel chan int, numbers []int) {
	defer close(Channel)
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	sum /= len(numbers)
	Channel <- sum

}
func main() {
	var (
		bch     = make(chan int, 1)
		X_Slice = []int{23, 45, 2, 34, 57, 8, 45, 56, 99, 21, 35, 47}
	)
	fmt.Println("Slice: ", X_Slice)
	x_time := time.Now()
	go Average(bch, X_Slice)

	fmt.Println("Average: ", <-bch, "\nx_time: ", time.Since(x_time))

}
