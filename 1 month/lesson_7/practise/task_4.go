package main

import "fmt"

func sum(n int)int{
	var x int = 0
	for n>0{
		x+=n%10
		n=n/10
	}
	return x
}

func numbers_sum(n int) int{
	// var summ = sum(n)
	// if summ <10{
	// 	return summ
	// }
	// 	return numbers_sum(summ)
	// var summ = sum(n)
	var x int = 0
	for n>0{
		x+=n%10
		n=n/10
		// fmt.Println(n)
	}
	if x <10{
		return x
	}
		return numbers_sum(x)  
}

func main(){
	
	fmt.Println(numbers_sum(456))

}