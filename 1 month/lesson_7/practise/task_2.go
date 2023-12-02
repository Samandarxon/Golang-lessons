package main

import "fmt"

func main(){
	var arr = []int{1,4,8,9,1,1,2,3,4}
	var average_value int 
	for _,el := range arr {
		average_value += el
	}
	average_value = average_value/len(arr)
	fmt.Println("O'rtacha qiymat :=",average_value)
}