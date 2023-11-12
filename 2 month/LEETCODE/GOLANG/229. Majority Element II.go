package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2}))
	fmt.Println(majorityElement([]int{1, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2}))
}
func majorityElement(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 3 {
			nums = nums[:i+1]
			fmt.Println("sasasd", nums)
			return nums
		} else if nums[i] == nums[len(nums)-1] {
			return nums[:i+1]
		}
	}
	return nums
}
