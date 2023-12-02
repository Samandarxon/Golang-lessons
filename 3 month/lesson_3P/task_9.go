package main

import (
	"fmt"
	"sort"
	"time"
)

func Sort(Channel chan []int, numbers []int) {
	defer close(Channel)
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	Channel <- numbers

}
func main() {
	var (
		bch     = make(chan []int, 1)
		X_Slice = []int{23, 45, 2, 34, 57, 8, 45, 56, 99, 21, 35, 47}
	)
	fmt.Println("Slice: ", X_Slice)
	x_time := time.Now()
	go Sort(bch, X_Slice)

	fmt.Println("Sorted: ", <-bch, "\nx_time: ", time.Since(x_time))

}
