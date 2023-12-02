package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("a=")
	fmt.Scanln(&a)

	fmt.Printf("Radian =: %v", a/math.Pi)
}
