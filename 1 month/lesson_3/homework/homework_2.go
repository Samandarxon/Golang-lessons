package main

import "fmt"

func main(){


	var num = 20345234567
	var number = 0

	for num>10{
		number = num%10
		num/=10
		
	}
	fmt.Println(number)
}