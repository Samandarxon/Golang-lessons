package main

import "fmt"

func counter(x []int,n int )int{
	var count int 

	for _,val:=range x {
		if n==val{
			count++
		}
	}
	return count
}

func main(){
	var x_slice = []int{23,24,3,25,34,35,2,56,66,3,76,68,4,3,33,24,5,2}

	fmt.Println("Count :=",counter(x_slice,3))
}