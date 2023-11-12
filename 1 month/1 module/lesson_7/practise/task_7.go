package main

import "fmt"
func twoSum(nums []int, target int) []int {
    for i,el:=range nums{
		for j:=i+1 ; j<len(nums); j++{
			if el+nums[j]==target{
				return []int{i,j}

			}
		}
	}
	return []int{0,0}
}
func main(){
	fmt.Println(twoSum([]int{3,2,3},6))
}