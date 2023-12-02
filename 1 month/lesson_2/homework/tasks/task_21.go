package main

import (
	"fmt"
	"math"
)

func main(){
	var x1,x2,y1,y2,x3,y3,a,b,c,P,S float64
	fmt.Print("x1=:")
	fmt.Scanln(&x1)
	fmt.Print("y1=:")
	fmt.Scanln(&y1)
	fmt.Print("x2=:")
	fmt.Scanln(&x2)
	fmt.Print("y2=:")
	fmt.Scanln(&y2)
	fmt.Print("x3=:")
	fmt.Scanln(&x3)
	fmt.Print("y3=:")
	fmt.Scanln(&y3)
	a = math.Sqrt(math.Pow(math.Abs(x1-x2),2)+math.Pow(math.Abs(y1-y2),2))
	b = math.Sqrt(math.Pow(math.Abs(x1-x3),2)+math.Pow(math.Abs(y1-y3),2))
	c = math.Sqrt(math.Pow(math.Abs(x3-x2),2)+math.Pow(math.Abs(y3-y2),2))

	P=(a+b+c)/2
	S=math.Sqrt(P*(P-a)*(P-b)*(P-c))

	fmt.Printf("A(%.2f,%.2f), B(%.2f,%.2f) va C(%.2f,%.2f)\nshu nuqtalar orasidagi Uchburchak yuzsai S=%.2f va Perimetri P=%.2f ga teng\n",x1,y1,x2,y2,x3,y3,S,P)

}