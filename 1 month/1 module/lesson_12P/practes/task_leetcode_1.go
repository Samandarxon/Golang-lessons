// https://leetcode.com/problems/concatenation-of-array/

package main

import "fmt"

func getConcatenation(nums []int) []int {
	// x_slice:=[]int{}
	nums = append(nums, nums...)
	fmt.Println(nums)
	return nums
}

func main() {
	getConcatenation([]int{1, 2, 1})
}
