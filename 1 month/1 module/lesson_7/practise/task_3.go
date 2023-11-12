package main

import (
	"fmt"
	// "math"
)

func tub(n int ){
	for i:=2 ; i<=n; i++{
		var count = 0
		for j := 2 ; j<=i; j++{
			if i%j==0{
				count++
			}
		}
		if count==1{
			fmt.Printf("%d ",i)
		}
	}
	fmt.Println()
}

func main(){
	tub(13)
}