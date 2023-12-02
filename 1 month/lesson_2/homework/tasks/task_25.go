package main

import (
	"fmt"
	"math"
)

func main(){
	var x ,y float64

	fmt.Print("x=")
	fmt.Scanln(&x)
	y = 3*math.Pow(x,6) - 6*x*x - 7
	fmt.Println("y =",y)

}