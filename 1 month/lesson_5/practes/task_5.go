package main

import "fmt"

func Log(x,y int)int{
	var count = 0
	j:=x
	for i:=1 ; i<=j ; i++{
		if x%y==0{
			count++
			// fmt.Println(x,"   ",i)
			x/=y
		}else{
			break
		}
	}
	return count
}

func main(){

	fmt.Println(Log(1024,2))
}
