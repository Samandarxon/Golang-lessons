package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countKDifference([]int{1, 2, 2, 1}, 1))
	fmt.Println(countKDifference([]int{3, 2, 1, 5, 4}, 2))
}
func countKDifference(nums []int, k int) int {
	count := 0
	// fmt.Println("i", "   ", "j")
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			// fmt.Println(i, "   ", j)
			if i < j && math.Abs(float64(nums[i]-nums[j])) == float64(k) {
				count++
			}

		}
	}
	return count
}
