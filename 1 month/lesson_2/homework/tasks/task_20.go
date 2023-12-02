package main

import (
	"fmt"
	"math"
)

func main(){
	var x1,x2,y1,y2,L float64
	fmt.Print("x1=:")
	fmt.Scanln(&x1)
	fmt.Print("y1=:")
	fmt.Scanln(&y1)
	fmt.Print("x2=:")
	fmt.Scanln(&x2)
	fmt.Print("y2=:")
	fmt.Scanln(&y2)
	L = math.Sqrt(math.Pow(math.Abs(x1-x2),2)+math.Pow(math.Abs(y1-y2),2))
	fmt.Printf("A(%.2f,%.2f) va B(%.2f,%.2f) shu ikki nuqta orasidagi masofa L =: %.2f ga teng\n",x1,y1,x2,y2,L)
}