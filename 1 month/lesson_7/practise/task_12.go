package main

import "fmt"

func main(){

	 x_slice :=[]int{-1,-2,-9,4,3,1,2,10,}
	var min=1000000000000000
	for _,el:=range x_slice{
		if (el>0 && min>el) {
			min=el
		} 
	}
	if min!=0{
		fmt.Println(min)
	}else{
		min=0
		fmt.Println(min)
	}
}