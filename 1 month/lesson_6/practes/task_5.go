package main

import "fmt"

func main(){
	var x_silce =[]int{33,4,2,55,45,65,67,78,91,32,21,24}
	var Total = 0

	for _,val:=range x_silce{
		Total+=val
	}

	fmt.Printf("%.2f\n",float64(Total)/float64(len(x_silce)))
}