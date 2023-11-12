package main

import (
	"fmt"
)

func main() {
	var Tf float64

	fmt.Print("a=")
	fmt.Scanln(&Tf)

	fmt.Printf("Selsiy =: %.2f", (Tf-32)*5/9)
}
