package main

import "fmt"

func max(nums ...int) int {

	max_num := nums[0]

	for _, val := range nums {
		if max_num < val {
			max_num = val
		}
	}

	return max_num
}

func main() {

	var nums = []int{3, 4, 5, 6, 7, 8, -2}
	fmt.Println(max(nums...))
}
