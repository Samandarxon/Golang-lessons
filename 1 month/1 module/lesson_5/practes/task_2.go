package main

import "fmt"


// EKUB
func EKUB(a,b int) int{
	var x,min,i int=1,0,2
	// var next = 2
	if a<b{
		min=a
	}else{
		min=b
	}
	for min>1{
		if a%i==0 && b%i==0{
			x*=i
		}
			min/=i
		i++
	}
	
	return x
}
func main(){
fmt.Println(EKUB(2,24))
}