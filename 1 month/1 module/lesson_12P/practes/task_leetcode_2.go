package main

import "fmt"


func shuffle(nums []int, n int) []int {
	fmt.Println(nums)
    x_slice := []int{}
	for i,_:=range nums[:n]{
		x_slice = append(x_slice,nums[i],nums[n+i])
	}
    
    // for i:=0 ; i<len()
	return x_slice
}

func main(){
fmt.Println(shuffle([]int{2,5,1,3,4,7},3))
}