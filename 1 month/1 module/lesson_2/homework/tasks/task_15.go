package main

import (
	"math"
	"fmt"
)

func main(){
	var (
		S = 3.4
		D float64
		R float64
	)
	const PI = 3.14
	
		R = math.Sqrt(S/PI)
		D = 2*R
		fmt.Printf("R = %.2f\nD = %.2f\n",R, D)


}