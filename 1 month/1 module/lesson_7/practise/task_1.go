package main

import "fmt"

func main(){
	var arr = []int{1, 4, 8,-1, 9, 1, 1, 2, 3, 4}
	var min = arr[0]
	var count = 0
	for _,el:=range arr{
		if min>el{
			min=el
		}
	}
	for _,el:=range arr{
		if min==el{
			count++
		}
	}

	fmt.Printf("min:=%d  count:=%d\n",min,count)

}