package main

import (
	"fmt"
)

func main() {
	var (
		R1 = 3
		R2 = 4
		S1 float64
		S2 float64
		S3 float64
	)
	const PI = 3.14

	if R1 > R2 {
		S1 = PI * float64(R1)
		S2 = PI * float64(R2)
		S3 = PI * float64(R1-R2)
		fmt.Printf("S1 = %v\nS2 = %v\nS3 = %v\n", S1, S2, S3)
	} else {
		fmt.Println("R1 katta bo'lishi kerak R2 dan")
	}
}
