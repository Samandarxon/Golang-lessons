package main

import "fmt"

func EKUK(a, b int) int {
	var ekub,min,i int=1,0,2
	// var next = 2
	if a<b{
		min=a
	}else{
		min=b
	}
	for min>1{
		if a%i==0 && b%i==0{
			ekub*=i
		}
			min/=i
		i++
	}
	ekuk:=a*b/ekub
	
	// fmt.Println(ekub)
	return ekuk
}

func main() {
	fmt.Println(EKUK(7,9))

}
