package main

import "fmt"

func main(){
	var arr = []int{1,2,3,1,1,3}
	var slice=[]map[int]int{}
	for i,el_1:=range arr{
		for j:=i+1;j<len(arr);j++{
			x_map:=make(map[int]int)
			if el_1==arr[j]{
				// fmt.Println(i,j)
				x_map[i]=j
				slice=append(slice,x_map)
			} 
		}
	}
	fmt.Println(slice)
}