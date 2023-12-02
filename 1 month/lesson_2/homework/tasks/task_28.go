package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("A=")
	fmt.Scanln(&a)

	fmt.Printf("A^2 =: %v \nA^3 =: %v\nA^5 =: %v\nA^10 =: %v\nA^15 =: %v\n",
		math.Pow(a, 2),
		math.Pow(a, 3),
		math.Pow(a, 5),
		math.Pow(a, 10),
		math.Pow(a, 15),
	)
}
