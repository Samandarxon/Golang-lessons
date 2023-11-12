package main

import "fmt"

func removeElement(nums []int, val int) int {
  
	fmt.Println(len(nums))
	var count = 0
    for _,el:=range nums{
        if el != val{
			// fmt.Println(el,nums)
            nums[count]=el
            count++
        }
    }
	return count
}
func main(){
	fmt.Println(removeElement([]int{3,2,2,4,5,6,3,6,7,8,3},2))
}