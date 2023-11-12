package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}
func singleNumber(nums []int) int {
	exactly_once := 0
	for _, el := range nums {
		j := 0
		for _, v := range nums {
			if el == v {
				j++
			}
		}
		if j < 2 {
			exactly_once = el
			break
		}
	}
	return exactly_once
}
