package main

import "fmt"

func main(){
	// 1. 782838 sondan eng katta maximum raqamini toping
	var number = 782838
	var max = number%10
	for number>0{
		if max<number%10{
			max=number
		}
		number/=10
	}
	fmt.Println(max)
}
