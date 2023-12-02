package main

import "fmt"

func runSum(nums []int )[]int{
	sum:=[]int{}
	for i,_:=range nums{
		summ:=0
		for j:=0 ; j<=i;j++{
			summ+=nums[j]	
		}
		sum = append(sum,summ)	
	}
	fmt.Println(nums)
	return sum
}

func main(){
	fmt.Println(runSum([]int{1,2,3,4}))
}