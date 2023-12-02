package main

import (
	"fmt"
	"time"
)

func Multiply(Channel chan []int, numbers []int) {
	defer close(Channel)
	for i, v := range numbers {
		numbers[i] = v * 2
	}
	Channel <- numbers

}
func main() {
	var (
		bch     = make(chan []int, 1)
		X_Slice = []int{23, 45, 2, 34, 57, 8, 45, 56, 99, 21, 35, 47}
	)
	fmt.Println("Befor slice: ", X_Slice)
	x_time := time.Now()
	go Multiply(bch, X_Slice)

	fmt.Println("After: ", <-bch, "\nx_time: ", time.Since(x_time))

}
