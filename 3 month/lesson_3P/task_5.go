package main

import (
	"fmt"
	"time"
)

func Max(Channel chan int, numbers []int) {
	defer close(Channel)
	max := numbers[0]
	for _, v := range numbers {
		if max < v {
			max = v
		}
	}
	Channel <- max

}
func main() {
	var (
		bch     = make(chan int, 1)
		X_Slice = []int{23, 45, 2, 34, 57, 8, 45, 56, 99, 21, 35, 47}
	)
	x_time := time.Now()
	go Max(bch, X_Slice)

	fmt.Println("Max: ", <-bch, "\nx_time: ", time.Since(x_time))

}
