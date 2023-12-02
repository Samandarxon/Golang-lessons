package main

import (
	"fmt"
	"time"
)

func SumOfEvenNumbers(Channel chan int, x_slice []int) {
	defer close(Channel)
	sum := 0
	for _, v := range x_slice {
		if v%2 == 0 {
			sum += v
		}
	}
	Channel <- sum

}
func main() {
	var (
		bch     = make(chan int, 1)
		X_Slice = []int{23, 45, 2, 34, 57, 8, 45, 56, 99, 21, 35, 47}
	)
	fmt.Println("Slice: ", X_Slice)
	x_time := time.Now()
	go SumOfEvenNumbers(bch, X_Slice)

	fmt.Println("Sum of even numbers: ", <-bch, "\nx_time: ", time.Since(x_time))

}
