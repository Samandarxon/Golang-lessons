package main

import (
	"fmt"
	"math"
)

func main(){
	var a float64


	fmt.Print("A=")
	fmt.Scanln(&a)

	fmt.Printf("A^2 =: %v \nA^4 =: %v\nA^8 =: %v\n",math.Pow(a,2),math.Pow(a,4),math.Pow(a,8))
}