package main

import "fmt"

func sumOfUnique(nums []int) int {
    var summ=0
	var x_map=make(map[int]int)
	for _,el :=range nums{
		x_map[el]++
	}
	for key,val:=range x_map{
		if val<2{
			summ+=key
		}
	}
	return summ
}

func main(){
	var nums = []int{1,2,3,2}
	fmt.Println(sumOfUnique(nums))
}

