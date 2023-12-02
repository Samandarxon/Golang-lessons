package main

import (
	"fmt"
)

func main(){
	var num = 23
	var number = 0
	var sum =0
	for num>0{
		number = num%10
		num/=10
		if number%2==0{
			sum+=number
		}
		
	}
	fmt.Println(sum)
}