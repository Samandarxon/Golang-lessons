package main

import (
	"fmt"
)

func main() {
	var Tc float64

	fmt.Print("a=")
	fmt.Scanln(&Tc)

	fmt.Printf("Selsiy =: %.2f", Tc*9/5+32)
}
