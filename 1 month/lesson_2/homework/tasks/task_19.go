package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	var P, S float64

	fmt.Scanln(&x1)
	fmt.Scanln(&y1)
	fmt.Scanln(&x2)
	fmt.Scanln(&y2)
	fmt.Printf("x1=%f \nx2=%f\ny1=%v\ny2=%v\n", x1, x2, y1, y2)
	P = 2 * (math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1)))
	S = math.Abs(float64(x2-x1)) * math.Abs(float64(y2-y1))
	fmt.Println("Turtburchak perimetiri P =:", P, "\nTurtburchak yuzasi S =:", S)
}
