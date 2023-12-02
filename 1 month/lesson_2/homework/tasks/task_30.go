package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Print("a=")
	fmt.Scanln(&a)

	fmt.Printf("Gradus =: %.2f", a*math.Pi)
}
